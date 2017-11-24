package blockspam

import (
	"fmt"
	"github.com/jen6/BlockSpam/link"
	"github.com/jen6/BlockSpam/parser"
	"github.com/jen6/BlockSpam/redirect"
	"github.com/jen6/rabinkarp"
	"runtime"
	"sync"
)

func IsSpam(content string, spamLinkDomains []string, redirectionDepth int) bool {
	head := link.NewLinkHead(content)
	reqQueue := make(chan *link.Link, 5000)
	reqQueue <- head

	numCpu := runtime.NumCPU()
	var wg sync.WaitGroup
	var workDoneOnce, workWaitOnce sync.Once
	goWait := make(chan struct{})

	for i := 0; i < numCpu; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				var linkIter *link.Link
				var ok bool
				select {
				case linkIter, ok = <-reqQueue:
					if !ok {
						return
					}
				default:
					fmt.Println("novalue")
					return
				}

				redirectionResult, err := spamreq.GetRedirectLinks(linkIter, redirectionDepth)
				if err != nil {
					workDoneOnce.Do(func() {
						goWait <- struct{}{}
						close(goWait)
					})
					continue
				}
				//if depth > max
				if redirectionResult.LastResp == nil {
					workDoneOnce.Do(func() {
						goWait <- struct{}{}
						close(goWait)
					})
					continue
				}

				statusCode := redirectionResult.LastResp.StatusCode
				statusCode = statusCode - (statusCode % 100)
				if statusCode == 300 {
					workDoneOnce.Do(func() {
						goWait <- struct{}{}
						close(goWait)
					})
					continue
				}

				childLinks, err := spamparser.ParseLinks(redirectionResult)
				if err != nil {
					continue
				}

				for _, clink := range childLinks {
					reqQueue <- clink
				}

				workDoneOnce.Do(func() {
					goWait <- struct{}{}
					close(goWait)
				})
			}
		}()

		workWaitOnce.Do(func() {
			_ = <-goWait
		})
	}
	wg.Wait()

	resultLinks := head.FindDepth(redirectionDepth)
	if len(resultLinks) == 0 {
		return false
	}
	flag := false
	for _, iterLink := range resultLinks {
		if iterLink.Depth > redirectionDepth {
			continue
		}
		domain, _ := iterLink.GetDomain()
		flag = flag || rabinkarp.Search(domain, spamLinkDomains)
	}
	return flag
}
