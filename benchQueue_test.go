package main

import (
	//"fmt"

	"testing"

	"github.com/tesujiro/OpenDataStructuresGo/ch01"
	"github.com/tesujiro/OpenDataStructuresGo/ch02"
	"github.com/tesujiro/OpenDataStructuresGo/ch03"
	"github.com/tesujiro/OpenDataStructuresGo/ch10"
)

func BenchmarkQueue_Ch2_ArrayQueue(b *testing.B) {
	constructor := ch02.NewArrayQueue
	b.Run("Add", func(b *testing.B) {
		q := constructor()
		benchmarkQueue_Add(q, b)
	})
	b.Run("Remove", func(b *testing.B) {
		q := constructor()
		benchmarkQueue_Remove(q, b)
	})
}

func BenchmarkQueue_Ch3_SLList(b *testing.B) {
	constructor := ch03.NewSLList
	b.Run("Add", func(b *testing.B) {
		q := constructor()
		benchmarkQueue_Add(q, b)
	})
	b.Run("Remove", func(b *testing.B) {
		q := constructor()
		benchmarkQueue_Remove(q, b)
	})
}

func benchmarkQueue_Add(q ch01.Queue, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Add(i)
	}
}

func benchmarkQueue_Remove(q ch01.Queue, b *testing.B) {
	for i := 0; i < b.N; i++ {
		q.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Remove()
	}
}
func BenchmarkPriorityQueue_Ch10_BinaryHeap(b *testing.B) {
	constructor := ch10.NewBinaryHeap
	b.Run("Add", func(b *testing.B) {
		q := constructor()
		benchmarkPriorityQueue_Add(q, b)
	})
	b.Run("Remove", func(b *testing.B) {
		q := constructor()
		benchmarkPriorityQueue_Remove(q, b)
	})
}

func BenchmarkPriorityQueue_Ch10_MeldableHeap(b *testing.B) {
	constructor := ch10.NewMeldableHeap
	b.Run("Add", func(b *testing.B) {
		q := constructor()
		benchmarkPriorityQueue_Add(q, b)
	})
	b.Run("Remove", func(b *testing.B) {
		q := constructor()
		benchmarkPriorityQueue_Remove(q, b)
	})
}

func benchmarkPriorityQueue_Add(q ch01.PriorityQueue, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Add(element(i))
	}
}

func benchmarkPriorityQueue_Remove(q ch01.PriorityQueue, b *testing.B) {
	for i := 0; i < b.N; i++ {
		q.Add(element(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Remove()
	}
}
