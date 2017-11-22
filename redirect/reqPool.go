package spamreq

import (
	"github.com/jen6/BlockSpam/link"
	"sync"
)

type RequestQueue struct {
	rwMx  sync.RWMutex
	queue []*link.Link
}

func (rq *RequestQueue) IsEmpty() bool {
	rq.rwMx.RLock()
	defer rq.rwMx.RUnlock()

	if len(rq.queue) == 0 {
		return true
	} else {
		return false
	}
}

func (rq *RequestQueue) Pop() *link.Link {
	var ret *link.Link
	if !rq.IsEmpty() {
		rq.rwMx.Lock()
		ret, rq.queue = rq.queue[0], rq.queue[1:]
		rq.rwMx.Unlock()
	} else {
		ret = nil
	}
	return ret
}

func (rq *RequestQueue) Push(link *link.Link) {
	rq.rwMx.Lock()
	defer rq.rwMx.Unlock()
	rq.queue = append(rq.queue, link)
}
