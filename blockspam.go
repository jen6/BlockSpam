package blockspam

import (
	"github.com/jen6/BlockSpam/link"
	"github.com/jen6/BlockSpam/parser"
	"github.com/jen6/BlockSpam/redirect"
	"github.com/jen6/rabinkarp"
)

func IsSpam(content string, spamLinkDomains []string, redirectionDepth int) bool {
	head := link.NewLinkHead(content)
	goWait := make(chan interface{})

	go fetch(head, goWait, redirectionDepth)
	<-goWait
	close(goWait)

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

func fetch(linkIter *link.Link, doneCh chan<- interface{}, redirectionDepth int) {
	redirectionResult, err := spamreq.GetRedirectLinks(linkIter, redirectionDepth)
	if err != nil {
		doneCh <- struct{}{}
		return
	}
	//if depth > max
	if redirectionResult.LastResp == nil {
		doneCh <- struct{}{}
		return
	}

	statusCode := redirectionResult.LastResp.StatusCode
	statusCode = statusCode - (statusCode % 100)
	if statusCode == 300 {
		doneCh <- struct{}{}
		return
	}

	childLinks, err := spamparser.ParseLinks(redirectionResult)
	if err != nil {
		doneCh <- struct{}{}
		return
	}

	resultChannels := make([]chan interface{}, len(childLinks))
	for i, clink := range childLinks {
		rch := make(chan interface{})
		resultChannels[i] = rch
		go fetch(clink, rch, redirectionDepth)
	}

	for _, rch := range resultChannels {
		<-rch
		close(rch)
	}
	doneCh <- struct{}{}
}
