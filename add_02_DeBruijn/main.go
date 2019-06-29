package main

import (
	"flag"
	"fmt"

	"github.com/pkg/profile"
)

type D struct {
	k         int // number of characters
	n         int // length of characters
	startWith rune
	A         []rune
	V         [][]rune
}

func newD(k, n int, zero bool) *D {
	d := D{k: k, n: n}
	if zero {
		d.startWith = '0'
	} else {
		d.startWith = 'a'
	}
	d.A = d.alph()
	d.V = d.combi()
	return &d
}

func (d *D) alph() []rune {
	a := []rune{}
	for i := 0; i < d.k; i++ {
		if d.startWith == '0' && i == 10 {
			d.startWith = 'A' - 10
		}
		a = append(a, d.startWith+rune(i))
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

// De Bruijn sequences built using brute force
func (d *D) _seq(cur string, p int, checklist map[string]bool) (string, bool) {
	if p == len(d.V) {
		return cur, true
	}
	if p > len(d.V)-d.n {
		key := (cur + cur)[p : p+d.n]
		if _, ok := checklist[key]; !ok {
			checklist[key] = true
			if ret, ok := d._seq(cur, p+1, checklist); ok {
				return ret, true
			} else {
				delete(checklist, key)
			}
		}
	} else {
		for _, r := range d.A {
			newCur := cur + string(r)
			if len(newCur) < d.n {
				return d._seq(newCur, 0, checklist)
			}
			key := string(newCur)[p:]
			if _, ok := checklist[key]; !ok {
				checklist[key] = true
				if ret, ok := d._seq(newCur, p+1, checklist); ok {
					return ret, true
				} else {
					delete(checklist, key)
				}
			}
		}
	}
	return "", false
}

func (d *D) seq() (string, bool) {
	checklist := make(map[string]bool)
	return d._seq("", 0, checklist)
}

func dump(a []rune) string {
	var ret string
	for i, r := range a {
		if i%8 == 0 {
			ret = fmt.Sprintf("%s ", ret)
		}
		ret = fmt.Sprintf("%s%s", ret, string(r))
	}
	return ret
}

func main() {
	var (
		k    = flag.Int("k", 2, "alphabets")
		n    = flag.Int("n", 3, "word length")
		zero = flag.Bool("z", false, "alphabets start with zero")
	)
	flag.Parse()
	d := newD(*k, *n, *zero)
	fmt.Printf("k: %v\n", d.k)
	fmt.Printf("n: %v\n", d.n)
	fmt.Printf("alphabets: %v\n", dump(d.A))
	fmt.Printf("combi: %v\n", len(d.V))
	defer profile.Start(profile.ProfilePath(".")).Stop()
	seq, ok := d.seq()
	if !ok {
		fmt.Println("Result is not OK!")
	}
	fmt.Println("De Bruijn sequence:", dump([]rune(seq)))
	fmt.Println("De Bruijn sequence length:", len(string(seq)))
}
