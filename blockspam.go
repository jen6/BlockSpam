package blockspam

import (
	"fmt"
	"github.com/jen6/BlockSpam/link"
	"github.com/jen6/BlockSpam/parser"
	"github.com/jen6/BlockSpam/redirect"
	"github.com/jen6/rabinkarp"
)

func IsSpam(content string, spamLinkDomains []string, redirectionDepth int) (bool, error) {
	head := link.NewLinkHead(content)
	reqQueue := spamreq.RequestQueue{}
	reqQueue.Push(head)

	for {
		if reqQueue.IsEmpty() {
			break
		}

		linkIter := reqQueue.Pop()
		redirectionResult, err := spamreq.GetRedirectLinks(linkIter, redirectionDepth)
		if err != nil {
			fmt.Println(err)
			return false, err
		}

		statusCode := redirectionResult.LastResp.StatusCode
		if statusCode == 301 || statusCode == 302 {
			continue
		}

		childLinks, err := spamparser.ParseLinks(redirectionResult)
		if err != nil {
			fmt.Println(err)
			return false, err
		}

		for _, clink := range childLinks {
			reqQueue.Push(clink)
		}
	}

	resultLinks := head.FindDepth(redirectionDepth)
	if len(resultLinks) == 0 {
		return false, nil
	}
	flag := false
	for _, iterLink := range resultLinks {
		domain, err := iterLink.GetDomain()
		if err != nil {
			return false, err
		}
		flag = flag || rabinkarp.Search(domain, spamLinkDomains)
	}
	return flag, nil
}