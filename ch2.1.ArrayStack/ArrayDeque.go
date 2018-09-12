package main

import (
	"fmt"
)

const initSize = 8

// ArrayDeque is a slice of interface{}
type ArrayDeque struct {
	array []interface{}
	//slice []interface{}
	start int
	len   int
}

func newArrayDeque() *ArrayDeque {
	ar := make([]interface{}, initSize)
	return &ArrayDeque{
		array: ar,
		//slice: ar[0:0],
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

func (as *ArrayDeque) Resize() {
	ar := make([]interface{}, len(as.array)*2)
	for i := 0; i < as.len; i++ {
		ar[i] = as.array[(as.start+i)%as.Cap()]
	}
	as.array = ar
	//as.slice = ar[:as.len]
}

func (as *ArrayDeque) Print() {
	fmt.Printf("ArrayDeque(len:%v,cap:%v,start:%v)=%#v\n", as.Len(), as.Cap(), as.start, as.array)
}

func (as *ArrayDeque) Get(i int) (interface{}, error) {
	return as.array[(as.start+i%as.len)%as.Cap()], nil
}

func (as *ArrayDeque) Set(i int, v interface{}) error {
	as.array[(as.start+i%as.len)%as.Cap()] = v
	return nil
}

func (as *ArrayDeque) Add(i int, v interface{}) error {
	if as.len+1 > len(as.array) {
		as.Resize()
	}

	if as.len == 0 {
		as.start = 0
		as.len = 1
		as.array[as.start] = v
		return nil
	}

	i = i % as.len
	as.len += 1

	if i < as.len/2 {
		fmt.Println("shift left")
		// shift left
		as.start = (as.start - 1 + as.Cap()) % as.Cap()
		//for j := as.len - 1; i < j; j-- {
		for j := 0; j < i; j++ {
			as.array[(as.start+j)%as.Cap()] = as.array[as.start+j+1]
		}
	} else {
		fmt.Println("shift right")
		// shift right
		for j := as.len - 1; i < j; j-- {
			as.array[(as.start+j)%as.Cap()] = as.array[(as.start+j-1)%as.Cap()]
		}
	}
	fmt.Printf("i=%v v=%v start=%v len=%v\n", i, v, as.start, as.len)
	as.array[(as.start+i)%as.Cap()] = v
	return nil
}

func (as *ArrayDeque) Remove(i int) error {
	if i < 0 || as.len < i {
		return fmt.Errorf("ArrayDeque.Add: index out of range (i:%v)", i)
	}
	for j := i; j < as.len-1; j++ {
		as.array[as.start+j] = as.array[as.start+j+1]
	}
	as.array[as.start+as.len-1] = nil
	//as.slice = as.array[:as.len-1]
	as.len -= 1
	return nil
}

func main() {
	fmt.Println("ArrayDeque")

	as := newArrayDeque()
	as.Print()

	as.Add(0, 10)
	as.Print()
	as.Add(0, 10)
	as.Print()
	as.Add(1, 40)
	as.Add(2, 50)
	as.Add(3, 60)
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
