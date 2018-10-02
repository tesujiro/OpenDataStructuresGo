package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/tesujiro/OpenDataStructuresGo/ch1"
	"github.com/tesujiro/OpenDataStructuresGo/ch2"
	"github.com/tesujiro/OpenDataStructuresGo/ch3"
	"github.com/tesujiro/OpenDataStructuresGo/ch4"
)

func BenchmarkList_Ch02_ArrayStack_AddFirst(b *testing.B) {
	s := ch2.NewArrayStack()
	benchmarkList_AddFirst(s, b)
}

func BenchmarkList_Ch02_ArrayStack_AddLast(b *testing.B) {
	s := ch2.NewArrayStack()
	benchmarkList_AddLast(s, b)
}

func BenchmarkList_Ch02_ArrayStack_AddRandom(b *testing.B) {
	s := ch2.NewArrayStack()
	benchmarkList_AddRandom(s, b)
}

func BenchmarkList_Ch02_ArrayDeque_AddFirst(b *testing.B) {
	s := ch2.NewArrayDeque()
	benchmarkList_AddFirst(s, b)
}

func BenchmarkList_Ch02_ArrayDeque_AddLast(b *testing.B) {
	s := ch2.NewArrayDeque()
	benchmarkList_AddLast(s, b)
}

func BenchmarkList_Ch02_ArrayDeque_AddRandom(b *testing.B) {
	s := ch2.NewArrayDeque()
	benchmarkList_AddRandom(s, b)
}

func BenchmarkList_Ch03_DLList_AddFirst(b *testing.B) {
	s := ch3.NewDLList()
	benchmarkList_AddFirst(s, b)
}

func BenchmarkList_Ch03_DLList_AddLast(b *testing.B) {
	s := ch3.NewDLList()
	benchmarkList_AddLast(s, b)
}

func BenchmarkList_Ch03_DLList_AddRandom(b *testing.B) {
	s := ch3.NewDLList()
	benchmarkList_AddRandom(s, b)
}

func BenchmarkList_Ch03_SEList_AddFirst(b *testing.B) {
	s := ch3.NewSEList()
	benchmarkList_AddFirst(s, b)
}

func BenchmarkList_Ch03_SEList_AddLast(b *testing.B) {
	s := ch3.NewSEList()
	benchmarkList_AddLast(s, b)
}

func BenchmarkList_Ch03_SEList_AddRandom(b *testing.B) {
	s := ch3.NewSEList()
	benchmarkList_AddRandom(s, b)
}

func BenchmarkList_Ch04_SkiplistList_AddFirst(b *testing.B) {
	s := ch4.NewSkiplistList()
	benchmarkList_AddFirst(s, b)
}

func BenchmarkList_Ch04_SkiplistList_AddLast(b *testing.B) {
	s := ch4.NewSkiplistList()
	benchmarkList_AddLast(s, b)
}

func BenchmarkList_Ch04_SkiplistList_AddRandom(b *testing.B) {
	s := ch4.NewSkiplistList()
	benchmarkList_AddRandom(s, b)
}

func benchmarkList_AddFirst(s ch1.List, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(0, i)
	}
}

func benchmarkList_AddLast(s ch1.List, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(i, i)
	}
}

func benchmarkList_AddRandom(s ch1.List, b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(rand.Intn(s.Size()+1), i)
	}
}
