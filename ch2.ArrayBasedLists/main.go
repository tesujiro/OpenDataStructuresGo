package main

import "fmt"

func main() {
	as()
	ad()
}

func as() {

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
