package ch11

import (
	"math/rand"

	"github.com/tesujiro/OpenDataStructuresGo/ch01"
)

func QuickSort(a []ch01.Comparable) {
	quickSort(a, 0, len(a))
}

func quickSort(a []ch01.Comparable, i, n int) {
	if n <= 1 {
		return
	}
	x := a[i+rand.Int()%n]
	p := i - 1
	j := i
	q := i + n
	// a[i..p]<x,  a[p+1..q-1]==x, a[q..i+n-1]>x

	for j < q {
		comp := a[j].Compare(x)
		switch {
		case comp < 0:
			p++
			a[j], a[p] = a[p], a[j]
			j++
		case comp > 0:
			q--
			a[j], a[q] = a[q], a[j]
		default:
			j++
		}
	}
	quickSort(a, i, p-i+1)
	quickSort(a, q, n-(q-i))
}
