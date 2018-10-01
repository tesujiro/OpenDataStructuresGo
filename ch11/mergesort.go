package ch11

import (
	"github.com/tesujiro/OpenDataStructuresGo/ch1"
)

func MergeSort(a []ch1.Comparable) {
	if len(a) <= 1 {
		return
	}
	a0 := make([]ch1.Comparable, len(a[:len(a)/2]))
	copy(a0, a[:len(a)/2])
	a1 := make([]ch1.Comparable, len(a[len(a)/2:]))
	copy(a1, a[len(a)/2:])
	MergeSort(a0)
	MergeSort(a1)
	merge(a0, a1, a)
}

func merge(a0, a1, a []ch1.Comparable) {
	var i0, i1 int
	for i := 0; i < len(a); i++ {
		switch {
		case i0 == len(a0):
			a[i] = a1[i1]
			i1++
		case i1 == len(a1):
			a[i] = a0[i0]
			i0++
		case a0[i0].Compare(a1[i1]) <= 0:
			a[i] = a0[i0]
			i0++
		default:
			a[i] = a1[i1]
			i1++
		}
	}
}
