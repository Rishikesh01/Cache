package api

import (
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	nCache := NewCache[int, int](300, time.Minute)
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
	read, err := nCache.Read(199)
	if err != nil {
		return
	}

	if read != 200 {
		t.Error("wrong val")
	}

}
