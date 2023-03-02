package api

import (
	"container/list"
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
}

func NewCache[K comparable, V any](capacity int) Cache[K, V] {
	return &lruCache[K, V]{
		cache:    cache[K]{cMap: make(map[K]*list.Element)},
		list:     list.New(),
		capacity: capacity,
	}
}
