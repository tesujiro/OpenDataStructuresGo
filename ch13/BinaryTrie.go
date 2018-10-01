package ch13

import (
	"fmt"

	"github.com/tesujiro/OpenDataStructuresGo/ch1"
)

type Node struct {
	x      ch1.Comparable
	child  [2]*Node
	jump   *Node
	parent *Node
}

func NewNode() *Node {
	return &Node{child: [2]*Node{}}
}

func (n *Node) left() *Node {
	if n == nil {
		return nil
	}
	return n.child[0]
}

func (n *Node) right() *Node {
	if n == nil {
		return nil
	}
	return n.child[1]
}

type BinaryTrie struct {
	r     *Node
	dummy *Node
	w     uint // bit length
	n     int
}

func NewBinaryTrie() *BinaryTrie {
	dummy := NewNode()
	dummy.child[0] = dummy // dummy.prev = dummy
	dummy.child[1] = dummy // dummy.next = dummy

	return &BinaryTrie{
		r:     NewNode(),
		dummy: dummy,
		w:     64,
	}
}

func (bt *BinaryTrie) Size() int {
	return bt.n
}

func (bt *BinaryTrie) GetAll() []interface{} {
	s := []interface{}{}
	w := bt.dummy.child[1]
	// get x till end
	for w != nil && w != bt.dummy {
		//s = append(s, w.x.HashCode())
		s = append(s, w.x)
		w = w.right()
	}
	return s
}

func (bt *BinaryTrie) Print() {
	fmt.Printf("BinaryTrie(n:%v)=%v\n", bt.Size(), bt.GetAll())
}

func (bt *BinaryTrie) Find(x ch1.Comparable) ch1.Comparable {
	ix := x.HashCode()
	u := bt.r
	var c, i uint
	for i = 0; i < bt.w; i++ {
		c = ix >> (bt.w - i - 1) & 1
		if u.child[c] == nil {
			break
		}
		u = u.child[c]
	}
	// found x
	if i == bt.w {
		return u.x
	}
	// search for next value
	if c == 0 {
		u = u.jump
	} else {
		u = u.jump.right()
	}
	if u == bt.dummy {
		return nil
	} else {
		return u.x
	}
}

func (bt *BinaryTrie) Add(x ch1.Comparable) bool {
	ix := x.HashCode()
	u := bt.r
	var c, i uint
	// 1 - search for ix until falling out of the trie
	for i = 0; i < bt.w; i++ {
		c = ix >> (bt.w - i - 1) & 1
		if u == nil || u.child[c] == nil {
			break
		}
		u = u.child[c]
	}
	if i == bt.w {
		//fmt.Printf("Add failed x=%v\n", x)
		return false
	}
	var pred *Node
	if c == 1 { //right
		pred = u.jump
	} else { //left
		pred = u.jump.left()
	}
	if pred == nil {
		pred = bt.dummy
	}
	u.jump = nil
	// 2 - add path to ix
	for ; i < bt.w; i++ {
		c = ix >> (bt.w - i - 1) & 1
		u.child[c] = NewNode()
		u.child[c].parent = u
		u = u.child[c]
	}
	u.x = x
	// 3 - add u to linked list
	u.child[0] = pred         //set pred
	u.child[1] = pred.right() //set next
	u.child[0].child[1] = u   //u.pred.next=u
	u.child[1].child[0] = u   //u.next.pred=u
	// 4 - walk back up, updating jump pointers
	v := u.parent
	for v != nil {
		if (v.left() == nil && (v.jump == nil || v.jump.x.HashCode() > ix)) ||
			(v.right() == nil && (v.jump == nil || v.jump.x.HashCode() < ix)) {
			v.jump = u
		}
		v = v.parent
	}
	bt.n++
	return true
}

func (bt *BinaryTrie) Remove(x ch1.Comparable) bool {
	ix := x.HashCode()
	u := bt.r
	var c, i uint
	// 1 - find leaf, u, containing x
	for i = 0; i < bt.w; i++ {
		c = ix >> (bt.w - i - 1) & 1
		if u.child[c] == nil {
			return false
		}
		u = u.child[c]
	}
	// 2 - remove u from linked list
	u.child[0].child[1] = u.child[1] //u.prev.next=u.next
	u.child[1].child[0] = u.child[0] //u.next.pred=u.prev
	v := u
	// 3 - delete nodes on path to u
	for i = bt.w - 1; int(i) >= 0; i-- {
		c = ix >> (bt.w - i - 1) & 1
		v = v.parent
		//fmt.Println("Remove i=", i, " c=", c)
		v.child[c] = nil
		if v.child[1-c] != nil {
			break
		}
	}
	// 4 - update jump pointers
	c = ix >> (bt.w - i - 1) & 1
	v.jump = u.child[1-c]
	v = v.parent
	i--
	for ; int(i) >= 0; i-- {
		c = ix >> (bt.w - i - 1) & 1
		if v.jump == u {
			v.jump = u.child[1-c]
		}
		v = v.parent
	}
	bt.n--
	return true
}
