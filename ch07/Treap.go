package ch07

import (
	"fmt"
	"math/rand"

	"github.com/tesujiro/OpenDataStructuresGo/ch01"
)

type TreapNode struct {
	x      ch01.Comparable
	left   *TreapNode
	right  *TreapNode
	parent *TreapNode
	p      int
}

func NewTreapNode() *TreapNode {
	return &TreapNode{}
}

type Treap struct {
	r *TreapNode
	n int
}

func NewTreap() *Treap {
	//return &Treap{r: NewNode()}
	return &Treap{r: nil}
}

func (bt *Treap) Size() int {
	return bt.n
}

func (bt *Treap) getAll(u *TreapNode) []interface{} {
	if u == nil {
		return []interface{}{}
	}
	s := []interface{}{u.x}
	s = append(s, bt.getAll(u.left)...)
	s = append(s, bt.getAll(u.right)...)
	return s
}

func (bt *Treap) GetAll() []interface{} {
	return bt.getAll(bt.r)
}

func (bt *Treap) Print() {
	fmt.Printf("Treap(n:%v)=%v\n", bt.Size(), bt.GetAll())
}

func (bt *Treap) Find(x ch01.Comparable) ch01.Comparable {
	w := bt.r
	var z *TreapNode
	for w != nil {
		comp := x.Compare(w.x)
		switch {
		case comp < 0:
			z = w
			w = w.left
		case comp > 0:
			w = w.right
		default:
			return w.x
		}
	}
	if z == nil {
		return nil
	} else {
		return z.x
	}
}

func (bt *Treap) rotateLeft(u *TreapNode) {
	w := u.right
	w.parent = u.parent
	if w.parent != nil {
		if w.parent.left == u {
			w.parent.left = w
		} else {
			w.parent.right = w
		}
	}
	u.right = w.left
	if u.right != nil {
		u.right.parent = u
	}
	u.parent = w
	w.left = u
	if u == bt.r {
		bt.r = w
		bt.r.parent = nil
	}
}

func (bt *Treap) rotateRight(u *TreapNode) {
	w := u.left
	w.parent = u.parent
	if w.parent != nil {
		if w.parent.left == u {
			w.parent.left = w
		} else {
			w.parent.right = w
		}
	}
	u.left = w.right
	if u.left != nil {
		u.left.parent = u
	}
	u.parent = w
	w.right = u
	if u == bt.r {
		bt.r = w
		bt.r.parent = nil
	}
}

func (bt *Treap) Add(x ch01.Comparable) bool {
	u := NewTreapNode()
	u.x = x
	u.p = rand.Int()
	if bt.AddNode(u) {
		bt.bubbleUp(u)
		return true
	}
	return false
}

func (bt *Treap) bubbleUp(u *TreapNode) {
	for u.parent != nil && u.parent.p > u.p {
		if u.parent.right == u {
			bt.rotateLeft(u.parent)
		} else {
			bt.rotateRight(u.parent)
		}
	}
	if u.parent == nil {
		bt.r = u
	}
}

func (bt *Treap) AddNode(u *TreapNode) bool {
	p := bt.findLast(u.x)
	return bt.addChild(p, u)
}

func (bt *Treap) findLast(x ch01.Comparable) *TreapNode {
	w := bt.r
	var prev *TreapNode
	for w != nil {
		prev = w
		comp := x.Compare(w.x)
		switch {
		case comp < 0:
			w = w.left
		case comp > 0:
			w = w.right
		default:
			return w
		}
	}
	return prev
}

func (bt *Treap) addChild(p, u *TreapNode) bool {
	if p == nil {
		bt.r = u
	} else {
		comp := u.x.Compare(p.x)
		switch {
		case comp < 0:
			p.left = u
		case comp > 0:
			p.right = u
		default:
			return false
		}
		u.parent = p
		//fmt.Printf("addChild p.x=%v u.x=%v u.parent.x=%v\n", p.x, u.x, u.parent.x)
	}
	bt.n++
	return true
}

func (bt *Treap) Remove(x ch01.Comparable) bool {
	u := bt.findLast(x)
	if u != nil && u.x.Compare(x) == 0 {
		bt.trickleDown(u)
		bt.splice(u)
		return true
	}
	return false
}

func (bt *Treap) trickleDown(u *TreapNode) {
	for u.left != nil && u.right != nil {
		switch {
		case u.left == nil:
			bt.rotateLeft(u)
		case u.right == nil:
			bt.rotateRight(u)
		case u.left.p < u.right.p:
			bt.rotateRight(u)
		default:
			bt.rotateLeft(u)
		}
		if bt.r == u {
			bt.r = u.parent
		}
	}
}

func (bt *Treap) splice(u *TreapNode) {
	var s, p *TreapNode
	if u.left != nil {
		s = u.left
	} else {
		s = u.right
	}
	if u == bt.r {
		bt.r = s
		p = nil
	} else {
		p = u.parent
		if p.left == u {
			p.left = s
		} else {
			p.right = s
		}
	}
	if s != nil {
		s.parent = p
	}
	bt.n--
}
