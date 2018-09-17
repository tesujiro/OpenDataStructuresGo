package ch2

import (
	"fmt"
)

const initSizeDQ = 8

type AD_Overridable interface {
	Resize()
}

type ArrayDeque struct {
	override AD_Overridable
	A        []interface{}
	start    int
	len      int
}

func NewArrayDeque() *ArrayDeque {
	ar := make([]interface{}, initSizeDQ)
	d := &ArrayDeque{
		A:     ar,
		start: 0,
		len:   0,
	}
	d.override = d
	return d
}

func (as *ArrayDeque) cap() int {
	return len(as.A)
}

func (as *ArrayDeque) Len() int {
	return as.len
}

func (as *ArrayDeque) GetAll() []interface{} {
	slice := []interface{}{}
	for i := 0; i < as.len; i++ {
		slice = append(slice, as.Get(i))
	}
	return slice
}

func (as *ArrayDeque) Print() {
	fmt.Printf("ArrayDeque(len:%v,cap:%v,start:%v)=%v\n", as.Len(), as.cap(), as.start, as.GetAll())
}

func (as *ArrayDeque) Get(i int) interface{} {
	i = i % as.len
	return as.A[(as.start+i)%as.cap()]
}

func (as *ArrayDeque) Set(i int, x interface{}) interface{} {
	i = i % as.len
	y := as.A[(as.start+i)%as.cap()]
	as.A[(as.start+i)%as.cap()] = x
	return y
}

func (as *ArrayDeque) Resize() {
	var new []interface{}
	if as.len > 1 {
		new = make([]interface{}, as.len*2)
	} else {
		new = make([]interface{}, 1)
	}
	for i := 0; i < as.len; i++ {
		new[i] = as.Get(i)
	}
	as.start = 0
	as.A = new
}

func (as *ArrayDeque) Add(i int, x interface{}) {
	i = i % (as.len + 1)
	if as.len+1 > as.cap() {
		as.override.Resize()
	}
	if i < as.len/2 {
		if as.start == 0 {
			as.start = len(as.A) - 1
		} else {
			as.start = as.start - 1
		}
		for k := 0; k <= i-1; k++ {
			as.A[(as.start+k)%as.cap()] = as.A[(as.start+k+1)%as.cap()]
		}
	} else {
		for k := as.len; k > i; k-- {
			as.A[(as.start+k)%as.cap()] = as.A[(as.start+k-1)%as.cap()]
		}
	}
	as.A[(as.start+i)%as.cap()] = x
	as.len += 1
}

func (as *ArrayDeque) Remove(i int) interface{} {
	i = i % as.len

	x := as.A[(as.start+i)%as.cap()]
	if i < as.len/2 {
		for k := i; k > 0; k-- {
			as.A[(as.start+k)%as.cap()] = as.A[(as.start+k-1)%as.cap()]
		}
		as.start = (as.start + 1) % as.cap()
	} else {
		for k := i; k <= as.len-1; k++ {
			as.A[(as.start+k)%as.cap()] = as.A[(as.start+k+1)%as.cap()]
		}
	}
	as.len -= 1
	if 3*as.len < as.cap() {
		as.override.Resize()
	}
	return x
}
