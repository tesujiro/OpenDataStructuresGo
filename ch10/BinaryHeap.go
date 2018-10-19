package ch10

import (
	"fmt"

	"github.com/tesujiro/OpenDataStructuresGo/ch1"
)

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func parent(i int) int {
	return (i - 1) / 2
}

type BinaryHeap struct {
	a []ch1.Comparable
	n int
}

const initSize = 1

func NewBinaryHeap() *BinaryHeap {
	return &BinaryHeap{
		a: make([]ch1.Comparable, initSize),
		n: 0,
	}
}

func (h *BinaryHeap) cap() int {
	return cap(h.a)
}

func (h *BinaryHeap) Size() int {
	return h.n
}

func (h *BinaryHeap) GetAll() []interface{} {
	is := make([]interface{}, h.n)
	for i := 0; i < h.n; i++ {
		is[i] = h.a[i]
	}
	return is
}

func (h *BinaryHeap) Print() {
	fmt.Printf("BinaryHeap(n:%v,cap:%v)=%v\n", h.Size(), h.cap(), h.GetAll())
}

func (h *BinaryHeap) Get(i int) ch1.Comparable {
	//i = i % h.n
	return h.a[i]
}

func (h *BinaryHeap) Set(i int, v ch1.Comparable) ch1.Comparable {
	//i = i % h.n
	y := h.a[i]
	h.a[i] = v
	return y
}

func (h *BinaryHeap) Add(x ch1.Comparable) bool {
	if h.n+1 > cap(h.a) {
		h.resize()
	}
	h.a[h.n] = x
	h.n++
	h.bubbleUp(h.n - 1)
	return true
}

func (h *BinaryHeap) bubbleUp(i int) {
	p := parent(i)
	for i > 0 && h.a[i].Compare(h.a[p]) < 0 {
		h.a[i], h.a[p] = h.a[p], h.a[i]
		i = p
		p = parent(i)
	}
}

func (h *BinaryHeap) Remove() ch1.Comparable {
	x := h.a[0]
	h.n--
	h.a[0] = h.a[h.n]
	h.trickleDown(0)
	if 3*h.n < h.cap() {
		h.resize()
	}
	return x
}

func (h *BinaryHeap) trickleDown(i int) {
	for i >= 0 {
		j := -1
		r := right(i)
		if r < h.n && h.a[r].Compare(h.a[i]) < 0 {
			l := left(i)
			if h.a[l].Compare(h.a[r]) < 0 {
				j = l
			} else {
				j = r
			}
		} else {
			l := left(i)
			if l < h.n && h.a[l].Compare(h.a[i]) < 0 {
				j = l
			}
		}
		if j >= 0 {
			h.a[i], h.a[j] = h.a[j], h.a[i]
		}
		i = j
	}
}

func (h *BinaryHeap) resize() {
	var new []ch1.Comparable
	new = make([]ch1.Comparable, h.n*2)
	for i := 0; i < h.n; i++ {
		new[i] = h.a[i]
	}
	h.a = new
}
