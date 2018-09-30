package ch5

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/tesujiro/OpenDataStructuresGo/ch1"
)

type LinearHashTable struct {
	t   []ch1.Comparable
	del []bool // delete flag for values in T
	n   int    // number of values in T
	q   int    // number of non-null entries in T
	d   uint32 // t.length = 2^d
	w   uint32 // bit size of int
	z   uint   // random number
}

func NewLinearHashTable() *LinearHashTable {
	rand.Seed(time.Now().UnixNano())
	return &LinearHashTable{
		t:   make([]ch1.Comparable, 1),
		del: make([]bool, 1),
		n:   0,
		q:   0,
		d:   0,
		w:   strconv.IntSize,
		z:   uint(rand.Int()) * 98767,
	}
}

func (ht *LinearHashTable) cap() int {
	return len(ht.t)
}

func (ht *LinearHashTable) Size() int {
	return ht.n // number of element
}

func (ht *LinearHashTable) GetAll() []interface{} {
	slice := []interface{}{}
	for i := 0; i < ht.cap(); i++ {
		//fmt.Printf("ht.cap=%v i=%v\n", ht.cap(), i)
		if !ht.del[i] && ht.t[i] != nil {
			//fmt.Printf("i=%v ht.t[i]=%v\n", i, ht.t[i])
			slice = append(slice, ht.t[i])
		}
	}
	return slice
}

func (ht *LinearHashTable) Print() {
	fmt.Printf("LinearHashTable(n:%v,cap:%v)=%v\n", ht.Size(), ht.cap(), ht.GetAll())
}

func (ht *LinearHashTable) hash(x ch1.Comparable) int {
	h := (ht.z * x.HashCode()) >> (ht.w - ht.d)
	return int(h)
}

func (ht *LinearHashTable) Find(x ch1.Comparable) ch1.Comparable {
	i := ht.hash(x)
	//ht.Print()
	//fmt.Printf("x=%v hash(x)=%v cap=%v ht.d=%v\n", x, i, ht.cap(), ht.d)
	for ht.t[i] != nil {
		if !ht.del[i] && ht.t[i].HashCode() == x.HashCode() {
			//fmt.Printf("Found %v\n", x)
			return ht.t[i]
		}
		i = (i + 1) % ht.cap()
		//if i == ht.hash(x) {
		//break
		//}
	}
	return nil
}

func (ht *LinearHashTable) Add(x ch1.Comparable) bool {
	if ht.Find(x) != nil {
		//fmt.Printf("Add Find(%v)!=nil\n", x)
		return false
	}
	if 2*(ht.q+1) > ht.cap() {
		ht.resize()
	}
	i := ht.hash(x)
	for ht.t[i] != nil && !ht.del[i] {
		i = (i + 1) % ht.cap()
	}
	if ht.t[i] == nil {
		ht.q++
	}
	ht.n++
	ht.t[i] = x
	//fmt.Printf("Add i=%v ht.t[i]=%v\n", i, ht.t[i])
	return true
}

func (ht *LinearHashTable) Remove(x ch1.Comparable) bool {
	i := ht.hash(x)
	for ht.t[i] != nil {
		y := ht.t[i]
		if !ht.del[i] && x == y {
			ht.del[i] = true
			ht.n--
			if 8*ht.n < ht.cap() {
				ht.resize()
			}
			return true
		}
		i = (i + 1) % ht.cap()
	}
	return false
}

func (ht *LinearHashTable) resize() {
	d := uint32(1)
	for 1<<d < 3*ht.n {
		d++
	}
	tnew := make([]ch1.Comparable, 1<<d)
	ht.d = d
	ht.q = ht.n
	for k := 0; k < ht.cap(); k++ {
		if ht.t[k] != nil && !ht.del[k] {
			i := ht.hash(ht.t[k])
			for tnew[i] != nil {
				i = (i + 1) % len(tnew)
			}
			tnew[i] = ht.t[k]
		}
	}
	ht.del = make([]bool, 1<<d)
	ht.t = tnew
}
