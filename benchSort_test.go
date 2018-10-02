package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/tesujiro/OpenDataStructuresGo/ch1"
	"github.com/tesujiro/OpenDataStructuresGo/ch11"
)

func BenchmarkSort_Ch11_MergeSort_NoSort(b *testing.B) {
	benchmarkSort_NoSort(ch11.MergeSort, b)
}

func BenchmarkSort_Ch11_MergeSort_Reverse(b *testing.B) {
	benchmarkSort_Reverse(ch11.MergeSort, b)
}

func BenchmarkSort_Ch11_MergeSort_Random(b *testing.B) {
	benchmarkSort_Random(ch11.MergeSort, b)
}

func benchmarkSort_NoSort(f sortFunc, b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	s := []ch1.Comparable{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = append(s, element(i))
	}
	//fmt.Printf("len(s)=%v\n", len(s))
	b.ResetTimer()
	f(s)
}

func benchmarkSort_Reverse(f sortFunc, b *testing.B) {
	s := []ch1.Comparable{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = append(s, element(b.N-i))
	}
	//fmt.Printf("len(s)=%v\n", len(s))
	b.ResetTimer()
	f(s)
}

func benchmarkSort_Random(f sortFunc, b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	s := []ch1.Comparable{}
	b.ResetTimer()
	for i := 0; len(s) < b.N; i++ {
		s = append(s, element(rand.Intn(b.N)))
	}
	//fmt.Printf("len(s)=%v\n", len(s))
	b.ResetTimer()
	f(s)
}
