package ch03

// A Space-Efficient Linked List

import (
	"fmt"

	"github.com/tesujiro/OpenDataStructuresGo/ch02"
)

const b = 8 // Block size

type BDeque struct {
	ch02.ArrayDeque
}

func NewBDeque(b int) *BDeque {
	a := make([]interface{}, b+1)
	d := &BDeque{*ch02.NewArrayDeque()}
	d.A = a
	d.Override = d
	return d
}

func (bd *BDeque) AddLast(x interface{}) bool {
	bd.Add(bd.Size(), x)
	return true
}

func (bd *BDeque) Resize() {
	// Do nothing
}

type SENode struct {
	d          *BDeque
	prev, next *SENode
}

func newSENode(b int) *SENode {
	return &SENode{
		//d: newBDeque(b),
		d: NewBDeque(b),
	}
}

type SEList struct {
	dummy *SENode
	n     int
}

func NewSEList() *SEList {
	dummy := newSENode(b)
	dummy.prev = dummy
	dummy.next = dummy
	return &SEList{
		dummy: dummy,
		n:     0,
	}
}

func (l *SEList) Size() int {
	return l.n
}

func (l *SEList) GetAll() []interface{} {
	slice := []interface{}{}
	for i := 0; i < l.n; i++ {
		slice = append(slice, l.Get(i))
	}
	return slice
}

func (l *SEList) Print() {
	fmt.Printf("list=%v\n", l.GetAll())
}

type SELocation struct {
	u *SENode
	j int
}

//func NewSELocation() *SELocation {
//return &SELocation{}
//}

func (l *SEList) getLocation(i int) *SELocation {
	if i < l.n/2 {
		u := l.dummy.next
		for i >= u.d.Size() {
			i -= u.d.Size()
			u = u.next
		}
		return &SELocation{u: u, j: i}
	} else {
		u := l.dummy
		idx := l.n
		for i < idx {
			u = u.prev
			idx -= u.d.Size()
		}
		return &SELocation{u: u, j: i - idx}
	}
}

func (l *SEList) Get(i int) interface{} {
	loc := l.getLocation(i)
	return loc.u.d.Get(loc.j)
}

func (l *SEList) Set(i int, x interface{}) interface{} {
	loc := l.getLocation(i)
	y := loc.u.d.Get(loc.j)
	loc.u.d.Set(loc.j, x)
	return y
}

func (l *SEList) addBefore(w *SENode) *SENode {
	//u := &SENode{prev: w.prev, next: w, d: NewBDeque(b)}
	u := newSENode(b)
	u.prev = w.prev
	u.next = w
	u.prev.next = u
	u.next.prev = u
	return u
}

func (l *SEList) AddLast(x interface{}) {
	last := l.dummy.prev
	if last == l.dummy || last.d.Size() == b+1 {
		//fmt.Printf("addBefore\n")
		last = l.addBefore(l.dummy)
	}
	//fmt.Printf("*last.d=%#v\n", *last.d)
	last.d.AddLast(x)
	l.n++
}

func (l *SEList) spread(u *SENode) {
	w := u
	for j := 0; j < b-1; j++ {
		w = w.next
	}
	w = l.addBefore(w)
	for w != u {
		for w.d.Size() < b {
			w.d.Add(0, w.prev.d.Size()-1)
		}
		w = w.prev
	}
	l.removeNode(w)
}

func (l *SEList) Add(i int, x interface{}) {
	if i == l.n {
		l.AddLast(x)
		return
	}
	loc := l.getLocation(i)
	u := loc.u
	r := 0
	for r < b && u != l.dummy && u.d.Size() == b+1 {
		u = u.next
		r++
	}
	if r == b {
		l.spread(loc.u)
		u = loc.u
	}
	if u == l.dummy {
		u = l.addBefore(u)
	}
	for u != loc.u {
		u.d.Add(0, u.prev.d.Remove(u.prev.d.Size()-1))
		u = u.prev
	}
	u.d.Add(loc.j, x)
	l.n++
}

func (l *SEList) removeNode(w *SENode) {
	w.prev.next = w.next
	w.next.prev = w.prev
}

func (l *SEList) gather(u *SENode) {
	w := u
	for j := 0; j < b-1; j++ {
		for w.d.Size() < b {
			w.d.AddLast(w.next.d.Remove(0))
		}
		w = w.next
	}
	l.removeNode(w)
}

func (l *SEList) Remove(i int) interface{} {
	loc := l.getLocation(i)
	y := loc.u.d.Get(loc.j)
	u := loc.u
	r := 0
	for r < b && u != l.dummy && u.d.Size() == b-1 {
		u = u.next
		r++
	}
	if r == b {
		l.gather(loc.u)
	}
	u = loc.u
	u.d.Remove(loc.j)
	for u.d.Size() < b-1 && u.next != l.dummy {
		u.d.AddLast(u.next.d.Remove(0))
		u = u.next
	}
	if u.d.Size() == 0 {
		l.removeNode(u)
	}
	l.n--
	return y
}
