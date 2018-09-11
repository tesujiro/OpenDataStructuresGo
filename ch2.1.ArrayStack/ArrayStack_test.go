package main

import (
	"reflect"
	"testing"
)

func TestArrayStack(t *testing.T) {
	t.Log("init")
	as := newArrayStack()
	as.Print()
	if !reflect.DeepEqual(as.slice, []interface{}{}) {
		t.Fatalf("failed init %#v", as)
	}
	t.Log("add")
	as.Add(0, 40)
	as.Add(1, 50)
	as.Add(2, 60)
	as.Add(3, 70)
	as.Add(4, 80)
	as.Add(0, 3)
	as.Add(0, 2)
	as.Add(0, 1)
	as.Add(0, 0)
	as.Print()
	if !reflect.DeepEqual(as.slice, []interface{}{0, 1, 2, 3, 40, 50, 60, 70, 80}) {
		t.Fatalf("failed init %#v", as)
	}
}

func BenchmarkAddLast(b *testing.B) {
	as := newArrayStack()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := as.Add(i, i)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkAddFirst(b *testing.B) {
	as := newArrayStack()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := as.Add(0, i)
		if err != nil {
			b.Error(err)
		}
	}
}
