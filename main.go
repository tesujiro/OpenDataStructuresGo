package main

import (
	"fmt"

	"github.com/tesujiro/OpenDataStructuresGo/ch2"
	"github.com/tesujiro/OpenDataStructuresGo/ch3"
)

type Printable interface {
	Print()
}

type List interface {
	Printable
	Len() int
	Cap() int
	Get(i int) interface{}
	Set(i int, v interface{})
	Resize()
	GetAll() []interface{}
	Add(i int, v interface{})
	Remove(i int)
}

type FIFO interface {
	Printable
	Add(x interface{})
	Remove() interface{}
}

type Stack interface {
	Push(x interface{})
	Pop() interface{}
	Print()
}

func main() {

	fmt.Println("ch2/ArrayStack")
	testAdd(ch2.NewArrayStack())

	fmt.Println("ch2/ArrayDeque")
	testAdd(ch2.NewArrayDeque())

	fmt.Println("ch3/SLList")
	testStack(ch3.NewSLList())
	testFIFO(ch3.NewSLList())
}

func testAdd(s List) {

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

func testStack(s Stack) {
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Print()
	fmt.Printf("Pop()=%v\n", s.Pop())
	fmt.Printf("Pop()=%v\n", s.Pop())
	fmt.Printf("Pop()=%v\n", s.Pop())
	s.Print()
}

func testFIFO(s FIFO) {
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Print()
	fmt.Printf("Remove()=%v\n", s.Remove())
	fmt.Printf("Remove()=%v\n", s.Remove())
	fmt.Printf("Remove()=%v\n", s.Remove())
	s.Print()
}
