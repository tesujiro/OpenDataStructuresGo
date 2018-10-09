package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/tesujiro/OpenDataStructuresGo/ch1"
	"github.com/tesujiro/OpenDataStructuresGo/ch13"
	"github.com/tesujiro/OpenDataStructuresGo/ch4"
	"github.com/tesujiro/OpenDataStructuresGo/ch5"
	"github.com/tesujiro/OpenDataStructuresGo/ch6"
	"github.com/tesujiro/OpenDataStructuresGo/ch7"
	"github.com/tesujiro/OpenDataStructuresGo/ch8"
)

func BenchmarkSSet_Ch04_SkiplistSSet(b *testing.B) {
	constructor := ch4.NewSkiplistSSet
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
	constructor := ch5.NewChainedHashTable
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
	constructor := ch5.NewLinearHashTable
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
	constructor := ch6.NewBinaryTree
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
	constructor := ch7.NewTreap
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
	constructor := ch8.NewScapegoatTree
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

func benchmarkSSet_AddFirst(s ch1.SSet, b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(element(b.N - i))
	}
}

func benchmarkSSet_AddRandom(s ch1.SSet, b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(element(rand.Int()))
	}
}

func benchmarkSSet_FindFrom1M(s ch1.SSet, b *testing.B) {
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
