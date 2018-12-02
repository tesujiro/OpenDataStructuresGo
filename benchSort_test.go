package main

import (
	//"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/tesujiro/OpenDataStructuresGo/ch01"
	"github.com/tesujiro/OpenDataStructuresGo/ch11"
)

func BenchmarkSort_Ch11_MergeSort(b *testing.B) {
	benchmarkSort(ch11.MergeSort, b)
}

func BenchmarkSort_Ch11_QuickSort(b *testing.B) {
	benchmarkSort(ch11.QuickSort, b)
}

func BenchmarkCSort_Ch11_CountingSort(b *testing.B) {
	benchmarkCSort(ch11.CountingSort, b)
}

func BenchmarkCSort_Ch11_RadixSort(b *testing.B) {
	benchmarkCSort(ch11.RadixSort, b)
}

func benchmarkSort(f ch01.SortFunc, b *testing.B) {
	b.Run("NoSort", func(b *testing.B) {
		benchmarkSort_NoSort(f, b)
	})
	b.Run("Reverse", func(b *testing.B) {
		benchmarkSort_Reverse(f, b)
	})
	b.Run("Random", func(b *testing.B) {
		benchmarkSort_Random(f, b)
	})

}

func benchmarkSort_NoSort(f ch01.SortFunc, b *testing.B) {
	s := []ch01.Comparable{}
	for i := 0; i < b.N; i++ {
		s = append(s, element(i))
	}
	//fmt.Printf("len(s)=%v\n", len(s))
	b.ResetTimer()
	f(s)
}

func benchmarkSort_Reverse(f ch01.SortFunc, b *testing.B) {
	s := []ch01.Comparable{}
	for i := 0; i < b.N; i++ {
		s = append(s, element(b.N-i))
	}
	//fmt.Printf("len(s)=%v\n", len(s))
	b.ResetTimer()
	f(s)
}

func benchmarkSort_Random(f ch01.SortFunc, b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	s := []ch01.Comparable{}
	for i := 0; len(s) < b.N; i++ {
		s = append(s, element(rand.Intn(b.N)))
	}
	//fmt.Printf("len(s)=%v\n", len(s))
	b.ResetTimer()
	f(s)
}

func benchmarkCSort(f ch01.CountingSortFunc, b *testing.B) {
	b.Run("NoSort", func(b *testing.B) {
		benchmarkCSort_NoSort(f, b)
	})
	b.Run("Reverse", func(b *testing.B) {
		benchmarkCSort_Reverse(f, b)
	})
	b.Run("Random", func(b *testing.B) {
		benchmarkCSort_Random(f, b)
	})

}

func benchmarkCSort_NoSort(f ch01.CountingSortFunc, b *testing.B) {
	//fmt.Printf("b.N=%v start\n", b.N)
	s := []int{}
	k := 256 * 256 * 256 * 1
	for i := 0; i < b.N; i++ {
		s = append(s, i%k)
	}
	//fmt.Printf("len(s)=%v\n", len(s))
	b.ResetTimer()
	f(&s, k)
	//fmt.Printf("b.N=%v finish\n", b.N)
}

func benchmarkCSort_Reverse(f ch01.CountingSortFunc, b *testing.B) {
	s := []int{}
	k := 256 * 256 * 256 * 8
	for i := 0; i < b.N; i++ {
		s = append(s, (b.N-i)%k)
	}
	//fmt.Printf("len(s)=%v\n", len(s))
	b.ResetTimer()
	f(&s, k)
}

func benchmarkCSort_Random(f ch01.CountingSortFunc, b *testing.B) {
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
