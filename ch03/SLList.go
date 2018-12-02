package ch03

import "fmt"

type SLNode struct {
	next *SLNode
	x    interface{}
}

type SLList struct {
	head *SLNode
	tail *SLNode
	n    int
}

func NewSLList() *SLList {
	return &SLList{
		head: nil,
		tail: nil,
		n:    0,
	}
}

func (l *SLList) GetAll() []interface{} {
	s := []interface{}{}
	for n := l.head; n != nil; n = n.next {
		s = append(s, n.x)
	}
	return s
}

func (l *SLList) Print() {
	fmt.Printf("list=%v\n", l.GetAll())
}

func (l *SLList) Push(x interface{}) {
	u := SLNode{
		x:    x,
		next: l.head,
	}
	l.head = &u
	if l.n == 0 {
		l.tail = &u
	}
	l.n++
	//fmt.Printf("l.head.x=%v\n", l.head.x)
}

func (l *SLList) Pop() interface{} {
	if l.n == 0 {
		return nil
	}
	x := l.head.x
	l.head = l.head.next
	l.n--
	if l.n == 0 {
		l.tail = nil
	}
	return x
}

func (l *SLList) Remove() interface{} {
	return l.Pop()
}

func (l *SLList) Add(x interface{}) bool {
	u := SLNode{
		x: x,
	}
	if l.n == 0 {
		l.head = &u
	} else {
		l.tail.next = &u
	}
	l.tail = &u
	l.n++
	return true
}
