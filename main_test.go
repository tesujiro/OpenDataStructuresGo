package main

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/tesujiro/OpenDataStructuresGo/ch2"
	"github.com/tesujiro/OpenDataStructuresGo/ch3"
)

func TestArrayStack(t *testing.T) {
	s := ch2.NewArrayStack()
	testList(s, t)
}

func TestArrayDeque(t *testing.T) {
	s := ch2.NewArrayDeque()
	testList(s, t)
}

func TestDLList(t *testing.T) {
	s := ch3.NewDLList()
	testList(s, t)
}

func testList(s List, t *testing.T) {
	if !reflect.DeepEqual(s.GetAll(), []interface{}{}) {
		t.Fatalf("failed init %#v", s.GetAll())
	}
	t.Log("add")
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

func benchmarkList_AddFirst(s List, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(0, i)
	}
}

func benchmarkList_AddLast(s List, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(i, i)
	}
}

func benchmarkList_AddRandom(s List, b *testing.B) {

	rand.Seed(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(rand.Intn(s.Len()+1), i)
	}
}
