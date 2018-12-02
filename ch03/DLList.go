package ch03

import "fmt"

type DLNode struct {
	prev, next *DLNode
	x          interface{}
}

type DLList struct {
	dummy *DLNode
	n     int
}

func NewDLList() *DLList {
	dummy := &DLNode{}
	dummy.prev = dummy
	dummy.next = dummy
	return &DLList{
		dummy: dummy,
		n:     0,
	}
}

func (l *DLList) Size() int {
	return l.n
}

func (l *DLList) GetAll() []interface{} {
	slice := []interface{}{}
	for i := 0; i < l.n; i++ {
		slice = append(slice, l.Get(i))
	}
	return slice
}

func (l *DLList) Print() {
	fmt.Printf("list=%v\n", l.GetAll())
}

func (l *DLList) getNode(i int) *DLNode {
	var p *DLNode
	if i < l.n/2 {
		p = l.dummy.next
		for j := 0; j < i; j++ {
			p = p.next
		}
	} else {
		p = l.dummy
		for j := l.n; j > i; j-- {
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
	l.n++
	return u
}

func (l *DLList) Add(i int, x interface{}) {
	l.addBefore(l.getNode(i), x)
}

func (l *DLList) remove(w *DLNode) {
	w.prev.next = w.next
	w.next.prev = w.prev
	l.n--
}

func (l *DLList) Remove(i int) interface{} {
	w := l.getNode(i)
	x := w.x
	l.remove(w)
	return x
}
