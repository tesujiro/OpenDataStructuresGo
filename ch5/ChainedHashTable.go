package ch5

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/tesujiro/OpenDataStructuresGo/ch1"
	"github.com/tesujiro/OpenDataStructuresGo/ch4"
)

type ChainedHashTable struct {
	t []ch1.List
	n int    // number of element
	d uint32 // bit size of Hashtable
	w uint32 // bit size of int
	z uint   // random number
}

const initSize = 8
const initWidth = 3

func newList() ch1.List {
	return ch4.NewSkiplistList()
}

func NewChainedHashTable() *ChainedHashTable {
	t := make([]ch1.List, initSize)
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
	var new []ch1.List
	if ht.n > 1 {
		new = make([]ch1.List, ht.n*2)
		ht.d++
	} else {
		new = make([]ch1.List, 1)
		ht.d = 1
	}
	for i := 0; i < len(new); i++ {
		new[i] = newList()
	}
	for i := 0; i < ht.n; i++ {
		for j := 0; j < ht.t[i].Size(); j++ {
			x := ht.t[i].Get(j).(ch1.Comparable)
			new[ht.hash(x)].Add(0, x)
		}
	}
	ht.t = new
}

func (ht *ChainedHashTable) Add(x ch1.Comparable) bool {
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

func (ht *ChainedHashTable) Remove(x ch1.Comparable) bool {
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

func (ht *ChainedHashTable) Find(x ch1.Comparable) ch1.Comparable {
	j := ht.hash(x)
	for i := 0; i < ht.t[j].Size(); i++ {
		if x == ht.t[j].Get(i) {
			return ht.t[j].Get(i).(ch1.Comparable)
		}
	}
	return nil
}

func (ht *ChainedHashTable) hash(x ch1.Comparable) uint {
	h := (ht.z * x.HashCode()) >> (ht.w - ht.d)
	return h
}
