package main

import (
	"fmt"
	"math/rand"

	"github.com/tesujiro/OpenDataStructuresGo/ch01"
	"github.com/tesujiro/OpenDataStructuresGo/ch02"
	"github.com/tesujiro/OpenDataStructuresGo/ch03"
	"github.com/tesujiro/OpenDataStructuresGo/ch04"
	"github.com/tesujiro/OpenDataStructuresGo/ch05"
	"github.com/tesujiro/OpenDataStructuresGo/ch06"
	"github.com/tesujiro/OpenDataStructuresGo/ch07"
	"github.com/tesujiro/OpenDataStructuresGo/ch08"
	"github.com/tesujiro/OpenDataStructuresGo/ch11"
	"github.com/tesujiro/OpenDataStructuresGo/ch13"
)

type element int

func (e1 element) Compare(e2 ch01.Comparable) int {
	return int(e1 - e2.(element))
}

func (e1 element) HashCode() uint {
	return uint(e1)
}

/*
func (e1 element) Compare(e2 interface{}) int {
	switch e2.(type) {
	case element:
		return int(e1 - e2.(element))
	case int:
		return int(e1) - e2.(int)
	default:
		return int(e1)
	}
}
*/

func main() {

	fmt.Println("\nch02/ArrayQueue")
	checkQueue(ch02.NewArrayQueue())

	fmt.Println("\nch03/SLList")
	checkStack(ch03.NewSLList())
	checkQueue(ch03.NewSLList())

	//List
	fmt.Println("\nch02/ArrayStack")
	checkList(ch02.NewArrayStack())

	fmt.Println("\nch02/ArrayDeque")
	checkList(ch02.NewArrayDeque())

	fmt.Println("\nch02/DualArrayDeque")
	checkList(ch02.NewDualArrayDeque())

	fmt.Println("\nch02/RootishArrayStack")
	checkList(ch02.NewRootishArrayStack())

	fmt.Println("\nch03/DLList")
	checkList(ch03.NewDLList())

	fmt.Println("\nch03/SEList")
	checkList(ch03.NewSEList())

	fmt.Println("\nch04/SkiplistList")
	checkList(ch04.NewSkiplistList())

	//SSet
	fmt.Println("\nch04/SkiplistSSet")
	checkSSet(ch04.NewSkiplistSSet())

	fmt.Println("\nch05/ChainedHashTable")
	checkSSet(ch05.NewChainedHashTable())

	fmt.Println("\nch05/LinearHashTable")
	checkSSet(ch05.NewLinearHashTable())

	fmt.Println("\nch06/BinaryTree")
	checkSSet(ch06.NewBinaryTree())

	fmt.Println("\nch07/Treap")
	checkSSet(ch07.NewTreap())

	fmt.Println("\nch08/ScapegoatTree")
	checkSSet(ch08.NewScapegoatTree())

	fmt.Println("\nch13/BinaryTrie")
	checkSSet(ch13.NewBinaryTrie())

	fmt.Println("\nch13/XFastTrie")
	checkSSet(ch13.NewXFastTrie())

	// Sort
	fmt.Println("\nch11/MergeSort")
	checkSort(ch11.MergeSort)

	fmt.Println("\nch11/QuickSort")
	checkSort(ch11.QuickSort)

	fmt.Println("\nch11/CountingSort")
	checkCSort(ch11.CountingSort)

	fmt.Println("\nch11/RadixSort")
	checkCSort(ch11.RadixSort)

}

func checkList(s ch01.List) {

	s.Print()

	s.Add(0, 10)
	s.Print()
	//s.Add(0, 10)
	s.Add(0, 40)
	s.Print()

	s.Add(2, 50)
	s.Print()
	s.Add(3, 60)
	s.Print()
	s.Add(0, 3)
	s.Print()
	s.Add(0, 2)
	s.Add(0, 1)
	s.Add(0, 0)

	var i int
	var v interface{}
	i = 1
	v = s.Get(i)
	i = 2
	v = 10
	s.Set(i, v)
	s.Print()

	s.Remove(0)
	s.Print()
	s.Remove(2)
	s.Print()
	s.Remove(3)
	s.Print()
	s.Remove(0)
	s.Print()
}

func checkStack(s ch01.Stack) {
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Print()
	fmt.Printf("Pop()=%v\n", s.Pop())
	fmt.Printf("Pop()=%v\n", s.Pop())
	fmt.Printf("Pop()=%v\n", s.Pop())
	s.Print()
}

func checkQueue(s ch01.Queue) {
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Print()
	fmt.Printf("Remove()=%v\n", s.Remove())
	fmt.Printf("Remove()=%v\n", s.Remove())
	fmt.Printf("Remove()=%v\n", s.Remove())
	s.Print()
}

func checkUSet(s ch01.USet) {
	s.Print()
	s.Add(10)
	s.Print()
	for i := 200; i < 2000; i += 100 {
		if !s.Add(i) {
			fmt.Println("Add ERROR! i:", i)
		}
	}
	s.Print()
	fmt.Println("Size:", s.Size())
	fmt.Printf("Find(element(20))=%v\n", s.Find(20))
	fmt.Printf("Find(element(123))=%v\n", s.Find(123))
	s.Remove(10)
	s.Print()
	s.Remove(30)
	s.Print()
	s.Remove(20)
	s.Print()
}

func checkSSet(s ch01.SSet) {
	s.Print()
	s.Add(element(10))
	s.Print()
	for i := 20; i <= 200; i += 10 {
		s.Add(element(i))
	}
	s.Print()
	fmt.Println("Size:", s.Size())
	fmt.Printf("Find(element(5))=%v\n", s.Find(element(5)))
	fmt.Printf("Find(element(10))=%v\n", s.Find(element(10)))
	fmt.Printf("Find(element(120))=%v\n", s.Find(element(120)))
	fmt.Printf("Find(element(123))=%v\n", s.Find(element(123)))
	fmt.Printf("Find(element(205))=%v\n", s.Find(element(205)))
	s.Remove(element(10))
	s.Print()
	s.Remove(element(30))
	s.Print()
	s.Remove(element(20))
	s.Print()
}

func checkSort(f ch01.SortFunc) {
	s := []ch01.Comparable{}
	for i := 300; i > 0; i -= 10 {
		s = append(s, element(i))
	}
	fmt.Printf("s=%v\n", s)
	f(s)
	fmt.Printf("s=%v\n", s)
}

func checkCSort(f ch01.CountingSortFunc) {
	k := 256 * 256
	s := []int{}
	for i := 0; i < 100; i++ {
		s = append(s, rand.Intn(k))
	}
	fmt.Printf("s=%v\n", s)
	f(&s, k)
	fmt.Printf("s=%v\n", s)
}
