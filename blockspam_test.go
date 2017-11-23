package blockspam

import "testing"

const (
	link1       = "https://goo.gl/nVLutc"
	link2       = "http://bit.ly/2yTkW52"
	link3       = "http://tvtv24.com/view.php?id=intro&no=58"
	link4       = "http://www.filekok.com"
	link5       = "http://github.com/jen6"
	link1Domain = "goo.gl"
	link2Domain = "bit.ly"
	link3Domain = "tvtv24.com"
	link4Domain = "www.filekok.com"
	link5Domain = "github.com"
	link6Domain = "jen6.tistory.com"
	link7Domain = "facebook.com"
	errorLink   = "aaaa"
	errorDomain = "aaa"
)

func TestIsSpam(t *testing.T) {
	type args struct {
		content          string
		spamLinkDomains  []string
		redirectionDepth int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"t1", args{link1, []string{link4Domain}, 2}, false, false},
		{"t2", args{link1, []string{link3Domain}, 1}, true, false},
		{"t3", args{link2, []string{link1Domain}, 1}, true, false},
		{"t4", args{link2, []string{link3Domain}, 2}, true, false},
		{"t5", args{link5, []string{link6Domain}, 2}, true, false},
		{"t6", args{errorLink, []string{errorDomain}, 2}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsSpam(tt.args.content, tt.args.spamLinkDomains, tt.args.redirectionDepth)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsSpam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsSpam() = %v, want %v", got, tt.want)
			}
		})
	}
}
