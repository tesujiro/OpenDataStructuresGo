package ch2

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

func NewArrayDeque() *ArrayDeque {
	ar := make([]interface{}, initSizeDQ)
	return &ArrayDeque{
		array: ar,
		start: 0,
		len:   0,
	}
}

func (as *ArrayDeque) cap() int {
	return len(as.array)
}

func (as *ArrayDeque) Len() int {
	return as.len
}

func (as *ArrayDeque) GetAll() []interface{} {
	slice := []interface{}{}
	for i := 0; i < as.len; i++ {
		slice = append(slice, as.Get(i))
	}
	return slice
}

func (as *ArrayDeque) Print() {
	fmt.Printf("ArrayDeque(len:%v,cap:%v,start:%v)=%v\n", as.Len(), as.cap(), as.start, as.GetAll())
}

func (as *ArrayDeque) Get(i int) interface{} {
	i = i % as.len
	return as.array[(as.start+i)%as.cap()]
}

func (as *ArrayDeque) Set(i int, v interface{}) interface{} {
	i = i % as.len
	y := as.array[(as.start+i)%as.cap()]
	as.array[(as.start+i)%as.cap()] = v
	return y
}

func (as *ArrayDeque) resize() {
	ar := make([]interface{}, len(as.array)*2)
	for i := 0; i < as.len; i++ {
		ar[i] = as.Get(i)
	}
	as.start = 0
	as.array = ar
}

func (as *ArrayDeque) Add(i int, v interface{}) {
	if as.len+1 > len(as.array) {
		as.resize()
	}

	if as.len == 0 {
		as.start = 0
		as.len = 1
		as.Set(0, v)
		return
	}

	i = i % (as.len + 1)
	as.len += 1

	if i < as.len/2 {
		// shift left
		as.start = (as.start - 1 + as.cap()) % as.cap()
		for j := 0; j < i; j++ {
			as.Set(j, as.Get(j+1))
		}
	} else {
		// shift right
		for j := as.len - 1; i < j; j-- {
			as.Set(j, as.Get(j-1))
		}
	}
	as.Set(i, v)
}

func (as *ArrayDeque) Remove(i int) interface{} {
	i = i % as.len
	x := as.Get(i)
	//TODO if i< as.len/2
	for j := i; j < as.len-1; j++ {
		as.Set(j, as.Get(j+1))
	}
	as.len -= 1
	return x
}
