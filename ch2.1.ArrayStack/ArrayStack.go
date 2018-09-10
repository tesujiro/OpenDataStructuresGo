package main

import "fmt"

// ArrayStack is a slice of interface{}
type ArrayStack []interface{}

func newArrayStack() ArrayStack {
	return make(ArrayStack, 0)
}

func (as ArrayStack) Len() int {
	return len(as)
}

func (as ArrayStack) Cap() int {
	return cap(as)
}

func (as ArrayStack) Print() {
	fmt.Printf("ArrayStack(len:%v,cap:%v)=%#v\n", as.Len(), as.Cap(), as)
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

func (as ArrayStack) Add(i int, v interface{}) (ArrayStack, error) {
	if i < 0 || len(as) < i {
		return nil, fmt.Errorf("ArrayStack.Add: index out of range (i:%v)", i)
	}
	if i == len(as) {
		as = append(as, v)
	} else {
		as = append(as[:i+1], as[i:]...)
		as[i] = v
	}
	return as, nil
}

func (as ArrayStack) Remove(i int) (ArrayStack, error) {
	if i < 0 || len(as) < i {
		return nil, fmt.Errorf("ArrayStack.Add: index out of range (i:%v)", i)
	}
	if i == len(as) {
		as = as[:i]
	} else {
		as = append(as[:i], as[i+1:]...)
	}
	return as, nil
}

func main() {
	fmt.Println("ArrayStack")

	as := newArrayStack()
	as.Print()

	as, _ = as.Add(0, 10)
	as, _ = as.Add(0, 10)
	as.Print()
	as, _ = as.Add(1, 20)
	as, _ = as.Add(2, 30)
	as, _ = as.Add(3, 40)
	as, _ = as.Add(0, 5)
	as.Print()

	var i int
	var v interface{}
	i = 1
	v, _ = as.Get(i)
	i = 2
	v = 10
	as.Set(i, v)
	as.Print()

	as, _ = as.Remove(0)
	as.Print()
	as, _ = as.Remove(2)
	as.Print()
	as, _ = as.Remove(3)
	as.Print()
}
