package main

func (d *D) makeGraph() map[string][]string {
	graph := make(map[string][]string)
	for _, v := range d.V {
		word := string(v)
		graph[word] = make([]string, d.n)
		for _, n := range d.V {
			next := string(n)
			//if next==word {
			//continue
			//}
			if word[1:] == next[0:d.n-1] {
				graph[word] = append(graph[word], string(next))
				//fmt.Println("graph[", word, "]=", string(next))
			}
		}
	}
	return graph
}

func (d *D) _Seq(checklist map[string]bool, path []string) []string {
	return nil
}

func (d *D) Seq() []string {
	checklist := make(map[string]bool)
	initWord := string(d.V[0])
	checklist[initWord] = true
	return d._Seq(checklist, []string{initWord})
}
