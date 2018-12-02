package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/tesujiro/OpenDataStructuresGo/ch01"
	"github.com/tesujiro/OpenDataStructuresGo/ch02"
	"github.com/tesujiro/OpenDataStructuresGo/ch03"
	"github.com/tesujiro/OpenDataStructuresGo/ch04"
)

func BenchmarkList_Ch02_ArrayStack(b *testing.B) {
	constructor := ch02.NewArrayStack
	b.Run("AddFirst", func(b *testing.B) {
		s := constructor()
		benchmarkList_AddFirst(s, b)
	})
	b.Run("AddLast", func(b *testing.B) {
		s := constructor()
		benchmarkList_AddLast(s, b)
	})
	b.Run("AddRandom", func(b *testing.B) {
		s := constructor()
		benchmarkList_AddRandom(s, b)
	})
}

func BenchmarkList_Ch02_ArrayDeque(b *testing.B) {
	constructor := ch02.NewArrayDeque
	b.Run("AddFirst", func(b *testing.B) {
		s := constructor()
		benchmarkList_AddFirst(s, b)
	})
	b.Run("AddLast", func(b *testing.B) {
		s := constructor()
		benchmarkList_AddLast(s, b)
	})
	b.Run("AddRandom", func(b *testing.B) {
		s := constructor()
		benchmarkList_AddRandom(s, b)
	})
}

func BenchmarkList_Ch02_DualArrayDeque(b *testing.B) {
	constructor := ch02.NewDualArrayDeque
	b.Run("AddFirst", func(b *testing.B) {
		s := constructor()
		benchmarkList_AddFirst(s, b)
	})
	b.Run("AddLast", func(b *testing.B) {
		s := constructor()
		benchmarkList_AddLast(s, b)
	})
	b.Run("AddRandom", func(b *testing.B) {
		s := constructor()
		benchmarkList_AddRandom(s, b)
	})
}

func BenchmarkList_Ch02_RootishArrayStack(b *testing.B) {
	constructor := ch02.NewRootishArrayStack
	b.Run("AddFirst", func(b *testing.B) {
		s := constructor()
		benchmarkList_AddFirst(s, b)
	})
	b.Run("AddLast", func(b *testing.B) {
		s := constructor()
		benchmarkList_AddLast(s, b)
	})
	b.Run("AddRandom", func(b *testing.B) {
		s := constructor()
		benchmarkList_AddRandom(s, b)
	})
}

func BenchmarkList_Ch03_DLList(b *testing.B) {
	constructor := ch03.NewDLList
	b.Run("AddFirst", func(b *testing.B) {
		s := constructor()
		benchmarkList_AddFirst(s, b)
	})
	b.Run("AddLast", func(b *testing.B) {
		s := constructor()
		benchmarkList_AddLast(s, b)
	})
	b.Run("AddRandom", func(b *testing.B) {
		s := constructor()
		benchmarkList_AddRandom(s, b)
	})
}

func BenchmarkList_Ch03_SEList(b *testing.B) {
	constructor := ch03.NewSEList
	b.Run("AddFirst", func(b *testing.B) {
		s := constructor()
		benchmarkList_AddFirst(s, b)
	})
	b.Run("AddLast", func(b *testing.B) {
		s := constructor()
		benchmarkList_AddLast(s, b)
	})
	b.Run("AddRandom", func(b *testing.B) {
		s := constructor()
		benchmarkList_AddRandom(s, b)
	})
}

func BenchmarkList_Ch04_SkiplistList(b *testing.B) {
	constructor := ch04.NewSkiplistList
	b.Run("AddFirst", func(b *testing.B) {
		s := constructor()
		benchmarkList_AddFirst(s, b)
	})
	b.Run("AddLast", func(b *testing.B) {
		s := constructor()
		benchmarkList_AddLast(s, b)
	})
	b.Run("AddRandom", func(b *testing.B) {
		s := constructor()
		benchmarkList_AddRandom(s, b)
	})
}

func benchmarkList_AddFirst(s ch01.List, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(0, i)
	}
}

func benchmarkList_AddLast(s ch01.List, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(i, i)
	}
}

func benchmarkList_AddRandom(s ch01.List, b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(rand.Intn(s.Size()+1), i)
	}
}
