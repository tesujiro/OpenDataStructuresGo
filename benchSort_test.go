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

func BenchmarkSort_Ch11_QuickSort_NoSort(b *testing.B) {
	benchmarkSort_NoSort(ch11.QuickSort, b)
}

func BenchmarkSort_Ch11_QuickSort_Reverse(b *testing.B) {
	benchmarkSort_Reverse(ch11.QuickSort, b)
}

func BenchmarkSort_Ch11_QuickSort_Random(b *testing.B) {
	benchmarkSort_Random(ch11.QuickSort, b)
}

func BenchmarkCSort_Ch11_CountingSort_NoSort(b *testing.B) {
	benchmarkCSort_NoSort(ch11.CountingSort, b)
}

func BenchmarkCSort_Ch11_CountingSort_Reverse(b *testing.B) {
	benchmarkCSort_Reverse(ch11.CountingSort, b)
}

func BenchmarkCSort_Ch11_CountingSort_Random(b *testing.B) {
	benchmarkCSort_Random(ch11.CountingSort, b)
}

func BenchmarkCSort_Ch11_RadixSort_NoSort(b *testing.B) {
	benchmarkCSort_NoSort(ch11.RadixSort, b)
}

func BenchmarkCSort_Ch11_RadixSort_Reverse(b *testing.B) {
	benchmarkCSort_Reverse(ch11.RadixSort, b)
}

func BenchmarkCSort_Ch11_RadixSort_Random(b *testing.B) {
	benchmarkCSort_Random(ch11.RadixSort, b)
}

func benchmarkSort_NoSort(f ch1.SortFunc, b *testing.B) {
	s := []ch1.Comparable{}
	for i := 0; i < b.N; i++ {
		s = append(s, element(i))
	}
	//fmt.Printf("len(s)=%v\n", len(s))
	b.ResetTimer()
	f(s)
}

func benchmarkSort_Reverse(f ch1.SortFunc, b *testing.B) {
	s := []ch1.Comparable{}
	for i := 0; i < b.N; i++ {
		s = append(s, element(b.N-i))
	}
	//fmt.Printf("len(s)=%v\n", len(s))
	b.ResetTimer()
	f(s)
}

func benchmarkSort_Random(f ch1.SortFunc, b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	s := []ch1.Comparable{}
	for i := 0; len(s) < b.N; i++ {
		s = append(s, element(rand.Intn(b.N)))
	}
	//fmt.Printf("len(s)=%v\n", len(s))
	b.ResetTimer()
	f(s)
}

func benchmarkCSort_NoSort(f ch1.CountingSortFunc, b *testing.B) {
	s := []int{}
	k := 256 * 256 * 256 * 8
	for i := 0; i < b.N; i++ {
		s = append(s, i%k)
	}
	//fmt.Printf("len(s)=%v\n", len(s))
	b.ResetTimer()
	f(&s, k)
}

func benchmarkCSort_Reverse(f ch1.CountingSortFunc, b *testing.B) {
	s := []int{}
	k := 256 * 256 * 256 * 8
	for i := 0; i < b.N; i++ {
		s = append(s, (b.N-i)%k)
	}
	//fmt.Printf("len(s)=%v\n", len(s))
	b.ResetTimer()
	f(&s, k)
}

func benchmarkCSort_Random(f ch1.CountingSortFunc, b *testing.B) {
	s := []int{}
	k := 256 * 256 * 256 * 8
	rand.Seed(time.Now().UnixNano())
	for i := 0; len(s) < b.N; i++ {
		s = append(s, rand.Intn(b.N)%k)
	}
	//fmt.Printf("len(s)=%v\n", len(s))
	b.ResetTimer()
	f(&s, k)
}
