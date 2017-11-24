package spamutil

import (
	"sync"
)

type LinkMap struct {
	sync.Map
}

var instance *LinkMap
var once sync.Once

func GetLinkMap() *LinkMap {
	once.Do(func() {
		instance = &LinkMap{}
	})
	return instance
}
