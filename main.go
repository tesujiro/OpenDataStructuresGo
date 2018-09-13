package main

import (
	"fmt"

	"github.com/tesujiro/OpenDataStructuresGo/ch2"
	"github.com/tesujiro/OpenDataStructuresGo/ch3"
)

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

	fmt.Println("ch2/ArrayStack")
	s = ch2.NewArrayStack()
	addTest(s)

	fmt.Println("ch2/ArrayDeque")
	s = ch2.NewArrayDeque()
	addTest(s)

	fmt.Println("ch2/ArrayDeque")
	SLList()
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

func SLList() {
	fmt.Println("SLList")
	sll := ch3.NewSLList()
	sll.Print()
	sll.Push(1)
	sll.Push(2)
	sll.Push(3)
	sll.Print()
	fmt.Printf("Pop()=%v\n", sll.Pop())
	sll.Print()
}
