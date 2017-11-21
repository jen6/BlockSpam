package link

import (
	"net/url"
)

type Link struct {
	childNodes []*Link
	FullLink   string
	Depth      int
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
func (l *Link) FindDepth(depth int, domain string) (*Link, error) {
	if depth == 1 {
		lDomain, err := l.GetDomain()
		if err != nil {
			return nil, err
		}

		if lDomain == domain {
			return l, nil
		} else {
			return nil, nil
		}
	}

	var result *Link
	result = nil

	for _, link := range l.childNodes {
		ret, err := link.FindDepth(depth-1, domain)
		if err != nil {
			return nil, err
		}
		if ret != nil {
			result = ret
			break
		}
	}
	return result, nil
}
