package api

import (
	"container/list"
	"errors"
	"sync"
)

type lruCache[K comparable, V any] struct {
	mut      sync.RWMutex
	capacity int
	list     *list.List
	cache[K]
}

func (l *lruCache[K, V]) Add(key K, val V) {
	l.mut.Lock()
	defer l.mut.Unlock()

	if item, ok := l.cMap[key]; ok {
		item.Value = entry[K, V]{key: key, value: val}
		l.list.MoveToFront(item)
	}

	if len(l.cMap) == l.capacity {
		tail := l.list.Back()
		delete(l.cMap, tail.Value.(K))
		l.list.Remove(tail)
	}

	elem := l.list.PushFront(entry[K, V]{key: key, value: val})
	l.cMap[key] = elem
}

func (l *lruCache[K, V]) Read(key K) (V, error) {
	if l.list.Front().Value.(entry[K, V]).key == key {
		return l.list.Front().Value.(entry[K, V]).value, nil
	}
	l.mut.RLock()
	defer l.mut.RUnlock()
	if elem, ok := l.cMap[key]; ok {
		l.list.MoveToFront(elem)
		return elem.Value.(entry[K, V]).value, nil
	}
	return *new(V), errors.New("val does not exists")
}

func (l *lruCache[K, V]) Delete(key K) {
	l.mut.Lock()
	defer l.mut.Unlock()
	if elem, ok := l.cMap[key]; ok {
		l.list.Remove(elem)
		delete(l.cMap, key)
	}
}
