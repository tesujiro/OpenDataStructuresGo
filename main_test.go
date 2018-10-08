package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"

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

func TestDualArrayDeque(t *testing.T) {
	s := ch2.NewDualArrayDeque()
	testList(s, t)
}

func TestRootishArrayStack(t *testing.T) {
	s := ch2.NewRootishArrayStack()
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

func TestQuickSort(t *testing.T) {
	testSort(ch11.QuickSort, t)
}

func TestCountingSort(t *testing.T) {
	testCSort(ch11.CountingSort, t)
}

func TestRadixSort(t *testing.T) {
	testCSort(ch11.RadixSort, t)
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
	result := []interface{}{}
	for i := 10; i < 50; i += 10 {
		s.Add(element(i))
		result = append(result, element(i))
	}
	/*
		s.Add(element(10))
		s.Add(element(20))
		s.Add(element(30))
		s.Add(element(40))
	*/
	slice := s.GetAll()
	sort.Slice(slice, func(i, j int) bool { return slice[i].(element) < slice[j].(element) })
	//sort.Slice(slice, func(i, j int) bool { return slice[i].(uint) < slice[j].(uint) })
	if !reflect.DeepEqual(slice, result) {
		//if !reflect.DeepEqual(slice, []interface{}{element(10), element(20), element(30), element(40)}) {
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

func testSort(f ch1.SortFunc, t *testing.T) {
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

func testCSort(f ch1.CountingSortFunc, t *testing.T) {
	k := 256 * 256
	s := []int{}
	rand.Seed(0)
	for i := 0; i < 10; i++ {
		s = append(s, rand.Intn(k))
	}
	//fmt.Printf("s=%v\n", s)
	f(&s, k)
	//fmt.Printf("s=%v\n", s)
	if !reflect.DeepEqual(s, []int{1042, 6395, 12282, 24569, 29960, 45536, 48426, 54287, 57040, 60549}) {
		t.Fatalf("failed CountingSort %#v", s)
	}
}
