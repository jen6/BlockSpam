package spamparser

import (
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

func ParseLinks(domain string, trResult spamreq.RedirectResult) ([]*link.Link, error) {
	body := trResult.LastResp.Body
	defer body.Close()

	parsedBody, err := html.Parse(body)
	if err != nil {
		return []*link.Link{}, err
	}

	links := getAHref(parsedBody)
	lastLink := trResult.LastLink
	for idx, linkd := range links {
		if !isAbsoluteLink(linkd) {
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

func getAHref(node *html.Node) []string {
	var links []string
	var loopParse func(node *html.Node)
	loopParse = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == tokenATag {
			for _, attr := range node.Attr {
				if attr.Key == tokenHref {
					links = append(links, attr.Val)
					break
				}
			}
		}
		for iterNode := node.FirstChild; iterNode != nil; iterNode = iterNode.NextSibling {
			loopParse(iterNode)
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
