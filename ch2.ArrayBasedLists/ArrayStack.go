package main

import (
	"fmt"
)

const initSize = 8

// ArrayStack is a slice of interface{}
type ArrayStack struct {
	array []interface{}
	slice []interface{}
}

func newArrayStack() *ArrayStack {
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
	fmt.Printf("ArrayStack(len:%v,cap:%v)=%#v\n", as.Len(), as.Cap(), as.slice)
}

func (as *ArrayStack) Slice() []interface{} {
	return as.slice
}

func (as *ArrayStack) Add(i int, v interface{}) {
	if len(as.slice) > 0 && i > len(as.slice) {
		i = i % len(as.slice)
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

func main() {
	fmt.Println("ArrayStack")

	as := newArrayStack()
	as.Print()

	as.Add(0, 10)
	//as.Add(0, 10)
	as.Add(1, 40)
	as.Add(2, 50)
	as.Add(3, 60)
	as.Add(0, 3)
	as.Add(0, 2)
	as.Add(0, 1)
	as.Add(0, 0)

	var i int
	var v interface{}
	i = 1
	v = as.Get(i)
	i = 2
	v = 10
	as.Set(i, v)
	as.Print()

	as.Remove(0)
	as.Print()
	as.Remove(2)
	as.Print()
	as.Remove(3)
	as.Print()
}
