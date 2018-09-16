package ch2

import (
	"fmt"
)

const initSizeDQ = 8

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

func (as *ArrayDeque) Set(i int, x interface{}) interface{} {
	i = i % as.len
	y := as.array[(as.start+i)%as.cap()]
	as.array[(as.start+i)%as.cap()] = x
	return y
}

func (as *ArrayDeque) resize() {
	var new []interface{}
	if as.len > 1 {
		new = make([]interface{}, as.len*2)
	} else {
		new = make([]interface{}, 1)
	}
	for i := 0; i < as.len; i++ {
		new[i] = as.Get(i)
	}
	as.start = 0
	as.array = new
}

func (as *ArrayDeque) Add(i int, x interface{}) {
	i = i % (as.len + 1)
	if as.len+1 > as.cap() {
		as.resize()
	}
	if i < as.len/2 {
		if as.start == 0 {
			as.start = len(as.array) - 1
		} else {
			as.start = as.start - 1
		}
		for k := 0; k <= i-1; k++ {
			as.array[(as.start+k)%as.cap()] = as.array[(as.start+k+1)%as.cap()]
		}
	} else {
		for k := as.len; k > i; k-- {
			as.array[(as.start+k)%as.cap()] = as.array[(as.start+k-1)%as.cap()]
		}
	}
	as.array[(as.start+i)%as.cap()] = x
	as.len += 1
}

func (as *ArrayDeque) Remove(i int) interface{} {
	i = i % as.len

	x := as.array[(as.start+i)%as.cap()]
	if i < as.len/2 {
		for k := i; k > 0; k-- {
			as.array[(as.start+k)%as.cap()] = as.array[(as.start+k-1)%as.cap()]
		}
		as.start = (as.start + 1) % as.cap()
	} else {
		for k := i; k <= as.len-1; k++ {
			as.array[(as.start+k)%as.cap()] = as.array[(as.start+k+1)%as.cap()]
		}
	}
	/*
		x := as.Get(i)
		//TODO if i< as.len/2
		for j := i; j < as.len-1; j++ {
			as.Set(j, as.Get(j+1))
		}
	*/
	as.len -= 1
	if 3*as.len < as.cap() {
		as.resize()
	}
	return x
}
