package spamreq

import (
	"github.com/jen6/BlockSpam/link"
	"net/http"
)

type RedirectResult struct {
	LastLink    *link.Link
	LastResp    *http.Response
	RedirectCnt int
}

func GetRedirectLinks(head *link.Link, maxRedirect int) (RedirectResult, error) {
	lastLink := head
	result := RedirectResult{}
	for i := 0; i < maxRedirect; i++ {
		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			}}
		resp, err := client.Get(lastLink.FullLink)
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
