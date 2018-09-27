package ch5

import (
	"fmt"

	"github.com/tesujiro/OpenDataStructuresGo/ch1"
	"github.com/tesujiro/OpenDataStructuresGo/ch4"
)

type ChainedHashTable struct {
	t []ch1.List
	n int
	d uint32 //widht bit
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
	}
}

func (ht *ChainedHashTable) cap() int {
	return len(ht.t)
}

func (ht *ChainedHashTable) Size() int {
	return ht.n //TODO: sum
}

func (ht *ChainedHashTable) GetAll() []interface{} {
	slice := []interface{}{}
	for i := 0; i < ht.n; i++ {
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
	//fmt.Println("resize ht.n=", ht.n)
	var new []ch1.List
	if ht.n > 1 {
		new = make([]ch1.List, ht.n*2)
		ht.d++
	} else {
		new = make([]ch1.List, 1)
		ht.d = 1
	}
	for i := 0; i < ht.n; i++ {
		new[i] = ht.t[i]
	}
	for i := ht.n; i < len(new); i++ {
		//fmt.Println("add NewList() i:", i)
		new[i] = newList()
	}
	ht.t = new
}

func (ht *ChainedHashTable) Add(x interface{}) bool {
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

//func (ht *ChainedHashTable) Remove(x interface{}) interface{} {
func (ht *ChainedHashTable) Remove(x interface{}) bool {
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

func (ht *ChainedHashTable) Find(x interface{}) interface{} {
	j := ht.hash(x)
	for i := 0; i < ht.t[j].Size(); i++ {
		if x == ht.t[j].Get(i) {
			return ht.t[j].Get(i)
		}
	}
	return nil
}

const w = 64
const z = 41025416850000000

func hashCode(x interface{}) int {
	return x.(int)
}

func (ht *ChainedHashTable) hash(x interface{}) int {
	h := (z * hashCode(x)) >> (w - ht.d)
	fmt.Println("z*hashCode(x)=", z*hashCode(x))
	fmt.Printf("ht.d=%v hash=%v\n", ht.d, h)
	return h
}
