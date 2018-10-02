package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/tesujiro/OpenDataStructuresGo/ch1"
	"github.com/tesujiro/OpenDataStructuresGo/ch11"
)

func BenchmarkSort(b *testing.B) {
	benchmarkSort_Reverse(ch11.MergeSort, b)
}

func benchmarkSort_Reverse(f sortfunc, b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	s := []ch1.Comparable{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = append(s, element(b.N-i))
	}
	b.ResetTimer()
	f(s)
}
