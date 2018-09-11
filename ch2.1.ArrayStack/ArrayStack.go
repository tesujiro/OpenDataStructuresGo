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
	ar := [initSize]interface{}{}
	return &ArrayStack{
		array: ar[:],
		slice: ar[0:0],
	}
}

func (as *ArrayStack) Len() int {
	return len(as.slice)
}

func (as *ArrayStack) Cap() int {
	return cap(as.array)
}

func (as *ArrayStack) Print() {
	fmt.Printf("ArrayStack(len:%v,cap:%v)=%#v\n", as.Len(), as.Cap(), as.slice)
}

func (as *ArrayStack) Get(i int) (interface{}, error) {
	if i < 0 || len(as.slice) <= i {
		return nil, fmt.Errorf("ArrayStack.Get: index out of range (i:%v)", i)
	}
	return as.slice[i], nil
}

func (as *ArrayStack) Set(i int, v interface{}) error {
	if i < 0 || len(as.slice) <= i {
		return fmt.Errorf("ArrayStack.Set: index out of range (i:%v)", i)
	}
	as.slice[i] = v
	return nil
}

func (as *ArrayStack) Add(i int, v interface{}) error {
	if i < 0 || len(as.slice) < i {
		return fmt.Errorf("ArrayStack.Add: index out of range (i:%v)", i)
	}
	if len(as.slice)+1 > len(as.array) {
		return fmt.Errorf("ArrayStack.Add: cannot extend array, please resize at first")
	}

	as.slice = as.array[:len(as.slice)+1]
	for j := len(as.slice) - 1; i < j; j-- {
		as.slice[j] = as.slice[j-1]
	}
	as.slice[i] = v
	return nil
}

func (as *ArrayStack) Remove(i int) error {
	if i < 0 || len(as.slice) < i {
		return fmt.Errorf("ArrayStack.Add: index out of range (i:%v)", i)
	}
	for j := i; j < len(as.slice)-1; j++ {
		as.slice[j] = as.slice[j+1]
	}
	as.slice[len(as.slice)-1] = nil
	as.slice = as.array[:len(as.slice)-1]
	return nil
}

func main() {
	fmt.Println("ArrayStack")

	as := newArrayStack()
	as.Print()

	as.Add(0, 10)
	as.Add(0, 10)
	as.Print()
	as.Add(1, 60)
	as.Add(2, 50)
	as.Add(3, 40)
	as.Add(0, 3)
	as.Add(0, 2)
	as.Add(0, 1)
	as.Add(0, 0)
	as.Print()

	var i int
	var v interface{}
	i = 1
	v, _ = as.Get(i)
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
