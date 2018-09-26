package ch8

import (
	"fmt"
	"math"

	"github.com/tesujiro/OpenDataStructuresGo/ch1"
)

type Node struct {
	x      ch1.Comparable
	left   *Node
	right  *Node
	parent *Node
}

func NewNode() *Node {
	return &Node{}
}

type ScapegoatTree struct {
	r *Node
	n int
	q int
}

func NewScapegoatTree() *ScapegoatTree {
	return &ScapegoatTree{}
}

func (bt *ScapegoatTree) Size() int {
	return bt.n
}

func (bt *ScapegoatTree) SizeNode(u *Node) int {
	if u == nil {
		return 0
	}
	return 1 + bt.SizeNode(u.left) + bt.SizeNode(u.right)
}

func (bt *ScapegoatTree) getAll(u *Node) []interface{} {
	if u == nil {
		return []interface{}{}
	}
	s := []interface{}{u.x}
	s = append(s, bt.getAll(u.left)...)
	s = append(s, bt.getAll(u.right)...)
	return s
}

func (bt *ScapegoatTree) GetAll() []interface{} {
	return bt.getAll(bt.r)
}

func (bt *ScapegoatTree) Print() {
	fmt.Printf("ScapegoatTree(n:%v)=%v\n", bt.Size(), bt.GetAll())
}

func (bt *ScapegoatTree) findNode(x ch1.Comparable) *Node {
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

func (bt *ScapegoatTree) Find(x ch1.Comparable) ch1.Comparable {
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

func (bt *ScapegoatTree) rebuild(u *Node) {
	ns := bt.SizeNode(u)
	p := u.parent
	a := make([]*Node, ns)
	packIntoArray(u, a, 0)
	if p == nil {
		bt.r = buildBalanced(a, 0, ns)
		bt.r.parent = nil
	} else if p.right == u {
		p.right.parent = p
	} else {
		p.left = buildBalanced(a, 0, ns)
		p.left.parent = p
	}
}

func buildBalanced(a []*Node, i, ns int) *Node {
	if ns == 0 {
		return nil
	}
	m := ns / 2
	a[i+m] = buildBalanced(a, i, m)
	if a[i+m].left != nil {
		a[i+m].left.parent = a[i+m]
	}
	a[i+m].right = buildBalanced(a, i+m+1, ns-m-1)
	if a[i+m].right != nil {
		a[i+m].right.parent = a[i+m]
	}
	//a[i+m].q = int(math.Ceil(math.Log(len(a[i+m]))))
	return a[i+m]
}

func (bt *ScapegoatTree) addWithDepth(u *Node) int {
	p, d := bt.findLast(u.x)
	bt.addChild(p, u)
	return d
}

func (bt *ScapegoatTree) findLast(x ch1.Comparable) (*Node, int) {
	w := bt.r
	d := 0
	q := 1
	var prev *Node
	//TODO:??
	if w == nil {
		d = 1
	}
	for w != nil {
		prev = w
		comp := x.Compare(w.x)
		switch {
		case comp < 0:
			w = w.left
		case comp > 0:
			w = w.right
		default:
			return w, d
		}
		d++
		//bt.q++
		q += q * 2
		//fmt.Println("bt.q add! :", bt.q)
	}
	if q > bt.q {
		bt.q = q
	}
	return prev, d
}

func (bt *ScapegoatTree) addChild(p, u *Node) bool {
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

func packIntoArray(u *Node, a []*Node, i int) int {
	if u == nil {
		return i
	}
	i = packIntoArray(u.left, a, i)
	a[i] = u
	i++
	return packIntoArray(u.right, a, i)
}

func (bt *ScapegoatTree) Add(x ch1.Comparable) bool {
	u := NewNode()
	u.x = x
	/*
		u.parent = nil
		u.left = nil
		u.right = nil
	*/
	//u.parent = u
	d := bt.addWithDepth(u)
	//fmt.Println("d=", d)
	//fmt.Println("bt.q=", bt.q)
	//fmt.Println("bt.r.x=", bt.r.x)
	//if float64(d) > math.Log(float64(bt.q)) {
	if u.parent != nil && float64(d) > math.Log(float64(bt.q)) {
		fmt.Println("d > math.Log(bt.q)=", math.Log(float64(bt.q)))
		w := u.parent
		a := bt.SizeNode(w)
		b := bt.SizeNode(w.parent)
		//fmt.Println("a=", a, "  b=", b)
		//fmt.Println("w.x=", w.x)
		//for w != bt.r && 3*a <= 2*b {
		for 3*a <= 2*b {
			w = w.parent
			a = bt.SizeNode(w)
			b = bt.SizeNode(w.parent)
			fmt.Println("a=", a, "  b=", b)
		}
		bt.rebuild(w.parent)
	} else if d < 0 {
		return false
	}
	return true
}

func (bt *ScapegoatTree) splice(u *Node) {
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

func (bt *ScapegoatTree) removeNode(u *Node) {
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

func (bt *ScapegoatTree) Remove(x ch1.Comparable) bool {
	u := bt.findNode(x)
	if u == nil {
		return false
	} else {
		bt.removeNode(u)
		return true
	}
}
