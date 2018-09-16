package ch2

import "fmt"

type ArrayQueue struct {
	array []interface{}
	start int
	len   int
}

const initArrayQueueSize = 8

func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{
		array: make([]interface{}, initArrayQueueSize),
		start: 0,
		len:   0,
	}
}

func (l *ArrayQueue) cap() int {
	return cap(l.array)
}

func (l *ArrayQueue) Len() int {
	return l.len
}

func (l *ArrayQueue) GetAll() []interface{} {
	s := []interface{}{}
	for i := l.start; i < l.start+l.len; i++ {
		s = append(s, l.array[i%l.len])
	}
	return s
}

func (l *ArrayQueue) Print() {
	fmt.Printf("list=%v\n", l.GetAll())
}

func (l *ArrayQueue) resize() {
	var new []interface{}
	if l.len > 1 {
		new = make([]interface{}, l.len*2)
	} else {
		new = make([]interface{}, 1)
	}
	for i := 0; i < l.len; i++ {
		new[i] = l.array[(l.start+i)%l.cap()]
	}
	l.array = new
	l.start = 0
}

func (l *ArrayQueue) Add(x interface{}) bool {
	if l.len+1 > l.cap() {
		l.resize()
	}
	l.array[(l.start+l.len)%l.cap()] = x
	l.len += 1
	return true
}

func (l *ArrayQueue) Remove() interface{} {
	x := l.array[l.start]
	l.start = (l.start + 1) % l.cap()
	l.len -= 1
	if l.cap() >= l.len*3 {
		l.resize()
	}
	return x
}
