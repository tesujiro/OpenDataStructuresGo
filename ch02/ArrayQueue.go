package ch02

import "fmt"

type ArrayQueue struct {
	a []interface{}
	j int
	n int
}

const initArrayQueueSize = 8

func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{
		a: make([]interface{}, initArrayQueueSize),
		j: 0,
		n: 0,
	}
}

func (l *ArrayQueue) cap() int {
	return cap(l.a)
}

func (l *ArrayQueue) Size() int {
	return l.n
}

func (l *ArrayQueue) GetAll() []interface{} {
	s := []interface{}{}
	for i := l.j; i < l.j+l.n; i++ {
		s = append(s, l.a[i%l.n])
	}
	return s
}

func (l *ArrayQueue) Print() {
	fmt.Printf("list=%v\n", l.GetAll())
}

func (l *ArrayQueue) resize() {
	var new []interface{}
	if l.n > 1 {
		new = make([]interface{}, l.n*2)
	} else {
		new = make([]interface{}, 1)
	}
	for i := 0; i < l.n; i++ {
		new[i] = l.a[(l.j+i)%l.cap()]
	}
	l.a = new
	l.j = 0
}

func (l *ArrayQueue) Add(x interface{}) bool {
	if l.n+1 > l.cap() {
		l.resize()
	}
	l.a[(l.j+l.n)%l.cap()] = x
	l.n++
	return true
}

func (l *ArrayQueue) Remove() interface{} {
	x := l.a[l.j]
	l.j = (l.j + 1) % l.cap()
	l.n--
	if l.cap() >= l.n*3 {
		l.resize()
	}
	return x
}
