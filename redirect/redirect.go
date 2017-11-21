package spamreq

import (
	"github.com/jen6/BlockSpam/link"
	"github.com/jen6/BlockSpam/util"
	"net"
	"net/http"
	"time"
)

type RedirectResult struct {
	LastLink    *link.Link
	LastResp    *http.Response
	RedirectCnt int
}

func GetRedirectLinks(head *link.Link, maxRedirect int) (RedirectResult, error) {
	lastLink := head
	originDomain, _ := lastLink.GetDomain()

	result := RedirectResult{}
	for i := head.Depth; i <= maxRedirect; i++ {

		var netTransport = &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 5 * time.Second,
		}

		client := &http.Client{
			Timeout:   time.Second * 10,
			Transport: netTransport,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			}}

		var targetLink string
		if !spamutil.IsAbsoluteLink(lastLink.FullLink) {
			domain, err := lastLink.GetDomain()
			if err != nil {
				return RedirectResult{}, err
			}
			if domain == "" {
				targetLink = spamutil.ToAbsoluteLink(lastLink.FullLink, originDomain)
			} else {
				targetLink = spamutil.ToAbsoluteLink(lastLink.FullLink, domain)
			}
		} else {
			targetLink = lastLink.FullLink
		}
		resp, err := client.Get(targetLink)
		if err != nil {
			return RedirectResult{}, err
		}

		lastLink = lastLink.Append(resp.Header.Get("Location"))
		result.LastResp = resp
		result.RedirectCnt = i + 1

		if resp.StatusCode == 200 {
			break
		}
	}
	result.LastLink = lastLink
	return result, nil
}
