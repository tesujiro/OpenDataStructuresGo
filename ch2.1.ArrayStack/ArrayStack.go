package main

import "fmt"

// ArrayStack is a slice of interface{}
type ArrayStack []interface{}

func newArrayStack(c int) ArrayStack {
	return make(ArrayStack, 0, c)
}

func (as ArrayStack) Len() int {
	return len(as)
}

func (as ArrayStack) Cap() int {
	return cap(as)
}

func (as ArrayStack) Get(i int) (interface{}, error) {
	if i < 0 || len(as) <= i {
		return nil, fmt.Errorf("ArrayStack.Get: index out of range (i:%v)", i)
	}
	return as[i], nil
}

func (as ArrayStack) Set(i int, v interface{}) error {
	if i < 0 || len(as) <= i {
		return fmt.Errorf("ArrayStack.Set: index out of range (i:%v)", i)
	}
	as[i] = v
	return nil
}

func (as ArrayStack) Add(i int, v interface{}) error {
	if i < 0 || len(as) < i {
		return fmt.Errorf("ArrayStack.Add: index out of range (i:%v)", i)
	}
	/* TODO: check size and resize */
	// TODO: NG
	as = append(as, as[len(as)-1])
	as = append(as[:i+1], as[i:]...)
	as[i] = v
	return nil
}

func main() {
	fmt.Println("ArrayStack")

	as := newArrayStack(1024)
	fmt.Printf("ArrayStack(len:%v,cap:%v)=%#v\n", as.Len(), as.Cap(), as)

	as = append(as, 1)
	as = append(as, 2)
	as = append(as, 3)
	fmt.Printf("ArrayStack(len:%v,cap:%v)=%#v\n", as.Len(), as.Cap(), as)

	as.Add(1, 100)
	fmt.Printf("ArrayStack(len:%v,cap:%v)=%#v\n", as.Len(), as.Cap(), as)

	var i int
	var v interface{}
	i = 1
	v, _ = as.Get(i)
	fmt.Printf("ArrayStack[%v]:%#v\n", i, v)

	i = 2
	v = 10
	as.Set(i, v)
	fmt.Printf("ArrayStack(len:%v,cap:%v)=%#v\n", as.Len(), as.Cap(), as)
}
