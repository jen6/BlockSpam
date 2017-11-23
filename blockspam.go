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
	reqQueue := spamreq.RequestQueue{}
	reqQueue.Push(head)

	numCpu := runtime.NumCPU()
	var wg sync.WaitGroup

	for i := 0; i < numCpu; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				if reqQueue.IsEmpty() {
					return
				}

				linkIter := reqQueue.Pop()
				redirectionResult, err := spamreq.GetRedirectLinks(linkIter, redirectionDepth)
				if err != nil {
					continue
				}
				//if depth > max
				if redirectionResult.LastResp == nil {
					continue
				}

				statusCode := redirectionResult.LastResp.StatusCode
				statusCode = statusCode - (statusCode % 100)
				if statusCode == 300 {
					continue
				}

				childLinks, err := spamparser.ParseLinks(redirectionResult)
				if err != nil {
					continue
				}

				for _, clink := range childLinks {
					reqQueue.Push(clink)
				}
			}
		}()
	}
	wg.Wait()

	resultLinks := head.FindDepth(redirectionDepth)
	if len(resultLinks) == 0 {
		return false
	}
	flag := false
	for _, iterLink := range resultLinks {
		domain, _ := iterLink.GetDomain()
		flag = flag || rabinkarp.Search(domain, spamLinkDomains)
	}
	return flag
}
