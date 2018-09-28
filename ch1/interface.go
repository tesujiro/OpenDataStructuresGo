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
	Add(x interface{}) bool
	Remove() interface{}
}

type Stack interface {
	Push(x interface{})
	Pop() interface{}
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
