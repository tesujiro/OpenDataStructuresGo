package ch1

type Printable interface {
	GetAll() []interface{}
	Print()
}

type List interface {
	Printable
	Size() int
	Get(i int) interface{}
	Set(i int, v interface{}) interface{}
	Add(i int, v interface{})
	Remove(i int) interface{}
}

type Queue interface {
	Printable
	Add(x interface{}) bool
	Remove() interface{}
}

type Stack interface {
	Push(x interface{})
	Pop() interface{}
	Print()
}

type USet interface {
	Printable
	Size() int
	Add(x interface{}) bool
	Remove(x interface{}) interface{}
	Find(x interface{}) interface{}
}

type SSet interface {
	USet
	Compare(x, y interface{}) int
}
