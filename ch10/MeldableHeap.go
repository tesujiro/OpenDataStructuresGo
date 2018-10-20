package ch10

import (
	"fmt"
	"math/rand"

	"github.com/tesujiro/OpenDataStructuresGo/ch1"
)

type Node struct {
	x                   ch1.Comparable
	right, left, parent *Node
}

func NewNode() *Node {
	return &Node{}
}

type MeldableHeap struct {
	r *Node
	n int
}

func NewMeldableHeap() *MeldableHeap {
	return &MeldableHeap{
		r: nil,
		n: 0,
	}
}

func (h *MeldableHeap) Size() int {
	return h.n
}

func (h *MeldableHeap) getAll(u *Node) []interface{} {
	if h == nil || u == nil {
		return []interface{}{}
	}
	s := []interface{}{u.x}
	s = append(s, h.getAll(u.left)...)
	s = append(s, h.getAll(u.right)...)
	return s
}

func (h *MeldableHeap) GetAll() []interface{} {
	return h.getAll(h.r)
}

func (h *MeldableHeap) Print() {
	fmt.Printf("MeldableHeap(n:%v)=%v\n", h.Size(), h.GetAll())
}

func merge(h1, h2 *Node) *Node {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}
	if h1.x.Compare(h2.x) > 0 {
		return merge(h2, h1)
	}
	if rand.Intn(2)%2 == 0 {
		h1.left = merge(h1.left, h2)
		if h1.left != nil {
			h1.left.parent = h1
		}
	} else {
		h1.right = merge(h1.right, h2)
		if h1.right != nil {
			h1.right.parent = h1
		}
	}
	return h1
}

func (h *MeldableHeap) Add(x ch1.Comparable) bool {
	u := NewNode()
	u.x = x
	h.r = merge(u, h.r)
	h.r.parent = nil
	h.n++
	return true
}

func (h *MeldableHeap) Remove() ch1.Comparable {
	x := h.r.x
	//tmp := h.r
	h.r = merge(h.r.left, h.r.right)
	if h.r != nil {
		h.r.parent = nil
	}
	h.n--
	return x
}
