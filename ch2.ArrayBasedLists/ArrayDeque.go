package main

import (
	"fmt"
)

const initSizeDQ = 8

// ArrayDeque is a slice of interface{}
type ArrayDeque struct {
	array []interface{}
	start int
	len   int
}

func newArrayDeque() *ArrayDeque {
	ar := make([]interface{}, initSizeDQ)
	return &ArrayDeque{
		array: ar,
		start: 0,
		len:   0,
	}
}

func (as *ArrayDeque) Len() int {
	return as.len
}

func (as *ArrayDeque) Cap() int {
	return len(as.array)
}

func (as *ArrayDeque) Get(i int) interface{} {
	i = i % as.len
	return as.array[(as.start+i)%as.Cap()]
}

func (as *ArrayDeque) Set(i int, v interface{}) {
	i = i % as.len
	as.array[(as.start+i)%as.Cap()] = v
}

func (as *ArrayDeque) Resize() {
	ar := make([]interface{}, len(as.array)*2)
	for i := 0; i < as.len; i++ {
		ar[i] = as.Get(i)
	}
	as.start = 0
	as.array = ar
}

func (as *ArrayDeque) Print() {
	fmt.Printf("ArrayDeque(len:%v,cap:%v,start:%v)=%v\n", as.Len(), as.Cap(), as.start, as.Slice())
}

func (as *ArrayDeque) Slice() []interface{} {
	slice := []interface{}{}
	for i := 0; i < as.len; i++ {
		slice = append(slice, as.Get(i))
	}
	return slice
}

func (as *ArrayDeque) Add(i int, v interface{}) {
	if as.len+1 > len(as.array) {
		as.Resize()
	}

	if as.len == 0 {
		as.start = 0
		as.len = 1
		as.Set(0, v)
		return
	}

	if i > as.len {
		i = i % as.len
	}
	as.len += 1

	//fmt.Printf("i=%v as.len/2=%v\n", i, as.len/2)
	if i < as.len/2 {
		// shift left
		//fmt.Println("shift left")
		as.start = (as.start - 1 + as.Cap()) % as.Cap()
		for j := 0; j < i; j++ {
			as.Set(j, as.Get(j+1))
		}
	} else {
		// shift right
		//fmt.Println("shift right")
		for j := as.len - 1; i < j; j-- {
			as.Set(j, as.Get(j-1))
		}
	}
	as.Set(i, v)
	//as.Print()
}

func (as *ArrayDeque) Remove(i int) {
	i = i % as.len
	//TODO if i< as.len/2
	for j := i; j < as.len-1; j++ {
		as.Set(j, as.Get(j+1))
	}
	as.len -= 1
}
