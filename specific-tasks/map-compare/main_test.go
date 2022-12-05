package main

import "testing"

func BenchmarkRemoveDupl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RemoveDupl()
	}
}

func BenchmarkRemoveDuplWithOk(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RemoveDuplWithOk()
	}
}
