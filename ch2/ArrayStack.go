package ch2

import (
	"fmt"
)

const initSize = 8

// ArrayStack is a slice of interface{}
type ArrayStack struct {
	array []interface{}
	slice []interface{}
}

func NewArrayStack() *ArrayStack {
	ar := make([]interface{}, initSize)
	return &ArrayStack{
		array: ar,
		slice: ar[0:0],
	}
}

func (as *ArrayStack) Len() int {
	return len(as.slice)
}

func (as *ArrayStack) Cap() int {
	return cap(as.array)
}

func (as *ArrayStack) Get(i int) interface{} {
	i = i % len(as.slice)
	return as.slice[i]
}

func (as *ArrayStack) Set(i int, v interface{}) {
	i = i % len(as.slice)
	as.slice[i] = v
}

func (as *ArrayStack) Resize() {
	ar := make([]interface{}, len(as.array)*2)
	for i := 0; i < len(as.array); i++ {
		ar[i] = as.array[i]
	}
	as.array = ar
	as.slice = ar[:len(as.slice)]
}

func (as *ArrayStack) Print() {
	fmt.Printf("ArrayStack(len:%v,cap:%v)=%v\n", as.Len(), as.Cap(), as.GetAll())
}

func (as *ArrayStack) GetAll() []interface{} {
	return as.slice
}

func (as *ArrayStack) Add(i int, v interface{}) {
	if len(as.slice) > 0 {
		i = i % (len(as.slice) + 1)
	}
	if len(as.slice)+1 > len(as.array) {
		as.Resize()
	}

	as.slice = as.array[:len(as.slice)+1]
	for j := len(as.slice) - 1; i < j; j-- {
		as.slice[j] = as.slice[j-1]
	}
	as.slice[i] = v
	//as.Print()
}

func (as *ArrayStack) Remove(i int) {
	i = i % len(as.slice)
	for j := i; j < len(as.slice)-1; j++ {
		as.slice[j] = as.slice[j+1]
	}
	as.slice[len(as.slice)-1] = nil
	as.slice = as.array[:len(as.slice)-1]
}
