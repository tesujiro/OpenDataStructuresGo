package ch11

import "fmt"

func RadixSort(a *[]int, k int) {
	d := uint(8)
	w := uint(32)
	for p := uint(0); p < w/d; p++ {
		c := make([]int, 1<<d)
		b := make([]int, len(*a))
		// the next three for loops implement counting-sort
		for i := 0; i < len(*a); i++ {
			c[(uint((*a)[i])>>d*p)&((1<<d)-1)]++
		}
		for i := 1; i < 1<<d; i++ {
			//fmt.Printf("i=%v 1<<d=%v\n", i, 1<<d)
			c[i] += c[i-1]
		}
		//b := make([]int, len(*a))
		for i := len(*a) - 1; i >= 0; i-- {
			//fmt.Printf("i=%v \n", i)
			//fmt.Printf("index=%v n", (uint((*a)[i])>>d*p)&((1<<d)-1))
			//fmt.Printf("index=%v n", uint((*a)[i])>>d*p)
			//fmt.Printf("index=%v \n", d*p)
			//fmt.Printf("index=%v \n", uint((*a)[i]))
			j := (uint((*a)[i]) >> d * p) & ((1 << d) - 1)
			c[j]--
			//fmt.Printf("j=%v c[j]=%v\n", j, c[j])
			b[c[j]] = (*a)[i]
		}
		fmt.Printf("b=%v\n", b)
		*a = b
	}

}
