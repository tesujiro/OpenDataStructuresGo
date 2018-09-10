package main

import (
	"testing"
)

func BenchmarkAddLast(b *testing.B) {
	as := newArrayStack()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		as, _ = as.Add(i, i)
	}
}

func BenchmarkAddFirst(b *testing.B) {
	as := newArrayStack()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		as, _ = as.Add(0, i)
	}
}
