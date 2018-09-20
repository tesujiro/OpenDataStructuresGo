package ch4

import (
	"fmt"
	"math/rand"
)

func compare(x1, x2 interface{}) int {
	return x1.(int) - x2.(int)
}

type SSSNode struct {
	x      interface{}
	height int //height
	next   []*SSSNode
}

func NewSSSNode(h int, x interface{}) *SSSNode {
	return &SSSNode{
		x:      x,
		height: h,
		next:   make([]*SSSNode, 0), //TODO
	}
}

type SkiplistSSet struct {
	sentinel *SSSNode
	n        int
	h        int
}

func NewSkiplistSSet() *SkiplistSSet {
	return &SkiplistSSet{sentinel: NewSSSNode(0, nil)}
}

func (l *SkiplistSSet) Size() int {
	return l.n
}

func (l *SkiplistSSet) GetAll() []interface{} {
	return []interface{}{}
}

func (l *SkiplistSSet) Print() {
	fmt.Printf("SkiplistSSet(n:%v)=%v\n", l.Size(), l.GetAll())
}

func (l *SkiplistSSet) findPredNode(x interface{}) *SSSNode {
	u := l.sentinel
	r := l.h
	for r >= 0 {
		for u.next[r] != nil && compare(u.next[r].x, x) < 0 {
			u = u.next[r]
		}
		r--
	}
	return u
}

func (l *SkiplistSSet) Find(x interface{}) interface{} {
	u := l.findPredNode(x)
	if u.next[0] == nil {
		return nil
	} else {
		return u.next[0].x
	}
}

func pickHeight() int {
	z := rand.Int()
	k := 0
	m := 1
	for (z & m) != 0 {
		k++
		m <<= 1
	}
	return k
}

func (l *SkiplistSSet) Add(x interface{}) bool {
	u := l.sentinel
	r := l.h
	stack := make([]*SSSNode, 0) //TODO
	//comp := 0
	for r >= 0 {
		for u.next[r] != nil && compare(u.next[r].x, x) < 0 {
			u = u.next[r]
		}
		if u.next[r] != nil && compare(u.next[r].x, x) == 0 {
			return false
		}
		stack[r] = u
		r--
	}
	w := NewSSSNode(pickHeight(), x)
	for l.h < w.height {
		l.h++
		stack[l.h] = l.sentinel
	}
	for i := 0; i <= w.height; i++ {
		w.next[i] = stack[i].next[i] //TODO: append??
		stack[i].next[i] = w         //TODO: append?
	}
	l.n++
	return true
}

func (l *SkiplistSSet) Remove(x interface{}) bool {
	removed := false
	u := l.sentinel
	r := l.h
	//var del *SSSNode
	for r >= 0 {
		for u.next[r] != nil && compare(u.next[r].x, x) < 0 {
			u = u.next[r]
		}
		if u.next[r] != nil && compare(u.next[r].x, x) == 0 {
			removed = true
			//del = u.next[r]
			u.next[r] = u.next[r].next[r]
			if u == l.sentinel && u.next[r] == nil {
				l.h--
			}
		}
		r--
	}
	if removed {
		//delete(del)
		//del = nil
		l.n--
	}
	return removed
}
