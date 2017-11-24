package spamreq

import (
	"testing"

	"github.com/jen6/BlockSpam/link"
)

const (
	errorShort   = "aaa"
	testShort    = "http://bit.ly/2yTkW52"
	middleDomain = "https://goo.gl/nVLutc"
	finalDomain  = "http://tvtv24.com/view.php?id=intro&no=58"
)

func makeTestLink() *link.Link {
	return link.NewLinkHead(testShort)
}
func makeTestErrorLink() *link.Link {
	return link.NewLinkHead(errorShort)
}

func TestGetRedirectLinks(t *testing.T) {
	type args struct {
		head        *link.Link
		maxRedirect int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"makeError", args{makeTestErrorLink(), 4}, "", true},
		{"trackMiddleDomain", args{makeTestLink(), 1}, middleDomain, false},
		{"trackFinalDomain", args{makeTestLink(), 2}, finalDomain, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRedirectLinks(tt.args.head, tt.args.maxRedirect)
			if tt.wantErr {
				if err == nil {
					t.Errorf("GetRedirectLinks() error = %v, wantErr %v", err, tt.wantErr)
					return
				} else {
					return
				}
			}
			dom, _ := got.LastLink.GetDomain()
			if dom == tt.want {
				t.Errorf("GetRedirectLinks() = %v, want %v", got, tt.want)
			}
		})
	}
}
