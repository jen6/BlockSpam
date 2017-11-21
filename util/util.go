package spamutil

import (
	"regexp"
	"strings"
)

const (
	tokenHttp  = "http://"
	tokenHttps = "https://"
	tokenSlash = "/"

	regexHttp = "^(http|https)://"
)

func IsAbsoluteLink(link string) bool {
	//TODO check another absolute link
	matched, _ := regexp.MatchString(regexHttp, link)
	return matched
}

func ToAbsoluteLink(link string, domain string) string {
	if strings.HasPrefix(link, tokenSlash) {
		return (tokenHttp + domain + link)
	} else {
		return (tokenHttp + domain + tokenSlash + link)
	}
}
