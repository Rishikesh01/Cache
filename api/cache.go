package api

import (
	"container/list"
	"time"
)

type Cache[K comparable, V any] interface {
	Add(K, V)
	Read(K) (V, error)
	Delete(K)
}

type cache[T comparable] struct {
	cMap map[T]*list.Element
}

type entry[T, V any] struct {
	key   T
	value V
	exp   time.Time
}

func NewCache[K comparable, V any](capacity int, duration time.Duration) Cache[K, V] {
	lru := &lruCache[K, V]{
		cache:    cache[K]{cMap: make(map[K]*list.Element)},
		list:     list.New(),
		capacity: capacity,
		duration: duration,
	}
	go lru.remove()
	return lru
}
