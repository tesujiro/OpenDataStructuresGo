package ch2

import (
	"fmt"
)

const initSizeDQ = 8

type AD_Overridable interface {
	Resize()
}

type ArrayDeque struct {
	Override AD_Overridable
	A        []interface{}
	j        int
	n        int
}

func NewArrayDeque() *ArrayDeque {
	ar := make([]interface{}, initSizeDQ)
	d := &ArrayDeque{
		A: ar,
		j: 0,
		n: 0,
	}
	d.Override = d
	return d
}

func (as *ArrayDeque) cap() int {
	return len(as.A)
}

func (as *ArrayDeque) Size() int {
	return as.n
}

func (as *ArrayDeque) GetAll() []interface{} {
	slice := []interface{}{}
	for i := 0; i < as.n; i++ {
		slice = append(slice, as.Get(i))
	}
	return slice
}

func (as *ArrayDeque) Print() {
	fmt.Printf("ArrayDeque(n:%v,cap:%v,j:%v)=%v\n", as.Size(), as.cap(), as.j, as.GetAll())
}

func (as *ArrayDeque) Get(i int) interface{} {
	i = i % as.n
	return as.A[(as.j+i)%as.cap()]
}

func (as *ArrayDeque) Set(i int, x interface{}) interface{} {
	i = i % as.n
	y := as.A[(as.j+i)%as.cap()]
	as.A[(as.j+i)%as.cap()] = x
	return y
}

func (as *ArrayDeque) Resize() {
	var new []interface{}
	if as.n > 1 {
		new = make([]interface{}, as.n*2)
	} else {
		new = make([]interface{}, 1)
	}
	for i := 0; i < as.n; i++ {
		new[i] = as.Get(i)
	}
	as.j = 0
	as.A = new
}

func (as *ArrayDeque) Add(i int, x interface{}) {
	i = i % (as.n + 1)
	if as.n+1 > as.cap() {
		as.Override.Resize()
	}
	if i < as.n/2 {
		if as.j == 0 {
			as.j = as.cap() - 1
		} else {
			as.j = as.j - 1
		}
		for k := 0; k <= i-1; k++ {
			as.A[(as.j+k)%as.cap()] = as.A[(as.j+k+1)%as.cap()]
		}
	} else {
		for k := as.n; k > i; k-- {
			as.A[(as.j+k)%as.cap()] = as.A[(as.j+k-1)%as.cap()]
		}
	}
	as.A[(as.j+i)%as.cap()] = x
	as.n++
}

func (as *ArrayDeque) Remove(i int) interface{} {
	i = i % as.n

	x := as.A[(as.j+i)%as.cap()]
	if i < as.n/2 {
		for k := i; k > 0; k-- {
			as.A[(as.j+k)%as.cap()] = as.A[(as.j+k-1)%as.cap()]
		}
		as.j = (as.j + 1) % as.cap()
	} else {
		for k := i; k <= as.n-1; k++ {
			as.A[(as.j+k)%as.cap()] = as.A[(as.j+k+1)%as.cap()]
		}
	}
	as.n--
	if 3*as.n < as.cap() {
		as.Override.Resize()
	}
	return x
}
