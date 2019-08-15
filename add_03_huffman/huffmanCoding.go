package main

import "fmt"

func main() {

	sub(`ハフマン符号（ハフマンふごう、英: Huffman coding）とは、1952年にデビッド・ハフマンによって開発された符号で、文字列をはじめとするデータの可逆圧縮などに使用される。`)
	fmt.Println("================")
	sub(`What is the most frequent letter in English?`)
	fmt.Println("================")
}

func sub(text string) {
	f := newFreq(text)
	fmt.Printf("text:  %v\n", text)

	tree := NewHuffmanTree(f)
	PrintTree(tree)
}

func newFreq(text string) map[rune]int {
	f := make(map[rune]int)
	for _, r := range text {
		f[r]++
	}
	return f
}

type Leaf struct {
	freq  int
	value rune
}

type Node struct {
	freq  int
	left  HuffmanTree
	right HuffmanTree
}

func (l Leaf) Freq() int {
	return l.freq
}

func (n Node) Freq() int {
	return n.freq
}

type HuffmanTree interface {
	Freq() int
}

func NewHuffmanTree(f map[rune]int) HuffmanTree {
	// init
	var trees []HuffmanTree
	for k, v := range f {
		trees = append(trees, Leaf{value: k, freq: v})
	}

	for i := len(f); i > 1; i-- {
		// find minimal 1,2 frequency
		min1 := 0 // index of trees with minimal 1 freq
		min2 := 1 // index of trees with minimal 2 freq
		if trees[min1].Freq() > trees[min2].Freq() {
			min1, min2 = min2, min1
		}
		for j, tree := range trees[2:i] {
			if tree.Freq() < trees[min1].Freq() {
				min2 = min1
				min1 = j + 2
			} else if tree.Freq() < trees[min2].Freq() {
				min2 = j + 2
			}
		}
		// create one node with min1, min2
		node := Node{freq: trees[min1].Freq() + trees[min2].Freq(), left: trees[min1], right: trees[min2]}
		trees[min1] = node
		trees[min2] = trees[i-1]
	}
	return trees[0]
}

func PrintTree(h HuffmanTree) {
	_print(h, "")
}

func _print(h HuffmanTree, code string) {
	switch h := h.(type) {
	case Leaf:
		//fmt.Println("leaf")
		fmt.Printf("%c (%v)\t: 0x%v\n", h.value, h.freq, code)
	case Node:
		//fmt.Println("node:", code)
		_print(h.left, code+"0")
		_print(h.right, code+"1")
	default:
		//fmt.Printf("not leaf nor node:%#v:%v\n", h, code)
	}
}
