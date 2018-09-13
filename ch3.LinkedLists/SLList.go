package main

import "fmt"

type SLNode struct {
	next *SLNode
	x    interface{}
}

type SLList struct {
	head *SLNode
	tail *SLNode
	len  int
}

func newSLList() *SLList {
	return &SLList{
		head: nil,
		tail: nil,
		len:  0,
	}
}

func (l *SLList) List() []interface{} {
	var slice []interface{}
	for n := l.head; n != nil; n = n.next {
		slice = append(slice, n.x)
	}
	return slice
}

func (l *SLList) Print() {
	fmt.Printf("list=%v\n", l.List())
}

func (l *SLList) Push(x interface{}) {
	u := SLNode{
		x:    x,
		next: l.head,
	}
	l.head = &u
	if l.len == 0 {
		l.tail = &u
	}
	l.len += 1
	//fmt.Printf("l.head.x=%v\n", l.head.x)
}

func (l *SLList) Pop() interface{} {
	if l.len == 0 {
		return nil
	}
	x := l.head.x
	l.head = l.head.next
	l.len -= 1
	if l.len == 0 {
		l.tail = nil
	}
	return x
}

func main() {
	fmt.Println("SLList")
	sll := newSLList()
	sll.Print()
	sll.Push(1)
	sll.Push(2)
	sll.Push(3)
	sll.Print()
	fmt.Printf("Pop()=%v\n", sll.Pop())
	sll.Print()
}
