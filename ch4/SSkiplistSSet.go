package ch4

import (
	"fmt"
	"math/rand"

	"github.com/tesujiro/OpenDataStructuresGo/ch1"
)

type SSSNode struct {
	x      ch1.Comparable
	height int //height
	next   []*SSSNode
}

func NewSSSNode(h int, x ch1.Comparable) *SSSNode {
	return &SSSNode{
		x:      x,
		height: h,
		next:   make([]*SSSNode, h+1),
	}
}

type SkiplistSSet struct {
	sentinel *SSSNode
	n        int
	h        int
}

func NewSkiplistSSet() *SkiplistSSet {
	return &SkiplistSSet{sentinel: NewSSSNode(0, nil), h: 0}
}

func (l *SkiplistSSet) Size() int {
	return l.n
}

func (l *SkiplistSSet) GetAll() []interface{} {
	//return []interface{}{}
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

func (l *SkiplistSSet) Print() {
	fmt.Printf("SkiplistSSet(n:%v)=%v\n", l.Size(), l.GetAll())
}

func (l *SkiplistSSet) findPredNode(x ch1.Comparable) *SSSNode {
	u := l.sentinel
	r := l.h
	for r >= 0 {
		for u.next[r] != nil && u.next[r].x.Compare(x) < 0 {
			u = u.next[r]
		}
		r--
	}
	return u
}

func (l *SkiplistSSet) Find(x ch1.Comparable) ch1.Comparable {
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
	//fmt.Println("pickHeight()=", k)
	return k
}

func (l *SkiplistSSet) Add(x ch1.Comparable) bool {
	u := l.sentinel
	r := l.h
	stack := make([]*SSSNode, l.h+1) //TODO
	//comp := 0
	for r >= 0 {
		//fmt.Printf("r=%v len(u.next)=%v\n", r, len(u.next))
		//fmt.Printf("u.next[r]=%v\n", u.next[r])
		for u.next[r] != nil && u.next[r].x.Compare(x) < 0 {
			u = u.next[r]
		}
		if u.next[r] != nil && u.next[r].x.Compare(x) == 0 {
			return false
		}
		//fmt.Printf("r=%v len(stack)=%v len(u.next)=%v sentinel.x=%v u=%#v\n", r, len(stack), len(u.next), l.sentinel.x, u)
		stack[r] = u
		r--
	}
	w := NewSSSNode(pickHeight(), x)
	for l.h < w.height {
		l.h++
		stack = append(stack, l.sentinel)
		l.sentinel.next = append(l.sentinel.next, nil)
	}
	for i := 0; i <= w.height; i++ {
		//fmt.Println("stack[", i, "]=", stack[i])
		//fmt.Println("stack[", i, "].next[", i, "]=", stack[i].next[i])
		w.next[i] = stack[i].next[i]
		//w.next = append(w.next, stack[i].next[i])
		stack[i].next[i] = w
		//stack[i].next = append(stack[i].next, w)
	}
	l.n++
	return true
}

func (l *SkiplistSSet) Remove(x ch1.Comparable) bool {
	removed := false
	u := l.sentinel
	r := l.h
	//var del *SSSNode
	for r >= 0 {
		for u.next[r] != nil && u.next[r].x.Compare(x) < 0 {
			u = u.next[r]
		}
		if u.next[r] != nil && u.next[r].x.Compare(x) == 0 {
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
