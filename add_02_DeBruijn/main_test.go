package main

import "testing"

func BenchmarkSeq(b *testing.B) {
	for i := 1; i < b.N; i++ {
		d := newD(8, 5, true)
		_, _ = d.seq()
	}
}
