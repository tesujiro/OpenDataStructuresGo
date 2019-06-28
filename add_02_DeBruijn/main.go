package main

import (
	"flag"
	"fmt"
)

type D struct {
	k int // number of characters
	n int // length of characters
	A []rune
	V [][]rune
}

func newD(k, n int) *D {
	d := D{k: k, n: n}
	d.A = d.alph()
	d.V = d.combi()
	return &d
}

func (d *D) alph() []rune {
	a := []rune{}
	for i := 0; i < d.k; i++ {
		a = append(a, rune('a'+i))
	}
	return a
}

func (d *D) combi() [][]rune {
	return d._combi([]rune{}, d.n, [][]rune{})
}

func (d *D) _combi(cur []rune, n int, ret [][]rune) [][]rune {
	if n == 0 {
		word := make([]rune, len(cur))
		copy(word, cur)
		return append(ret, word)
	}
	for _, r := range d.A {
		ret = d._combi(append(cur, r), n-1, ret)
	}
	return ret
}

func (d *D) check(cur []rune, checklist map[string]bool) (bool, string) {
	//fmt.Println("cur=", string(cur))
	if len(cur) < d.n {
		return true, ""
	}

	var str string
	for i := len(cur) - d.n; i < len(cur); i++ {
		str = (string(cur) + string(cur))[i : i+d.n]
		//fmt.Printf("str=%v ", str)
		if _, ok := checklist[str]; ok {
			//fmt.Printf("->false checklist=%v\n", checklist)
			return false, ""
		}
		//fmt.Printf("->true checklist=%v\n", checklist)
		if len(cur) != len(d.V) {
			break
		}
	}
	return true, str
}

func (d *D) _seq(cur []rune, checklist map[string]bool) ([]rune, bool) {
	if len(cur) == len(d.V) {
		return cur, true
	}
	for _, r := range d.A {
		newCur := append(cur, r)
		if ok, key := d.check(newCur, checklist); ok {
			checklist[key] = true
			if ret, ok := d._seq(newCur, checklist); ok {
				return ret, true
			} else {
				delete(checklist, key)
			}
		}
	}
	return nil, false
}

func (d *D) seq() ([]rune, bool) {
	checklist := make(map[string]bool)
	return d._seq([]rune{}, checklist)
}

func main() {
	var (
		k = flag.Int("k", 2, "alphabets")
		n = flag.Int("n", 3, "word length")
	)
	flag.Parse()
	d := newD(*k, *n)
	fmt.Printf("k: %v\n", d.k)
	fmt.Printf("n: %v\n", d.n)
	fmt.Printf("alphabets: %v\n", string(d.A))
	seq, ok := d.seq()
	fmt.Println("ok:", ok)
	fmt.Println("De Bruijn sequence:", string(seq))
}
