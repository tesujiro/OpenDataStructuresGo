package ch02

import (
	"fmt"
	"math"
)

type RootishArrayStack struct {
	blocks *ArrayStack
	n      int
}

func NewRootishArrayStack() *RootishArrayStack {
	blocks := NewArrayStack()
	blocks.Set(0, make([]interface{}, 0, 1))
	return &RootishArrayStack{
		blocks: blocks,
		n:      0,
	}
}

func (as *RootishArrayStack) cap() int {
	c := 0
	for i := 0; i < as.blocks.Size(); i++ {
		c += cap(as.blocks.Get(i).([]interface{}))
	}
	return c
}

func (as *RootishArrayStack) Size() int {
	return as.n
}

func (as *RootishArrayStack) GetAll() []interface{} {
	slice := []interface{}{}
	for i := 0; i < as.blocks.Size(); i++ {
		block := as.blocks.Get(i).([]interface{})
		for j := 0; j < len(block); j++ {
			slice = append(slice, as.blocks.Get(i).([]interface{})[j])
		}
	}
	return slice
}

func (as *RootishArrayStack) Print() {
	fmt.Printf("RootishArrayStack(n:%v,cap:%v)=%v\n", as.Size(), as.cap(), as.GetAll())
}

func i2b(i int) int {
	db := (-3.0 + math.Sqrt(float64(9+8*i))) / 2.0
	return int(math.Ceil(db))
}

func (as *RootishArrayStack) Get(i int) interface{} {
	b := i2b(i)
	j := i - b*(b+1)/2
	return as.blocks.Get(b).([]interface{})[j]
}

func (as *RootishArrayStack) Set(i int, x interface{}) interface{} {
	b := i2b(i)
	j := i - b*(b+1)/2
	block := as.blocks.Get(b).([]interface{})
	if j < len(block) {
		y := block[j]
		block[j] = x
		//as.blocks.Set(b, block)
		return y
	} else {
		//fmt.Printf("block extend Add(%v) cap=%v j=%v\n", x, as.cap(), j)
		//as.blocks.Set(b, append(block, x))
		block = append(block, x)
		as.blocks.Set(b, block)
		return nil
	}
}

func (as *RootishArrayStack) grow() {
	newBlock := make([]interface{}, 0, as.blocks.Size()+1)
	as.blocks.Add(as.blocks.Size(), newBlock)
}

func (as *RootishArrayStack) Add(i int, x interface{}) {
	r := as.blocks.Size()
	if r*(r+1)/2 < as.n+1 {
		as.grow()
	}
	as.n++
	for j := as.n - 1; j > i; j-- {
		as.Set(j, as.Get(j-1))
	}
	//fmt.Printf("Add(%v) cap=%v\n", x, as.cap())
	as.Set(i, x)
}

func (as *RootishArrayStack) shrink() {
	as.blocks.Remove(as.blocks.Size() - 1)
	/*
		for r := as.blocks.Size(); r > 0 && (r-2)*(r-1)/2 >= as.n; r-- {
			as.blocks.Remove(as.blocks.Size() - 1)
		}
	*/
}

func (as *RootishArrayStack) Remove(i int) interface{} {
	x := as.Get(i)
	for j := i; j < as.n-1; j++ {
		as.Set(j, as.Get(j+1))
	}
	as.n--
	lastBlock := as.blocks.Get(as.blocks.Size() - 1).([]interface{})
	if len(lastBlock) > 1 {
		as.blocks.Set(as.blocks.Size()-1, lastBlock[:len(lastBlock)-1])
	} else {
		as.shrink()
	}
	/*
		r := as.blocks.Size()
		fmt.Printf("r=%v n=%v\n", r, as.n)
		for r > 0 && (r-2)*(r-1)/2 >= as.n {
			fmt.Println("shrink")
			as.shrink()
		}
	*/
	return x
}
