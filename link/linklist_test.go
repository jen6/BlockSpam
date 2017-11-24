package link

import (
	"reflect"
	"testing"
)

const (
	aFullLink  = "http://a.com/aa"
	bFullLink  = "http://b.com/bb"
	cFullLink  = "http://c.com/cc"
	aaFullLink = "http://aa.com/aaa"

	aDomain  = "a.com"
	bDomain  = "b.com"
	cDomain  = "c.com"
	aaDomain = "aa.com"
)

func makeTestLink() *Link {
	head := NewLinkHead(aFullLink)
	b := head.Append(bFullLink)
	_ = b.Append(aaFullLink)

	c := head.Append(cFullLink)
	_ = c.Append(aaFullLink)

	return head
}

func TestLink_FindDepth(t *testing.T) {
	type args struct {
		depth int
	}
	tests := []struct {
		name string
		link *Link
		args args
		want []string
	}{
		{"find depth1", makeTestLink(), args{depth: 1}, []string{bFullLink, cFullLink}},
		{"find depth2", makeTestLink(), args{depth: 2}, []string{aaFullLink}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.link
			got := l.FindDepth(tt.args.depth)
			gotLink := []string{}
			for _, g := range got {
				gotLink = append(gotLink, g.FullLink)
			}
			if !reflect.DeepEqual(gotLink, tt.want) {
				t.Errorf("Link.FindDepth() = %v, want %v", gotLink, tt.want)
			}
		})
	}
}
