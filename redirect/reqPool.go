package spamreq

import (
	"github.com/jen6/BlockSpam/link"
)

type RequestPool struct {
	pool []*link.Link
}

func (rp RequestPool) IsEmpty() bool {
	if len(rp.pool) > 0 {
		return true
	} else {
		return false
	}
}

func (rp *RequestPool) Pop() *link.Link {
	var ret *link.Link
	if rp.IsEmpty() {
		ret, rp.pool = rp.pool[0], rp.pool[1:]
	} else {
		ret = nil
	}
	return ret
}

func (rp *RequestPool) Push(link *link.Link) {
	rp.pool = append(rp.pool, link)
}
