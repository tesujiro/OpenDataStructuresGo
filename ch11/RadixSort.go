package ch11

import "math"

func RadixSort(a *[]int, k int) {
	d := uint(8)
	bytes := math.Ceil(math.Logb(float64(k)) / float64(d))
	w := uint(bytes) * d
	//w := uint(32)
	for p := uint(0); p < w/d; p++ {
		c := make([]int, 1<<d)
		b := make([]int, len(*a))
		// the next three for loops implement counting-sort
		for i := 0; i < len(*a); i++ {
			c[(uint((*a)[i])>>(d*p))&((1<<d)-1)]++
		}
		for i := 1; i < 1<<d; i++ {
			c[i] += c[i-1]
		}
		for i := len(*a) - 1; i >= 0; i-- {
			j := (uint((*a)[i]) >> (d * p)) & ((1 << d) - 1)
			c[j]--
			b[c[j]] = (*a)[i]
		}
		*a = b
	}

}
