package link

import (
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
	head := &Link{FullLink: aFullLink, Depth: 1}
	b := head.Append(bFullLink)
	_ = b.Append(aaFullLink)

	c := head.Append(cFullLink)
	_ = c.Append(aaFullLink)

	return head
}

func TestLink_FindDepth(t *testing.T) {
	type args struct {
		depth  int
		domain string
	}
	tests := []struct {
		name    string
		link    *Link
		args    args
		want    string
		wantErr bool
	}{
		{"find b", makeTestLink(), args{depth: 2, domain: bDomain}, bFullLink, false},
		{"find aa", makeTestLink(), args{depth: 3, domain: aaDomain}, aaFullLink, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.link
			got, err := l.FindDepth(tt.args.depth, tt.args.domain)
			if (err != nil) != tt.wantErr {
				t.Errorf("Link.FindDepth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil && !tt.wantErr {
				t.Errorf("Link.FindDepth() = %v, want %v", got, tt.want)
				return
			}
			if got.FullLink != tt.want {
				t.Errorf("Link.FindDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
