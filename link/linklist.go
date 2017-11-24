package link

import (
	"net/url"
)

type Link struct {
	childNodes []*Link
	FullLink   string
	Depth      int
}

func NewLinkHead(url string) *Link {
	return &Link{FullLink: url, Depth: 0}
}

func (l *Link) Append(fullLink string) *Link {
	newLink := &Link{FullLink: fullLink, Depth: l.Depth + 1}
	l.childNodes = append(l.childNodes, newLink)
	return newLink
}

func (l Link) GetDomain() (string, error) {
	u, err := url.Parse(l.FullLink)
	if err != nil {
		return "", err
	}
	return u.Host, nil
}

//if find Link return pointer of Link else nil
func (l *Link) FindDepth(depth int) []*Link {
	if depth == 0 {
		return []*Link{l}
	}

	result := []*Link{}
	for _, link := range l.childNodes {
		ret := link.FindDepth(depth - 1)
		if len(ret) != 0 {
			result = append(result, ret...)
		}
	}
	return removeDuplicateLink(result)
}

func removeDuplicateLink(links []*Link) []*Link {
	linkMap := map[string]bool{}
	result := []*Link{}

	for _, link := range links {
		ok, _ := linkMap[link.FullLink]
		if !ok {
			linkMap[link.FullLink] = true
			result = append(result, link)
		} else {
			continue
		}
	}
	return result
}
