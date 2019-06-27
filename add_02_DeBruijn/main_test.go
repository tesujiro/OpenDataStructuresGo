package main

import "testing"

func BenchmarkSeq(b *testing.B) {
	for i := 1; i < b.N; i++ {
		d := newD(2, 2)
		_, _ = d.seq()
	}
}
