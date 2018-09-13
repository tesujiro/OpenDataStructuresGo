package main

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/tesujiro/OpenDataStructuresGo/ch2"
)

func TestArrayStack(t *testing.T) {
	var s Slicable
	t.Log("init")
	s = ch2.NewArrayStack()
	testSlicable(s, t)
}

func TestArrayDeque(t *testing.T) {
	var s Slicable
	t.Log("init")
	s = ch2.NewArrayDeque()
	testSlicable(s, t)
}

func testSlicable(s Slicable, t *testing.T) {
	if !reflect.DeepEqual(s.Slice(), []interface{}{}) {
		t.Fatalf("failed init %#v", s)
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
	if !reflect.DeepEqual(s.Slice(), []interface{}{0, 1, 2, 3, 40, 50, 60, 70, 80}) {
		t.Fatalf("failed init %#v", s.Slice())
	}
}

func BenchmarkArrayStack_AddFirst(b *testing.B) {
	var s Slicable
	s = ch2.NewArrayStack()
	benchmarkSlicable_AddFirst(s, b)
}

func BenchmarkArrayStack_AddLast(b *testing.B) {
	var s Slicable
	s = ch2.NewArrayStack()
	benchmarkSlicable_AddLast(s, b)
}

func BenchmarkArrayStack_AddRandom(b *testing.B) {
	var s Slicable
	s = ch2.NewArrayStack()
	benchmarkSlicable_AddRandom(s, b)
}

func BenchmarkArrayDeque_AddFirst(b *testing.B) {
	var s Slicable
	s = ch2.NewArrayDeque()
	benchmarkSlicable_AddFirst(s, b)
}

func BenchmarkArrayDeque_AddLast(b *testing.B) {
	var s Slicable
	s = ch2.NewArrayDeque()
	benchmarkSlicable_AddLast(s, b)
}

func BenchmarkArrayDeque_AddRandom(b *testing.B) {
	var s Slicable
	s = ch2.NewArrayDeque()
	benchmarkSlicable_AddRandom(s, b)
}

func benchmarkSlicable_AddFirst(s Slicable, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(0, i)
	}
}

func benchmarkSlicable_AddLast(s Slicable, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(i, i)
	}
}

func benchmarkSlicable_AddRandom(s Slicable, b *testing.B) {

	rand.Seed(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(rand.Intn(s.Len()+1), i)
	}
}
