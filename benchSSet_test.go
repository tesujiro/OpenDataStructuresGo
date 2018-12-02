package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/tesujiro/OpenDataStructuresGo/ch01"
	"github.com/tesujiro/OpenDataStructuresGo/ch04"
	"github.com/tesujiro/OpenDataStructuresGo/ch05"
	"github.com/tesujiro/OpenDataStructuresGo/ch06"
	"github.com/tesujiro/OpenDataStructuresGo/ch07"
	"github.com/tesujiro/OpenDataStructuresGo/ch08"
	"github.com/tesujiro/OpenDataStructuresGo/ch13"
)

func BenchmarkSSet_Ch04_SkiplistSSet(b *testing.B) {
	constructor := ch04.NewSkiplistSSet
	b.Run("AddFirst", func(b *testing.B) {
		s := constructor()
		benchmarkSSet_AddFirst(s, b)
	})
	b.Run("AddRandom", func(b *testing.B) {
		s := constructor()
		benchmarkSSet_AddRandom(s, b)
	})
	b.Run("FindFrom1M", func(b *testing.B) {
		s := constructor()
		benchmarkSSet_FindFrom1M(s, b)
	})
}

func BenchmarkSSet_Ch05_ChainedHashTable(b *testing.B) {
	constructor := ch05.NewChainedHashTable
	b.Run("AddFirst", func(b *testing.B) {
		s := constructor()
		benchmarkSSet_AddFirst(s, b)
	})
	b.Run("AddRandom", func(b *testing.B) {
		s := constructor()
		benchmarkSSet_AddRandom(s, b)
	})
	b.Run("FindFrom1M", func(b *testing.B) {
		s := constructor()
		benchmarkSSet_FindFrom1M(s, b)
	})
}

func BenchmarkSSet_Ch05_LinearHashTable(b *testing.B) {
	constructor := ch05.NewLinearHashTable
	b.Run("AddFirst", func(b *testing.B) {
		s := constructor()
		benchmarkSSet_AddFirst(s, b)
	})
	b.Run("AddRandom", func(b *testing.B) {
		s := constructor()
		benchmarkSSet_AddRandom(s, b)
	})
	b.Run("FindFrom1M", func(b *testing.B) {
		s := constructor()
		benchmarkSSet_FindFrom1M(s, b)
	})
}

func BenchmarkSSet_Ch06_BinaryTree(b *testing.B) {
	constructor := ch06.NewBinaryTree
	b.Run("AddFirst", func(b *testing.B) {
		s := constructor()
		benchmarkSSet_AddFirst(s, b)
	})
	b.Run("AddRandom", func(b *testing.B) {
		s := constructor()
		benchmarkSSet_AddRandom(s, b)
	})
	b.Run("FindFrom1M", func(b *testing.B) {
		s := constructor()
		benchmarkSSet_FindFrom1M(s, b)
	})
}

func BenchmarkSSet_Ch07_Treap(b *testing.B) {
	constructor := ch07.NewTreap
	b.Run("AddFirst", func(b *testing.B) {
		s := constructor()
		benchmarkSSet_AddFirst(s, b)
	})
	b.Run("AddRandom", func(b *testing.B) {
		s := constructor()
		benchmarkSSet_AddRandom(s, b)
	})
	b.Run("FindFrom1M", func(b *testing.B) {
		s := constructor()
		benchmarkSSet_FindFrom1M(s, b)
	})
}

func BenchmarkSSet_Ch08_ScapegoatTree(b *testing.B) {
	constructor := ch08.NewScapegoatTree
	b.Run("AddFirst", func(b *testing.B) {
		s := constructor()
		benchmarkSSet_AddFirst(s, b)
	})
	b.Run("AddRandom", func(b *testing.B) {
		s := constructor()
		benchmarkSSet_AddRandom(s, b)
	})
	b.Run("FindFrom1M", func(b *testing.B) {
		s := constructor()
		benchmarkSSet_FindFrom1M(s, b)
	})
}

func BenchmarkSSet_Ch13_BinaryTrie(b *testing.B) {
	constructor := ch13.NewBinaryTrie
	b.Run("AddFirst", func(b *testing.B) {
		s := constructor()
		benchmarkSSet_AddFirst(s, b)
	})
	b.Run("AddRandom", func(b *testing.B) {
		s := constructor()
		benchmarkSSet_AddRandom(s, b)
	})
	b.Run("FindFrom1M", func(b *testing.B) {
		s := constructor()
		benchmarkSSet_FindFrom1M(s, b)
	})
}

func BenchmarkSSet_Ch13_XFastTrie(b *testing.B) {
	constructor := ch13.NewXFastTrie
	b.Run("AddFirst", func(b *testing.B) {
		s := constructor()
		benchmarkSSet_AddFirst(s, b)
	})
	b.Run("AddRandom", func(b *testing.B) {
		s := constructor()
		benchmarkSSet_AddRandom(s, b)
	})
	b.Run("FindFrom1M", func(b *testing.B) {
		s := constructor()
		benchmarkSSet_FindFrom1M(s, b)
	})
}

func benchmarkSSet_AddFirst(s ch01.SSet, b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(element(b.N - i))
	}
}

func benchmarkSSet_AddRandom(s ch01.SSet, b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(element(rand.Int()))
	}
}

func benchmarkSSet_FindFrom1M(s ch01.SSet, b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	n := 1000000
	for s.Size() < n {
		s.Add(element(rand.Intn(n)))
	}
	//fmt.Println("Size:", s.Size())
	count := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if s.Find(element(rand.Intn(n))) != nil {
			count++
		}
	}
	//fmt.Printf("Found/All:%3.2f%%\n", float64(count*100)/float64(b.N))
}
