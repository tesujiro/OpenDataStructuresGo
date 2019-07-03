package main

import "fmt"

func (d *D) makeGraph() map[string][]string {
	graph := make(map[string][]string)
	//for i, v := range d.V {
	//fmt.Printf("V[%v]=%v\n", i, string(v))
	//}
	for _, v := range d.V {
		word := string(v)
		graph[word] = []string{}
		//fmt.Printf("graph[%v]=%v\n", word, graph[word])
		//fmt.Println("len(d.V)=",
		for _, n := range d.V {
			next := string(n)
			//if next==word {
			//continue
			//}
			if next == "" {
				fmt.Println("NULL")
			}
			if next != "" && word[1:] == next[0:d.n-1] {
				graph[word] = append(graph[word], string(next))
				//fmt.Println(" graph[", word, "]=", string(next))
			}
		}
		//fmt.Printf("->graph[%v]=%v len=%v\n", word, graph[word], len(graph[word]))
	}
	return graph
}

func (d *D) _seqGraph(checklist map[string]bool, path []string) ([]string, bool) {
	if len(path) == len(d.V)-d.n+2 {
		fmt.Println("path:", path)
		fmt.Println("path[last]:", path[len(path)-1:])
		return path, true
	}
	last := path[len(path)-1]
	//fmt.Println("last=", last, " len(d.G[last])=", len(d.G[last]), " G[last]=", d.G[last])
	for _, word := range d.G[last] {
		//if word == "" {
		//fmt.Println("word=", word, " last=", last, " i=", i)
		//}
		if checklist[word] == true {
			continue
		}
		checklist[word] = true
		//fmt.Println("word=", word, " len(d.G[word])=", len(d.G[word]), " G[word]=", d.G[word])
		if ret, ok := d._seqGraph(checklist, append(path, string(word))); ok {
			return ret, true
		} else {
			delete(checklist, word)
		}
	}
	fmt.Printf("FALSE path=%v\n", path)
	return nil, false
}

func (d *D) SeqGraph() ([]string, bool) {
	checklist := make(map[string]bool)
	initWord := string(d.V[0])
	checklist[initWord] = true
	return d._seqGraph(checklist, []string{initWord})
}
