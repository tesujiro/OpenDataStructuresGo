package ch06

import (
	"fmt"

	"github.com/tesujiro/OpenDataStructuresGo/ch01"
)

type Node struct {
	x      ch01.Comparable
	left   *Node
	right  *Node
	parent *Node
}

func NewNode() *Node {
	return &Node{}
}

type BinaryTree struct {
	r *Node
	n int
}

func NewBinaryTree() *BinaryTree {
	//return &BinaryTree{r: NewNode()}
	return &BinaryTree{r: nil}
}

func (bt *BinaryTree) depth(u *Node) int {
	d := 0
	for u != bt.r {
		u = u.parent
		d++
	}
	return d
}

func (bt *BinaryTree) Size() int {
	return bt.n
}

func (bt *BinaryTree) getAll(u *Node) []interface{} {
	if u == nil {
		return []interface{}{}
	}
	s := []interface{}{u.x}
	s = append(s, bt.getAll(u.left)...)
	s = append(s, bt.getAll(u.right)...)
	return s
}

func (bt *BinaryTree) GetAll() []interface{} {
	return bt.getAll(bt.r)
}

func (bt *BinaryTree) Print() {
	fmt.Printf("BinaryTree(n:%v)=%v\n", bt.Size(), bt.GetAll())
}

func (bt *BinaryTree) findNode(x ch01.Comparable) *Node {
	w := bt.r
	for w != nil {
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
	return nil
}

func (bt *BinaryTree) Find(x ch01.Comparable) ch01.Comparable {
	w := bt.r
	var z *Node
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

func (bt *BinaryTree) Add(x ch01.Comparable) bool {
	p := bt.findLast(x)
	u := NewNode()
	u.x = x
	return bt.addChild(p, u)
}

func (bt *BinaryTree) findLast(x ch01.Comparable) *Node {
	w := bt.r
	var prev *Node
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

func (bt *BinaryTree) addChild(p, u *Node) bool {
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
	}
	bt.n++
	return true
}

func (bt *BinaryTree) splice(u *Node) {
	var s, p *Node
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

func (bt *BinaryTree) removeNode(u *Node) {
	if u.left == nil || u.right == nil {
		//fmt.Println("u.x=", u.x)
		bt.splice(u)
	} else {
		w := u.right
		for w.left != nil {
			w = w.left
		}
		u.x = w.x
		bt.splice(w)
	}
}

func (bt *BinaryTree) Remove(x ch01.Comparable) bool {
	u := bt.findNode(x)
	if u == nil {
		return false
	} else {
		bt.removeNode(u)
		return true
	}
}
