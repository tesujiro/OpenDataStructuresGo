package main

import (
	"flag"
	"fmt"
	"math"
	"os"

	"github.com/pkg/profile"
)

type D struct {
	k         int // number of characters
	n         int // length of characters
	startWith rune
	A         []rune
	V         [][]rune
	G         map[string][]string
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
	d.G = d.makeGraph()
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

func (d *D) seq2hex(seq string) string {
	bits := int(math.Log2(float64(d.k)))
	if float64(d.k) != math.Exp2(float64(bits)) {
		return "k is not 2^x"
	}

	hex := "0x"
	index := make(map[rune]int, d.k)
	for i, r := range d.A {
		index[r] = i
	}
	var tmp uint32
	for i, r := range seq {
		tmp = tmp*uint32(d.k) + uint32(index[r])
		if (i+1)*bits%32 == 0 {
			hex = fmt.Sprintf("%s%8.8x", hex, tmp)
			tmp = 0
		}
	}
	if bits*len(seq)%32 != 0 {
		var width string
		l := int(math.Ceil(float64(bits*len(seq)) / 4.0))
		if l != 0 {
			width = fmt.Sprintf("%d", l)
		} else {
			width = fmt.Sprintf("%d", 4)
		}
		hex = fmt.Sprintf("%s%"+width+"."+width+"x", hex, tmp)
	}
	return hex
}

func (d *D) bitPosition(s string) []int {
	tape := s + s
	var pos2value []int
	index := make(map[rune]int, d.k)
	for i, r := range d.A {
		index[r] = i
	}
	//fmt.Println("index:", index)
	for i := 0; i < len(s); i++ {
		pos := 0
		for j := 0; j < d.n; j++ {
			pos = pos*d.k + index[rune(tape[i+j])]
		}
		//fmt.Println("pos:", pos)
		pos2value = append(pos2value, pos)
	}
	value2pos := make([]int, len(pos2value))
	for pos, value := range pos2value {
		value2pos[value] = pos
	}
	return value2pos
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

func dumpPath(p []string) string {
	ret := p[0]
	length := len(p[0])
	for i := 1; i < len(p)-1; i++ {
		ret += p[i][len(p[i])-1:]
		length += 1
		if length%8 == 0 {
			ret += " "
		}
		//fmt.Println("ret:", ret)
		//ret += p[i]
	}
	return ret
}

func main() {
	var (
		k     = flag.Int("k", 2, "alphabets")
		n     = flag.Int("n", 3, "word length")
		zero  = flag.Bool("z", false, "alphabets start with zero")
		blute = flag.Bool("b", false, "make sequence with blute force")
	)
	flag.Parse()
	d := newD(*k, *n, *zero)
	fmt.Printf("k: %v\n", d.k)
	fmt.Printf("n: %v\n", d.n)
	fmt.Printf("alphabets: %v\n", dump(d.A))
	fmt.Printf("combi: %v\n", len(d.V))
	defer profile.Start(profile.ProfilePath(".")).Stop()
	if *blute {
		fmt.Println("[Blute Force]")
		seq, ok := d.seq()
		if !ok {
			fmt.Println("Result is not OK!")
		}
		fmt.Println("De Bruijn sequence:", dump([]rune(seq)))
		fmt.Println("De Bruijn sequence(HEX):", d.seq2hex(seq))
		fmt.Println("De Bruijn sequence length:", len(string(seq)))
		pos := d.bitPosition(seq)
		fmt.Printf("De Bruijn bit position:")
		for _, v := range pos {
			fmt.Printf("%v, ", v)
		}
		fmt.Printf("\n")
	} else {

		fmt.Println("[De Bruijn Graph]")
		fmt.Printf("Graph:%v\n", d.G)
		path, ok := d.SeqGraph()
		if !ok {
			fmt.Println("Result is not OK!")
			os.Exit(1)
		}
		fmt.Println("De Bruijn sequence:", dumpPath(path))
		//fmt.Println("De Bruijn sequence(HEX):", d.seq2hex(seq))
		fmt.Println("De Bruijn sequence length:", len(path))
	}
}
