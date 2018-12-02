package ch4

import (
	"fmt"
)

type SLNode struct {
	x      interface{}
	height int
	length []int     //  *int
	next   []*SLNode // **SLNode
}

func NewSLNode(x interface{}, h int) *SLNode {
	return &SLNode{
		x:      x,
		height: h,
		length: make([]int, h+1),
		next:   make([]*SLNode, h+1),
	}
}

type SkiplistList struct {
	sentinel *SLNode
	n        int
	h        int
}

func NewSkiplistList() *SkiplistList {
	return &SkiplistList{sentinel: NewSLNode(nil, 0), h: 0}
}

func (l *SkiplistList) Size() int {
	return l.n
}

func (l *SkiplistList) GetAll() []interface{} {
	s := []interface{}{}
	u := l.sentinel
	//fmt.Printf("l=%#v\n", l)
	//fmt.Printf("len(u.next)=%v\n", len(u.next))
	for u.next[0] != nil {
		s = append(s, u.next[0].x)
		u = u.next[0]
		//fmt.Printf("s=%#v", s)
	}
	return s
}

func (l *SkiplistList) Print() {
	fmt.Printf("SkiplistList(n:%v h:%v)=%v\n", l.Size(), l.h, l.GetAll())
}

func (l *SkiplistList) findPred(i int) *SLNode {
	u := l.sentinel
	r := l.h
	j := -1
	for r >= 0 {
		for u.next[r] != nil && j+u.length[r] < i {
			j += u.length[r]
			u = u.next[r]
		}
		r--
	}
	return u
}
func (l *SkiplistList) Get(i int) interface{} {
	return l.findPred(i).next[0].x
}

func (l *SkiplistList) Set(i int, x interface{}) interface{} {
	u := l.findPred(i).next[0]
	y := u.x
	u.x = x
	return y
}

func (l *SkiplistList) AddNode(i int, w *SLNode) *SLNode {
	u := l.sentinel
	k := w.height
	r := l.h
	j := -1
	for r >= 0 {
		//fmt.Printf("r=%v len(u.next)=%v\n", r, len(u.next))
		//fmt.Printf("u.next[r]=%v\n", u.next[r])
		for u.next[r] != nil && j+u.length[r] < i {
			j += u.length[r]
			u = u.next[r]
		}
		u.length[r]++
		if r <= k {
			w.next[r] = u.next[r]
			u.next[r] = w
			w.length[r] = u.length[r] - (i - j)
			u.length[r] = i - j
		}
		r--
	}
	l.n++
	return u
}

func (l *SkiplistList) Add(i int, x interface{}) {
	w := NewSLNode(x, pickHeight())
	/*
		if w.height > l.h {
			l.h = w.height
		}
	*/
	for ; l.h < w.height; l.h++ {
		l.sentinel.next = append(l.sentinel.next, nil)
		l.sentinel.length = append(l.sentinel.length, 0)
		l.sentinel.height++ // ??
	}
	l.AddNode(i, w)
}

func (l *SkiplistList) Remove(i int) interface{} {
	var x interface{}
	u := l.sentinel
	r := l.h
	j := -1
	for r >= 0 {
		for u.next[r] != nil && j+u.length[r] < i {
			j += u.length[r]
			u = u.next[r]
		}
		u.length[r]--
		if j+u.length[r]+1 == i && u.next[r] != nil {
			x = u.next[r].x
			u.length[r] += u.next[r].length[r]
			u.next[r] = u.next[r].next[r]
			if u == l.sentinel && u.next[r] == nil {
				l.h--
			}
		}
		r--
	}
	l.n--
	return x
}
