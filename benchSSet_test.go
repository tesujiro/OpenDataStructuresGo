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

func BenchmarkSSet_Ch04_SkiplistSSet_AddFirst(b *testing.B) {
	s := ch4.NewSkiplistSSet()
	benchmarkSSet_AddFirst(s, b)
}

func BenchmarkSSet_Ch04_SkiplistSSet_AddRandom(b *testing.B) {
	s := ch4.NewSkiplistSSet()
	benchmarkSSet_AddRandom(s, b)
}

func BenchmarkSSet_Ch04_SkiplistSSet_FindFrom1M(b *testing.B) {
	s := ch4.NewSkiplistSSet()
	benchmarkSSet_FindFrom1M(s, b)
}

func BenchmarkSSet_Ch05_ChainedHashTable_AddFirst(b *testing.B) {
	s := ch5.NewChainedHashTable()
	benchmarkSSet_AddFirst(s, b)
}

func BenchmarkSSet_Ch05_ChainedHashTable_AddRandom(b *testing.B) {
	s := ch5.NewChainedHashTable()
	benchmarkSSet_AddRandom(s, b)
}

func BenchmarkSSet_Ch05_ChainedHashTable_FindFrom1M(b *testing.B) {
	s := ch5.NewChainedHashTable()
	benchmarkSSet_FindFrom1M(s, b)
}

func BenchmarkSSet_Ch05_LinearHashTable_AddFirst(b *testing.B) {
	s := ch5.NewLinearHashTable()
	benchmarkSSet_AddFirst(s, b)
}

func BenchmarkSSet_Ch05_LinearHashTable_AddRandom(b *testing.B) {
	s := ch5.NewLinearHashTable()
	benchmarkSSet_AddRandom(s, b)
}

func BenchmarkSSet_Ch05_LinearHashTable_FindFrom1M(b *testing.B) {
	s := ch5.NewLinearHashTable()
	benchmarkSSet_FindFrom1M(s, b)
}

func BenchmarkSSet_Ch06_BinaryTree_AddFirst(b *testing.B) {
	s := ch6.NewBinaryTree()
	benchmarkSSet_AddFirst(s, b)
}

func BenchmarkSSet_Ch06_BinaryTree_AddRandom(b *testing.B) {
	s := ch6.NewBinaryTree()
	benchmarkSSet_AddRandom(s, b)
}

func BenchmarkSSet_Ch06_BinaryTree_FindFrom1M(b *testing.B) {
	s := ch6.NewBinaryTree()
	benchmarkSSet_FindFrom1M(s, b)
}

func BenchmarkSSet_Ch07_Treap_AddFirst(b *testing.B) {
	s := ch7.NewTreap()
	benchmarkSSet_AddFirst(s, b)
}

func BenchmarkSSet_Ch07_Treap_AddRandom(b *testing.B) {
	s := ch7.NewTreap()
	benchmarkSSet_AddRandom(s, b)
}

func BenchmarkSSet_Ch07_Treap_FindFrom1M(b *testing.B) {
	s := ch7.NewTreap()
	benchmarkSSet_FindFrom1M(s, b)
}

func BenchmarkSSet_Ch08_ScapegoatTree_AddFirst(b *testing.B) {
	s := ch8.NewScapegoatTree()
	benchmarkSSet_AddFirst(s, b)
}

func BenchmarkSSet_Ch08_ScapegoatTree_AddRandom(b *testing.B) {
	s := ch8.NewScapegoatTree()
	benchmarkSSet_AddRandom(s, b)
}

func BenchmarkSSet_Ch08_ScapegoatTree_FindFrom1M(b *testing.B) {
	s := ch8.NewScapegoatTree()
	benchmarkSSet_FindFrom1M(s, b)
}

func BenchmarkSSet_Ch13_BinaryTrie_AddFirst(b *testing.B) {
	s := ch13.NewBinaryTrie()
	benchmarkSSet_AddFirst(s, b)
}

func BenchmarkSSet_Ch13_BinaryTrie_AddRandom(b *testing.B) {
	s := ch13.NewBinaryTrie()
	benchmarkSSet_AddRandom(s, b)
}

func BenchmarkSSet_Ch13_BinaryTrie_FindFrom1M(b *testing.B) {
	s := ch13.NewBinaryTrie()
	benchmarkSSet_FindFrom1M(s, b)
}

func BenchmarkSSet_Ch13_XFastTrie_AddFirst(b *testing.B) {
	s := ch13.NewXFastTrie()
	benchmarkSSet_AddFirst(s, b)
}

func BenchmarkSSet_Ch13_XFastTrie_AddRandom(b *testing.B) {
	s := ch13.NewXFastTrie()
	benchmarkSSet_AddRandom(s, b)
}

func BenchmarkSSet_Ch13_XFastTrie_FindFrom1M(b *testing.B) {
	s := ch13.NewXFastTrie()
	benchmarkSSet_FindFrom1M(s, b)
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
