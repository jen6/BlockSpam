package spamparser

import (
	"fmt"
	"github.com/jen6/BlockSpam/link"
	"github.com/jen6/BlockSpam/redirect"
	"golang.org/x/net/html"
	"regexp"
	"strings"
)

const (
	tokenATag  = "a"
	tokenHref  = "href"
	tokenHttp  = "http://"
	tokenHttps = "https://"
	tokenSlash = "/"

	regexHttp = "^(http|https)://"
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
		if !isAbsoluteLink(linkd) || linkd == "" {
			domain, err := lastLink.GetDomain()
			if err != nil {
				return []*link.Link{}, err
			}
			links[idx] = toAbsoluteLink(linkd, domain)
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
			fmt.Println(links)
			break
		}
	}
	return links
}

func isAbsoluteLink(link string) bool {
	//TODO check another absolute link
	matched, _ := regexp.MatchString(regexHttp, link)
	return matched
}

func toAbsoluteLink(link string, domain string) string {
	if strings.HasPrefix(link, tokenSlash) {
		return (tokenHttp + domain + link)
	} else {
		return (tokenHttp + domain + tokenSlash + link)
	}
}
