package main

import "fmt"

type Slicable interface {
	Len() int
	Cap() int
	Get(i int) interface{}
	Set(i int, v interface{})
	Resize()
	Print()
	Slice() []interface{}
	Add(i int, v interface{})
	Remove(i int)
}

func main() {

	var s Slicable

	fmt.Println("ArrayStack")
	s = newArrayStack()
	addTest(s)

	fmt.Println("ArrayDeque")
	s = newArrayDeque()
	addTest(s)
}

func addTest(s Slicable) {

	s.Print()

	s.Add(0, 10)
	//s.Add(0, 10)
	s.Add(1, 40)
	s.Add(2, 50)
	s.Add(3, 60)
	s.Add(0, 3)
	s.Add(0, 2)
	s.Add(0, 1)
	s.Add(0, 0)

	var i int
	var v interface{}
	i = 1
	v = s.Get(i)
	i = 2
	v = 10
	s.Set(i, v)
	s.Print()

	s.Remove(0)
	s.Print()
	s.Remove(2)
	s.Print()
	s.Remove(3)
	s.Print()
}

/*
func ad() {
	fmt.Println("ArrayDeque")

	as := newArrayDeque()
	as.Print()

	as.Add(0, 40)
	as.Add(1, 50)
	as.Add(2, 60)
	as.Add(3, 70)
	as.Add(4, 80)
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
*/
