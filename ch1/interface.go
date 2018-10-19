package ch1

type Helper interface {
	GetAll() []interface{}
	Print()
}

type List interface {
	Helper
	Size() int
	Get(i int) interface{}
	Set(i int, v interface{}) interface{}
	Add(i int, v interface{})
	Remove(i int) interface{}
}

type Queue interface {
	Helper
	Add(x interface{}) bool // Enqueue()
	Remove() interface{}    // Dequeue()
}

type PriorityQueue interface {
	Helper
	Add(x Comparable) bool
	Remove() Comparable // DeleteMin()
}

type Stack interface {
	Push(x interface{}) // Add(x)
	Pop() interface{}   // Remove()
	Print()
}

type USet interface {
	Helper
	Size() int
	Add(x interface{}) bool
	Remove(x interface{}) bool
	Find(x interface{}) interface{}
}

type Comparable interface {
	Compare(Comparable) int
	//Compare(interface{}) int
	HashCode() uint
}

type SSet interface {
	Helper
	Size() int
	Add(x Comparable) bool
	Remove(x Comparable) bool
	Find(x Comparable) Comparable
}

type SortFunc func([]Comparable)
type CountingSortFunc func(*[]int, int)
