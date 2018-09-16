package ch2

import (
	"fmt"
)

const initSize = 8

type ArrayStack struct {
	array []interface{}
	len   int
}

func NewArrayStack() *ArrayStack {
	ar := make([]interface{}, initSize)
	return &ArrayStack{
		array: ar,
		len:   0,
	}
}

func (as *ArrayStack) cap() int {
	return cap(as.array)
}

func (as *ArrayStack) Len() int {
	return as.len
}

func (as *ArrayStack) GetAll() []interface{} {
	return as.array[:as.len]
}

func (as *ArrayStack) Print() {
	fmt.Printf("ArrayStack(len:%v,cap:%v)=%v\n", as.Len(), as.cap(), as.GetAll())
}

func (as *ArrayStack) Get(i int) interface{} {
	i = i % as.len
	return as.array[i]
}

func (as *ArrayStack) Set(i int, v interface{}) interface{} {
	i = i % as.len
	y := as.array[i]
	as.array[i] = v
	return y
}

func (as *ArrayStack) resize() {
	var new []interface{}
	if as.len > 1 {
		new = make([]interface{}, as.len*2)
	} else {
		new = make([]interface{}, 1)
	}
	for i := 0; i < as.len; i++ {
		new[i] = as.array[i]
	}
	as.array = new
}

func (as *ArrayStack) Add(i int, v interface{}) {
	if as.len > 0 {
		i = i % (as.len + 1)
	}
	if as.len+1 > as.cap() {
		as.resize()
	}
	for j := as.len; j > i; j-- {
		as.array[j] = as.array[j-1]
	}
	as.array[i] = v
	as.len += 1
}

func (as *ArrayStack) Remove(i int) interface{} {
	i = i % as.len
	x := as.array[i]
	for j := i; j < as.len-1; j++ {
		as.array[j] = as.array[j+1]
	}
	as.len -= 1
	if as.cap() >= 3*as.len {
		as.resize()
	}
	return x
}
