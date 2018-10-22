package ch10

type AdjacencyMatrix struct {
	n int
	a [][]bool
}

type List []int

func NewAdjacencyMatrix(n int) *AdjacencyMatrix {
	a := make([][]bool, n)
	for i := 0; i < n; i++ {
		a[i] = make([]bool, n)
	}
	return &AdjacencyMatrix{a: a, n: n}
}

func (m *AdjacencyMatrix) addEdge(i int, j int) {
	m.a[i][j] = true
}

func (m *AdjacencyMatrix) removeEdge(i int, j int) {
	m.a[i][j] = false
}

func (m *AdjacencyMatrix) hasEdge(i int, j int) bool {
	return m.a[i][j]
}

func (m *AdjacencyMatrix) outEdges(i int) List {
	edges := make(List, 0)
	for j := 0; j < m.n; j++ {
		if m.a[i][j] {
			edges = append(edges, j)
		}
	}
	return edges
}

func (m *AdjacencyMatrix) inEdges(i int) List {
	edges := make(List, 0)
	for j := 0; j < m.n; j++ {
		if m.a[j][i] {
			edges = append(edges, j)
		}
	}
	return edges
}
