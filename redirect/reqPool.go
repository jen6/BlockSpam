package spamreq

import (
	"github.com/jen6/BlockSpam/link"
)

type RequestQueue struct {
	queue []*link.Link
}

func (rq RequestQueue) IsEmpty() bool {
	if len(rq.queue) == 0 {
		return true
	} else {
		return false
	}
}

func (rq *RequestQueue) Pop() *link.Link {
	var ret *link.Link
	if !rq.IsEmpty() {
		ret, rq.queue = rq.queue[0], rq.queue[1:]
	} else {
		ret = nil
	}
	return ret
}

func (rq *RequestQueue) Push(link *link.Link) {
	rq.queue = append(rq.queue, link)
}
