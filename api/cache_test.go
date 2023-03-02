package api

import (
	"testing"
)

func TestNewCache(t *testing.T) {
	nCache := NewCache[int, int](300)
	var m []int
	for i := 1; i < 100; i++ {
		m = append(m, i)
	}

	for i := 100; i < 200; i++ {
		m = append(m, i)
	}

	for _, i := range m {
		v := i + 1
		nCache.Add(i, v)
	}

}
