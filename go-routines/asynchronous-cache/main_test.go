package main

import "testing"

func BenchmarkCacheWork(b *testing.B) {
	cache := NewCache()

	for i := 0; i < b.N; i++ {
		CacheWork(&cache)
	}
}

func TestCacheWork(t *testing.T) {
	cache := NewCache()

	CacheWork(&cache)
}
