package api

import (
	"container/list"
	"errors"
	"sync"
	"time"
)

type lruCache[K comparable, V any] struct {
	mut      sync.RWMutex
	capacity int
	list     *list.List
	duration time.Duration
	cache[K]
}

func (l *lruCache[K, V]) Add(key K, val V) {
	l.mut.Lock()
	defer l.mut.Unlock()

	if item, ok := l.cMap[key]; ok {
		item.Value = entry[K, V]{key: key, value: val, exp: time.Now().Add(l.duration)}
		l.list.MoveToFront(item)
	}

	for len(l.cMap) == l.capacity {
		tail := l.list.Back()
		delete(l.cMap, tail.Value.(K))
		l.list.Remove(tail)
	}

	elem := l.list.PushFront(entry[K, V]{key: key, value: val, exp: time.Now().Add(l.duration)})
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

func (l *lruCache[K, V]) isExpired(elem *list.Element) bool {
	return elem.Value.(entry[K, V]).exp.Before(time.Now())
}

func (l *lruCache[K, V]) evict(elem *list.Element) {
	l.mut.Lock()
	defer l.mut.Unlock()

	key := elem.Value.(entry[K, V]).key
	delete(l.cMap, key)
	l.list.Remove(elem)
}

func (l *lruCache[K, V]) remove() {
	time.AfterFunc(l.duration, func() {
		elem := l.list.Back()
		if elem == nil {
			return
		}
		if l.isExpired(elem) {
			l.evict(elem)
		}
	})
}

func (l *lruCache[K, V]) Delete(key K) {
	l.mut.Lock()
	defer l.mut.Unlock()
	if elem, ok := l.cMap[key]; ok {
		l.list.Remove(elem)
		delete(l.cMap, key)
	}
}
