package ch13

import (
	"fmt"

	"github.com/tesujiro/OpenDataStructuresGo/ch1"
	"github.com/tesujiro/OpenDataStructuresGo/ch5"
)

type XFastTrie struct {
	BinaryTrie
	t []*ch5.LinearHashTable
}

func NewXFastTrie() *XFastTrie {
	ft := &XFastTrie{*NewBinaryTrie(), nil}
	ft.t = make([]*ch5.LinearHashTable, ft.w+1)
	for i := 0; i < len(ft.t); i++ {
		ft.t[i] = ch5.NewLinearHashTable()
	}
	return ft
}

/*
//import from BinaryTrie
func (ft *XFastTrie) Size() int {
	return ft.n
}

//import from BinaryTrie
func (ft *XFastTrie) GetAll() []interface{} {
	s := []interface{}{}
	w := ft.dummy.child[1]
	// get x till end
	for w != nil && w != ft.dummy {
		s = append(s, w.x)
		w = w.right()
	}
	return s
}
*/

func (ft *XFastTrie) Print() {
	fmt.Printf("XFastTrie(n:%v)=%v\n", ft.Size(), ft.GetAll())
}

type XPair struct {
	x uint
	u *Node
}

func NewXPair(x uint) *XPair {
	return &XPair{x: x}
}

func (xp1 XPair) Compare(xp2 ch1.Comparable) int {
	return int(xp1.x - xp2.(XPair).x)
}

func (xp XPair) HashCode() uint {
	return xp.x
}

func (ft *XFastTrie) Find(x ch1.Comparable) ch1.Comparable {
	l := uint(0)
	h := ft.w + 1
	ix := x.HashCode()
	u := ft.r
	for h-l > 1 {
		i := (l + h) / 2
		p := NewXPair(ix >> (ft.w - i))
		//fmt.Printf("Find XPair ix=%v, i=%v, p=%v, len(ft.t[i])=%v,ft.t[i]=%v\n", ix, i, p.x, ft.t[i].Size(), ft.t[i].GetAll())
		xp := ft.t[i].Find(p)
		if xp == nil {
			h = i
		} else {
			//fmt.Println(" ==> FOUND!")
			//u = xp.(XPair).u
			u = xp.(*XPair).u
			l = i
		}
	}
	// found x
	if l == ft.w {
		return u.x
	}
	// search for next value

	c := ix >> (ft.w - l - 1) & 1
	if c == 0 {
		u = u.jump
	} else {
		u = u.jump.right()
	}
	var pred *Node
	//TODO???
	if u == nil {
		return nil
	}
	//if c == 1 {
	if c == 1 {
		pred = u.jump
	} else {
		//fmt.Printf("u=%#v\n", u)
		pred = u.jump.left()
		//pred = u.jump.child[0]
	}
	//if pred.right() == ft.dummy {
	if pred.right() == ft.dummy || pred.right() == nil {
		return nil
	} else {
		return pred.right().x
	}
}

func (ft *XFastTrie) Add(x ch1.Comparable) bool {
	ix := x.HashCode()
	u := ft.r
	var c, i uint
	// 1 - search for ix until falling out of the trie
	for i = 0; i < ft.w; i++ {
		c = ix >> (ft.w - i - 1) & 1
		//if u == nil || u.child[c] == nil {
		if u.child[c] == nil {
			break
		}
		u = u.child[c]
	}
	if i == ft.w {
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
		pred = ft.dummy
	}
	u.jump = nil
	// 2 - add path to ix
	for ; i < ft.w; i++ {
		c = ix >> (ft.w - i - 1) & 1
		u.child[c] = NewNode()
		u.child[c].parent = u

		p := NewXPair(ix >> (ft.w - i))
		//p.u = u.child[c]
		p.u = u
		//fmt.Printf("Add NewXPair ix=%v,i=%v,x=%v\n", ix, i, p.x)
		if !ft.t[i].Add(p) {
			//fmt.Printf("XFastTrie Add XPair failed. ix=%v,i=%v\n", ix, i)
		}
		//fmt.Printf(" XPair ix=%v, i=%v, p=%v, len(ft.t[i])=%v\n", ix, i, p.x, ft.t[i].Size())
		//ft.t[i].Print()
		u = u.child[c]
	}
	u.x = x
	// 3 - add u to linked list
	u.child[0] = pred         //set pred
	u.child[1] = pred.right() //set next
	u.child[0].child[1] = u   //u.pred.next=u
	u.child[1].child[0] = u   //u.next.pred=u
	p := NewXPair(ix >> (ft.w - i))
	p.u = u
	//fmt.Printf("Add NewXPair ix=%v,i=%v,x=%v\n", ix, i, p.x)
	if !ft.t[i].Add(p) {
		fmt.Printf("XFastTrie Add XPair failed. ix=%v,i=%v\n", ix, i)
	}
	// 4 - walk back up, updating jump pointers
	v := u.parent
	for v != nil {
		if (v.left() == nil && (v.jump == nil || v.jump.x.HashCode() > ix)) ||
			(v.right() == nil && (v.jump == nil || v.jump.x.HashCode() < ix)) {
			v.jump = u
		}
		v = v.parent
	}
	ft.n++
	return true
}

func (ft *XFastTrie) Remove(x ch1.Comparable) bool {
	ix := x.HashCode()
	u := ft.r
	var c, i uint
	// 1 - find leaf, u, containing x
	for i = 0; i < ft.w; i++ {
		c = ix >> (ft.w - i - 1) & 1
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
	for i = ft.w - 1; int(i) >= 0; i-- {
		c = ix >> (ft.w - i - 1) & 1
		v = v.parent
		//fmt.Println("Remove i=", i, " c=", c)
		v.child[c] = nil
		if v.child[1-c] != nil {
			break
		}
	}
	// 4 - update jump pointers
	c = ix >> (ft.w - i - 1) & 1
	v.jump = u.child[1-c]
	v = v.parent
	i--
	for ; int(i) >= 0; i-- {
		c = ix >> (ft.w - i - 1) & 1
		if v.jump == u {
			v.jump = u.child[1-c]
		}
		v = v.parent
	}
	ft.n--
	return true
}
