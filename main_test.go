package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"

	"github.com/tesujiro/OpenDataStructuresGo/ch1"
	"github.com/tesujiro/OpenDataStructuresGo/ch11"
	"github.com/tesujiro/OpenDataStructuresGo/ch13"
	"github.com/tesujiro/OpenDataStructuresGo/ch2"
	"github.com/tesujiro/OpenDataStructuresGo/ch3"
	"github.com/tesujiro/OpenDataStructuresGo/ch4"
	"github.com/tesujiro/OpenDataStructuresGo/ch5"
	"github.com/tesujiro/OpenDataStructuresGo/ch6"
	"github.com/tesujiro/OpenDataStructuresGo/ch7"
	"github.com/tesujiro/OpenDataStructuresGo/ch8"
)

func TestArrayStack(t *testing.T) {
	s := ch2.NewArrayStack()
	testList(s, t)
}

func TestArrayQueue(t *testing.T) {
	s := ch2.NewArrayQueue()
	testQueue(s, t)
}

func TestArrayDeque(t *testing.T) {
	s := ch2.NewArrayDeque()
	testList(s, t)
}

func TestSLList(t *testing.T) {
	s := ch3.NewSLList()
	testQueue(s, t)
}

func TestDLList(t *testing.T) {
	s := ch3.NewDLList()
	testList(s, t)
}

func TestSEList(t *testing.T) {
	s := ch3.NewSEList()
	testList(s, t)
}

func TestSkiplistSSet(t *testing.T) {
	s := ch4.NewSkiplistSSet()
	testSSet(s, t)
}

func TestSkiplistList(t *testing.T) {
	s := ch4.NewSkiplistList()
	testList(s, t)
}

func TestChainedHashTable(t *testing.T) {
	s := ch5.NewChainedHashTable()
	testSSet(s, t)
}

func TestLinearHashTable(t *testing.T) {
	s := ch5.NewLinearHashTable()
	testSSet(s, t)
}

func TestBinaryTree(t *testing.T) {
	s := ch6.NewBinaryTree()
	testSSet(s, t)
}

func TestTreap(t *testing.T) {
	s := ch7.NewTreap()
	testSSet(s, t)
}

func TestScapegoatTree(t *testing.T) {
	s := ch8.NewScapegoatTree()
	testSSet(s, t)
}

func TestMergeSort(t *testing.T) {
	testSort(ch11.MergeSort, t)
}

func TestBinaryTrie(t *testing.T) {
	s := ch13.NewBinaryTrie()
	testSSet(s, t)
}

func TestXFastTrie(t *testing.T) {
	s := ch13.NewXFastTrie()
	testSSet(s, t)
}

func testList(s ch1.List, t *testing.T) {
	if !reflect.DeepEqual(s.GetAll(), []interface{}{}) {
		t.Fatalf("failed init %#v", s.GetAll())
	}
	t.Log("Add")
	s.Add(0, 40)
	s.Add(1, 50)
	s.Add(2, 60)
	s.Add(3, 70)
	s.Add(4, 80)
	s.Add(0, 3)
	s.Add(0, 2)
	s.Add(0, 1)
	s.Add(0, 0)
	if !reflect.DeepEqual(s.GetAll(), []interface{}{0, 1, 2, 3, 40, 50, 60, 70, 80}) {
		t.Fatalf("failed Add %#v", s.GetAll())
	}
	t.Log("Remove")
	s.Remove(8)
	s.Remove(0)
	if !reflect.DeepEqual(s.GetAll(), []interface{}{1, 2, 3, 40, 50, 60, 70}) {
		t.Fatalf("failed Remove %#v", s.GetAll())
	}
}

func testQueue(q ch1.Queue, t *testing.T) {
	if !reflect.DeepEqual(q.GetAll(), []interface{}{}) {
		t.Fatalf("failed init %#v", q.GetAll())
	}
	t.Log("Add")
	q.Add(0)
	q.Add(1)
	q.Add(2)
	q.Add(3)
	q.Add(4)
	if !reflect.DeepEqual(q.GetAll(), []interface{}{0, 1, 2, 3, 4}) {
		t.Fatalf("failed init %#v", q.GetAll())
	}
	t.Log("Remove")
}

func testSSet2(s ch1.SSet, t *testing.T) {
	num := 50000
	for i := 0; i < num; i++ {
		s.Add(element(i))
	}
	fmt.Println(num, "add finished")
}

func testSSet(s ch1.SSet, t *testing.T) {
	if !reflect.DeepEqual(s.GetAll(), []interface{}{}) {
		t.Fatalf("failed init %#v", s.GetAll())
	}
	t.Log("Add")
	s.Add(element(10))
	s.Add(element(20))
	s.Add(element(30))
	s.Add(element(40))
	slice := s.GetAll()
	sort.Slice(slice, func(i, j int) bool { return slice[i].(element) < slice[j].(element) })
	//sort.Slice(slice, func(i, j int) bool { return slice[i].(uint) < slice[j].(uint) })
	if !reflect.DeepEqual(slice, []interface{}{element(10), element(20), element(30), element(40)}) {
		//if !reflect.DeepEqual(slice, []interface{}{10, 20, 30, 40}) {
		t.Fatalf("failed Add %#v", slice)
	}
	t.Log("Find")
	if !(s.Find(element(20)).Compare(element(20)) == 0) {
		t.Fatalf("failed Find %#v", s.GetAll())
	}
	if !(s.Find(element(123)) == nil) {
		t.Fatalf("failed Find %#v", s.GetAll())
	}
	t.Log("Remove")
	s.Remove(element(10))
	s.Remove(element(30))
	s.Remove(element(40))
	s.Remove(element(20))
	if !reflect.DeepEqual(s.GetAll(), []interface{}{}) {
		t.Fatalf("failed Remove %#v", s.GetAll())
	}
}

func testSort(f sortfunc, t *testing.T) {
	s := []ch1.Comparable{}
	for i := 40; i > 0; i -= 10 {
		s = append(s, element(i))
	}
	//fmt.Printf("s=%v\n", s)
	f(s)
	//fmt.Printf("s=%v\n", s)
	if !reflect.DeepEqual(s, []ch1.Comparable{element(10), element(20), element(30), element(40)}) {
		//if !reflect.DeepEqual(slice, []interface{}{10, 20, 30, 40}) {
		t.Fatalf("failed Sort %#v", s)
	}
}

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
