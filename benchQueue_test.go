package main

import (
	//"fmt"

	"testing"

	"github.com/tesujiro/OpenDataStructuresGo/ch1"
	"github.com/tesujiro/OpenDataStructuresGo/ch10"
)

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

func benchmarkPriorityQueue_Add(q ch1.PriorityQueue, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Add(element(i))
	}
}

func benchmarkPriorityQueue_Remove(q ch1.PriorityQueue, b *testing.B) {
	for i := 0; i < b.N; i++ {
		q.Add(element(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Remove()
	}
}
