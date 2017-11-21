package spamreq

import (
	"testing"

	"github.com/jen6/BlockSpam/link"
)

const (
	testShort   = "https://goo.gl/nVLutc"
	finalDomain = "tvtv24.com"
)

func makeTestLink() *link.Link {
	return &link.Link{FullLink: testShort, Depth: 1}
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
		{"shortTest", args{makeTestLink(), 3}, finalDomain, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRedirectLinks(tt.args.head, tt.args.maxRedirect)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRedirectLinks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.LastLink == nil {
				t.Errorf("GetRedirectLinks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			dom, _ := got.LastLink.GetDomain()
			if dom == tt.want {
				t.Errorf("GetRedirectLinks() = %v, want %v", got, tt.want)
			}
		})
	}
}
