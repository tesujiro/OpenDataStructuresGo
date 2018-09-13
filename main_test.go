package main

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/tesujiro/OpenDataStructuresGo/ch2"
)

func TestArrayStack(t *testing.T) {
	var s Sliceable
	t.Log("init")
	s = ch2.NewArrayStack()
	testSliceable(s, t)
}

func TestArrayDeque(t *testing.T) {
	var s Sliceable
	t.Log("init")
	s = ch2.NewArrayDeque()
	testSliceable(s, t)
}

func testSliceable(s Sliceable, t *testing.T) {
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
	var s Sliceable
	s = ch2.NewArrayStack()
	benchmarkSliceable_AddFirst(s, b)
}

func BenchmarkArrayStack_AddLast(b *testing.B) {
	var s Sliceable
	s = ch2.NewArrayStack()
	benchmarkSliceable_AddLast(s, b)
}

func BenchmarkArrayStack_AddRandom(b *testing.B) {
	var s Sliceable
	s = ch2.NewArrayStack()
	benchmarkSliceable_AddRandom(s, b)
}

func BenchmarkArrayDeque_AddFirst(b *testing.B) {
	var s Sliceable
	s = ch2.NewArrayDeque()
	benchmarkSliceable_AddFirst(s, b)
}

func BenchmarkArrayDeque_AddLast(b *testing.B) {
	var s Sliceable
	s = ch2.NewArrayDeque()
	benchmarkSliceable_AddLast(s, b)
}

func BenchmarkArrayDeque_AddRandom(b *testing.B) {
	var s Sliceable
	s = ch2.NewArrayDeque()
	benchmarkSliceable_AddRandom(s, b)
}

func benchmarkSliceable_AddFirst(s Sliceable, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(0, i)
	}
}

func benchmarkSliceable_AddLast(s Sliceable, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(i, i)
	}
}

func benchmarkSliceable_AddRandom(s Sliceable, b *testing.B) {

	rand.Seed(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(rand.Intn(s.Len()+1), i)
	}
}
