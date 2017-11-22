# BlockSpam
## go version
```
# jen6 at localhost in ~/dev/go/src/github.com/jen6/BlockSpam on git:test ✖︎ [11:01:39]
→ go version
go version go1.8.3 linux/amd64
```

## install
first install glide
```
curl https://glide.sh/get | sh
```
On Mac OS X you can also install the latest release via Homebrew:
```
$ brew install glide
```
On Ubuntu Precise (12.04), Trusty (14.04), Wily (15.10) or Xenial (16.04) you can install from our PPA:
```
sudo add-apt-repository ppa:masterminds/glide && sudo apt-get update
sudo apt-get install glide
```
After install glide get package
```
go get github.com/jen6/BlockSpam
glide install
go test
```
