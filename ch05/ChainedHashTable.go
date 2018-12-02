package ch05

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/tesujiro/OpenDataStructuresGo/ch01"
	"github.com/tesujiro/OpenDataStructuresGo/ch04"
)

type ChainedHashTable struct {
	t []ch01.List
	n int    // number of element
	d uint32 // bit size of Hashtable
	w uint32 // bit size of int
	z uint   // random number
}

const initWidth = 20

func newList() ch01.List {
	return ch04.NewSkiplistList()
}

func NewChainedHashTable() *ChainedHashTable {
	initSize := 2 << (initWidth - 1)
	t := make([]ch01.List, initSize)
	for i := 0; i < initSize; i++ {
		t[i] = newList()
	}
	return &ChainedHashTable{
		t: t,
		n: 0,
		d: initWidth,
		w: strconv.IntSize,
		z: uint(rand.Int()) * 98767,
	}
}

func (ht *ChainedHashTable) cap() int {
	return len(ht.t)
}

func (ht *ChainedHashTable) Size() int {
	return ht.n // number of element
}

func (ht *ChainedHashTable) GetAll() []interface{} {
	slice := []interface{}{}
	for i := 0; i < ht.cap(); i++ {
		for j := 0; j < ht.t[i].Size(); j++ {
			slice = append(slice, ht.t[i].Get(j))
		}
	}
	return slice
}

func (ht *ChainedHashTable) Print() {
	fmt.Printf("ChainedHashTable(n:%v,cap:%v)=%v\n", ht.Size(), ht.cap(), ht.GetAll())
}

func (ht *ChainedHashTable) resize() {
	var new []ch01.List
	if ht.n > 1 {
		new = make([]ch01.List, ht.n*2)
		ht.d++
	} else {
		new = make([]ch01.List, 1)
		ht.d = 1
	}
	for i := 0; i < len(new); i++ {
		new[i] = newList()
	}
	for i := 0; i < ht.n; i++ {
		for j := 0; j < ht.t[i].Size(); j++ {
			x := ht.t[i].Get(j).(ch01.Comparable)
			new[ht.hash(x)].Add(0, x)
		}
	}
	ht.t = new
}

func (ht *ChainedHashTable) Add(x ch01.Comparable) bool {
	if ht.Find(x) != nil {
		return false
	}
	if ht.n+1 > ht.cap() {
		ht.resize()
	}
	ht.t[ht.hash(x)].Add(0, x)
	ht.n++
	return true
}

func (ht *ChainedHashTable) Remove(x ch01.Comparable) bool {
	j := ht.hash(x)
	for i := 0; i < ht.t[j].Size(); i++ {
		y := ht.t[j].Get(i)
		if x == y {
			ht.t[j].Remove(i)
			ht.n--
			return true
		}
	}
	return false
}

func (ht *ChainedHashTable) Find(x ch01.Comparable) ch01.Comparable {
	j := ht.hash(x)
	for i := 0; i < ht.t[j].Size(); i++ {
		if x == ht.t[j].Get(i) {
			return ht.t[j].Get(i).(ch01.Comparable)
		}
	}
	return nil
}

func (ht *ChainedHashTable) hash(x ch01.Comparable) uint {
	h := (ht.z * x.HashCode()) >> (ht.w - ht.d)
	return h
}
