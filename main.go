package main

import (
	"fmt"

	"github.com/tesujiro/OpenDataStructuresGo/ch2"
	"github.com/tesujiro/OpenDataStructuresGo/ch3"
)

type Printable interface {
	GetAll() []interface{}
	Print()
}

type List interface {
	Printable
	Len() int
	Get(i int) interface{}
	Set(i int, v interface{}) interface{}
	Add(i int, v interface{})
	Remove(i int) interface{}
}

type Queue interface {
	Printable
	Add(x interface{}) bool
	Remove() interface{}
}

type Stack interface {
	Push(x interface{})
	Pop() interface{}
	Print()
}

func main() {

	fmt.Println("ch2/ArrayStack")
	checkList(ch2.NewArrayStack())

	fmt.Println("ch2/ArrayQueue")
	checkQueue(ch2.NewArrayQueue())

	fmt.Println("ch2/ArrayDeque")
	checkList(ch2.NewArrayDeque())

	fmt.Println("ch3/SLList")
	checkStack(ch3.NewSLList())
	checkQueue(ch3.NewSLList())

	fmt.Println("ch3/DLList")
	checkList(ch3.NewDLList())

	fmt.Println("ch3/SEList")
	checkList(ch3.NewSEList())
}

func checkList(s List) {

	s.Print()

	s.Add(0, 10)
	s.Print()
	//s.Add(0, 10)
	s.Add(0, 40)
	s.Print()

	/*
		x := s.Remove(0)
		fmt.Println("Removed:", x)
		s.Print()
	*/

	s.Add(2, 50)
	s.Print()
	s.Add(3, 60)
	s.Print()
	s.Add(0, 3)
	s.Print()
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

func checkStack(s Stack) {
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Print()
	fmt.Printf("Pop()=%v\n", s.Pop())
	fmt.Printf("Pop()=%v\n", s.Pop())
	fmt.Printf("Pop()=%v\n", s.Pop())
	s.Print()
}

func checkQueue(s Queue) {
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Print()
	fmt.Printf("Remove()=%v\n", s.Remove())
	fmt.Printf("Remove()=%v\n", s.Remove())
	fmt.Printf("Remove()=%v\n", s.Remove())
	s.Print()
}
