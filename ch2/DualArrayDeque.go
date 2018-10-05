package ch2

import (
	"fmt"
)

type DualArrayDeque struct {
	front *ArrayStack
	back  *ArrayStack
}

func NewDualArrayDeque() *DualArrayDeque {
	return &DualArrayDeque{
		front: NewArrayStack(),
		back:  NewArrayStack(),
	}
}

func (as *DualArrayDeque) cap() int {
	return as.front.cap() + as.back.cap()
}

func (as *DualArrayDeque) Size() int {
	return as.front.Size() + as.back.Size()
}

func (as *DualArrayDeque) GetAll() []interface{} {
	slice := []interface{}{}
	for i := 0; i < as.front.n; i++ {
		slice = append(slice, as.front.Get(as.front.n-i-1))
	}
	for i := 0; i < as.back.n; i++ {
		slice = append(slice, as.back.Get(i))
	}
	return slice
}

func (as *DualArrayDeque) Print() {
	fmt.Printf("DualArrayDeque(n:%v,cap:%v)=%v\n", as.Size(), as.cap(), as.GetAll())
}

func (as *DualArrayDeque) Get(i int) interface{} {
	i = i % as.Size()
	if i < as.front.Size() {
		return as.front.Get(as.front.Size() - i - 1)
	} else {
		return as.back.Get(i - as.front.Size())
	}
}

func (as *DualArrayDeque) Set(i int, x interface{}) interface{} {
	i = i % as.Size()
	if i < as.front.Size() {
		return as.front.Set(as.front.Size()-i-1, x)
	} else {
		return as.back.Set(i-as.front.Size(), x)
	}
}

func (as *DualArrayDeque) balance() {
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	if 3*as.front.Size() < as.back.Size() ||
		3*as.back.Size() < as.front.Size() {
	}
	n := as.Size()
	nf := n / 2
	af := make([]interface{}, max(2*nf, 1))
	for i := 0; i < nf; i++ {
		af[nf-i-1] = as.Get(i)
	}
	nb := n - nf
	ab := make([]interface{}, max(2*nb, 1))
	for i := 0; i < nb; i++ {
		ab[i] = as.Get(nf + i)
	}
	as.front.a = af
	as.front.n = nf
	as.back.a = ab
	as.back.n = nb
}

func (as *DualArrayDeque) Add(i int, x interface{}) {
	i = i % (as.Size() + 1)
	if i < as.front.Size() {
		as.front.Add(as.front.Size()-i, x)
	} else {
		as.back.Add(i-as.front.Size(), x)
	}
	as.balance()
}

func (as *DualArrayDeque) Remove(i int) interface{} {
	i = i % as.Size()
	var x interface{}
	if i < as.front.Size() {
		x = as.front.Remove(as.front.Size() - i - 1)
	} else {
		x = as.back.Remove(i - as.front.Size())
	}
	as.balance()
	return x
}
