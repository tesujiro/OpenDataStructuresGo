package ch3

import "fmt"

type DLNode struct {
	prev, next *DLNode
	x          interface{}
}

type DLList struct {
	dummy *DLNode
	len   int
}

func NewDLList() *DLList {
	dummy := &DLNode{}
	dummy.prev = dummy
	dummy.next = dummy
	return &DLList{
		dummy: dummy,
		len:   0,
	}
}

func (l *DLList) Len() int {
	return l.len
}

func (l *DLList) GetAll() []interface{} {
	slice := []interface{}{}
	for i := 0; i < l.len; i++ {
		slice = append(slice, l.Get(i))
	}
	return slice
}

func (l *DLList) Print() {
	fmt.Printf("list=%v\n", l.GetAll())
}

func (l *DLList) getNode(i int) *DLNode {
	var p *DLNode
	if i < l.len/2 {
		p = l.dummy.next
		for j := 0; j < i; j++ {
			p = p.next
		}
	} else {
		p = l.dummy
		for j := l.len; j > i; j-- {
			p = p.prev
		}
	}
	return p
}

func (l *DLList) Get(i int) interface{} {
	return l.getNode(i).x
}

func (l *DLList) Set(i int, v interface{}) interface{} {
	u := l.getNode(i)
	y := u.x
	u.x = v
	return y
}

func (l *DLList) addBefore(w *DLNode, x interface{}) *DLNode {
	u := &DLNode{prev: w.prev, next: w, x: x}
	u.next.prev = u
	u.prev.next = u
	l.len += 1
	return u
}

func (l *DLList) Add(i int, x interface{}) {
	l.addBefore(l.getNode(i), x)
}

func (l *DLList) remove(w *DLNode) {
	w.prev.next = w.next
	w.next.prev = w.prev
	l.len -= 1
}

func (l *DLList) Remove(i int) interface{} {
	w := l.getNode(i)
	x := w.x
	l.remove(w)
	return x
}
