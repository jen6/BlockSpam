package spamreq

import (
	"github.com/jen6/BlockSpam/link"
	"sync"
)

type RequestQueue struct {
	mx    sync.Mutex
	queue []*link.Link
}

func (rq *RequestQueue) isEmpty() bool {
	if len(rq.queue) == 0 {
		return true
	} else {
		return false
	}
}

func (rq *RequestQueue) IsEmpty() bool {
	rq.mx.Lock()
	result := rq.isEmpty()
	rq.mx.Unlock()
	return result
}

func (rq *RequestQueue) Pop() *link.Link {
	rq.mx.Lock()
	defer rq.mx.Unlock()
	var ret *link.Link
	if !rq.isEmpty() {
		ret, rq.queue = rq.queue[0], rq.queue[1:]
	} else {
		ret = nil
	}
	return ret
}

func (rq *RequestQueue) Push(link *link.Link) {
	rq.mx.Lock()
	defer rq.mx.Unlock()
	rq.queue = append(rq.queue, link)
}
