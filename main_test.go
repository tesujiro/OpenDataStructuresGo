package main

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/tesujiro/OpenDataStructuresGo/ch1"
	"github.com/tesujiro/OpenDataStructuresGo/ch2"
	"github.com/tesujiro/OpenDataStructuresGo/ch3"
	"github.com/tesujiro/OpenDataStructuresGo/ch4"
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
	s.Print()
	if !reflect.DeepEqual(s.GetAll(), []interface{}{0, 1, 2, 3, 40, 50, 60, 70, 80}) {
		t.Fatalf("failed init %#v", s.GetAll())
	}
	t.Log("Remove")
	s.Remove(8)
	s.Remove(0)
	s.Print()
	if !reflect.DeepEqual(s.GetAll(), []interface{}{1, 2, 3, 40, 50, 60, 70}) {
		t.Fatalf("failed init %#v", s.GetAll())
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
	q.Print()
	if !reflect.DeepEqual(q.GetAll(), []interface{}{0, 1, 2, 3, 4}) {
		t.Fatalf("failed init %#v", q.GetAll())
	}
	t.Log("Remove")
}

func testSSet(s ch1.SSet, t *testing.T) {
	if !reflect.DeepEqual(s.GetAll(), []interface{}{}) {
		t.Fatalf("failed init %#v", s.GetAll())
	}
	t.Log("Add")
	s.Add(10)
	s.Add(20)
	s.Add(30)
	if !reflect.DeepEqual(s.GetAll(), []interface{}{10, 20, 30}) {
		t.Fatalf("failed Add %#v", s.GetAll())
	}
	t.Log("Find")
	if !(s.Find(20).(int) == 20) {
		t.Fatalf("failed Find %#v", s.GetAll())
	}
	if !(s.Find(123) == nil) {
		t.Fatalf("failed Find %#v", s.GetAll())
	}
	t.Log("Remove")
	s.Remove(10)
	s.Remove(30)
	s.Remove(20)
	if !reflect.DeepEqual(s.GetAll(), []interface{}{}) {
		t.Fatalf("failed Remove %#v", s.GetAll())
	}
}

func BenchmarkArrayStack_AddFirst(b *testing.B) {
	s := ch2.NewArrayStack()
	benchmarkList_AddFirst(s, b)
}

func BenchmarkArrayStack_AddLast(b *testing.B) {
	s := ch2.NewArrayStack()
	benchmarkList_AddLast(s, b)
}

func BenchmarkArrayStack_AddRandom(b *testing.B) {
	s := ch2.NewArrayStack()
	benchmarkList_AddRandom(s, b)
}

func BenchmarkArrayDeque_AddFirst(b *testing.B) {
	s := ch2.NewArrayDeque()
	benchmarkList_AddFirst(s, b)
}

func BenchmarkArrayDeque_AddLast(b *testing.B) {
	s := ch2.NewArrayDeque()
	benchmarkList_AddLast(s, b)
}

func BenchmarkArrayDeque_AddRandom(b *testing.B) {
	s := ch2.NewArrayDeque()
	benchmarkList_AddRandom(s, b)
}

func BenchmarkDLList_AddFirst(b *testing.B) {
	s := ch3.NewDLList()
	benchmarkList_AddFirst(s, b)
}

func BenchmarkDLList_AddLast(b *testing.B) {
	s := ch3.NewDLList()
	benchmarkList_AddLast(s, b)
}

func BenchmarkDLList_AddRandom(b *testing.B) {
	s := ch3.NewDLList()
	benchmarkList_AddRandom(s, b)
}

func BenchmarkSEList_AddFirst(b *testing.B) {
	s := ch3.NewSEList()
	benchmarkList_AddFirst(s, b)
}

func BenchmarkSEList_AddLast(b *testing.B) {
	s := ch3.NewSEList()
	benchmarkList_AddLast(s, b)
}

func BenchmarkSEList_AddRandom(b *testing.B) {
	s := ch3.NewSEList()
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
