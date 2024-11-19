package lru

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	lru := NewLRUCache[int, int](2)
	lru.Set(1, 1)
	lru.Set(2, 2)
	t.Log(lru)
	lru.Set(3, 3)
	t.Log(lru)
	t.Log(lru.Get(2), lru)
}
