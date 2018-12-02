package ch02

import (
	"fmt"
)

const initSize = 1

type ArrayStack struct {
	a []interface{}
	n int
}

func NewArrayStack() *ArrayStack {
	ar := make([]interface{}, initSize)
	return &ArrayStack{
		a: ar,
		n: 0,
	}
}

func (as *ArrayStack) cap() int {
	return cap(as.a)
}

func (as *ArrayStack) Size() int {
	return as.n
}

func (as *ArrayStack) GetAll() []interface{} {
	return as.a[:as.n]
}

func (as *ArrayStack) Print() {
	fmt.Printf("ArrayStack(n:%v,cap:%v)=%v\n", as.Size(), as.cap(), as.GetAll())
}

func (as *ArrayStack) Get(i int) interface{} {
	//i = i % as.n
	return as.a[i]
}

func (as *ArrayStack) Set(i int, v interface{}) interface{} {
	//i = i % as.n
	y := as.a[i]
	as.a[i] = v
	return y
}

func (as *ArrayStack) resize() {
	var new []interface{}
	if as.n > 0 {
		new = make([]interface{}, as.n*2)
	} else {
		new = make([]interface{}, 1)
	}
	for i := 0; i < as.n; i++ {
		new[i] = as.a[i]
	}
	as.a = new
}

func (as *ArrayStack) Add(i int, v interface{}) {
	if as.n > 0 {
		i = i % (as.n + 1)
	}
	if as.n+1 > as.cap() {
		//fmt.Println("before ArrayStack#resize() as.n=", as.n, " as.cap()=", as.cap())
		as.resize()
		//fmt.Println("after  ArrayStack#resize() as.n=", as.n, " as.cap()=", as.cap())
	}
	for j := as.n; j > i; j-- {
		as.a[j] = as.a[j-1]
	}
	as.a[i] = v
	as.n++
}

func (as *ArrayStack) Remove(i int) interface{} {
	i = i % as.n
	x := as.a[i]
	for j := i; j < as.n-1; j++ {
		as.a[j] = as.a[j+1]
	}
	as.n--
	if as.cap() >= 3*as.n {
		as.resize()
	}
	return x
}
