package spamparser

import (
	"github.com/jen6/BlockSpam/link"
	"github.com/jen6/BlockSpam/redirect"
	"github.com/jen6/BlockSpam/util"
	"golang.org/x/net/html"
)

const (
	tokenATag = "a"
	tokenHref = "href"
)

func ParseLinks(trResult spamreq.RedirectResult) ([]*link.Link, error) {
	body := trResult.LastResp.Body
	defer body.Close()

	parsedBody := html.NewTokenizer(body)
	links := []string{}
	Exitflag := false
	for {
		ptok := parsedBody.Next()
		switch ptok {
		case html.StartTagToken:
			tok := parsedBody.Token()
			if tok.Data == "a" {
				buf := getAHref(&tok)
				links = append(links, buf...)
			}
		case html.ErrorToken:
			Exitflag = true
			break
		}
		if Exitflag {
			break
		}
	}
	lastLink := trResult.LastLink
	for idx, linkd := range links {
		if !spamutil.IsAbsoluteLink(linkd) {
			domain, err := lastLink.GetDomain()
			if err != nil {
				return []*link.Link{}, err
			}
			links[idx] = spamutil.ToAbsoluteLink(linkd, domain)
		}
	}

	var result []*link.Link
	for _, link := range links {
		bufLink := lastLink.Append(link)
		result = append(result, bufLink)
	}
	return result, nil
}

func getAHref(node *html.Token) []string {
	var links []string
	for _, attr := range node.Attr {
		if attr.Key == tokenHref {
			links = append(links, attr.Val)
			break
		}
	}
	return links
}
