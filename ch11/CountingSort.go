package ch11

//func CountingSort(a []ch1.Comparable,k int) {
func CountingSort(a *[]int, k int) {
	c := make([]int, k)
	for i := 0; i < len(*a); i++ {
		c[(*a)[i]]++
	}
	for i := 1; i < k; i++ {
		c[i] += c[i-1]
	}
	b := make([]int, len(*a))
	for i := len(*a) - 1; i >= 0; i-- {
		//fmt.Printf("i=%v\n", i)
		c[(*a)[i]]--
		b[c[(*a)[i]]] = (*a)[i]
	}
	*a = b
}
