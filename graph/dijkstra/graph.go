package dijkstra

import "container/heap"

type edge struct {
	node   string
	weight int
}

type Graph struct {
	nodes map[string][]edge
}

func NewGraph() *Graph {
	return &Graph{nodes: make(map[string][]edge)}
}

func (g *Graph) AddEdge(origin, destiny string, weight int) {
	g.nodes[origin] = append(g.nodes[origin], edge{node: destiny, weight: weight})
	g.nodes[destiny] = append(g.nodes[destiny], edge{node: origin, weight: weight})
}

func (g *Graph) GetPath(origin, destiny string) (int, []string) {
	h := newHeap()
	h.push(path{value: 0, nodes: []string{origin}})
	visited := make(map[string]bool)

	for len(*h.values) > 0 {
		// Find the nearest yet to visit node
		p := h.pop()
		node := p.nodes[len(p.nodes)-1]

		if visited[node] {
			continue
		}

		if node == destiny {
			return p.value, p.nodes
		}

		for _, e := range g.getEdges(node) {
			if !visited[e.node] {
				// We calculate the total spent so far plus the cost and the path of getting here
				h.push(path{value: p.value + e.weight, nodes: append([]string{}, append(p.nodes, e.node)...)})
			}
		}

		visited[node] = true
	}

	return 0, nil
}

func (g *Graph) getEdges(node string) []edge {
	return g.nodes[node]
}

// ----

type path struct {
	value int
	nodes []string
}

type minPath []path

func (h minPath) Len() int           { return len(h) }
func (h minPath) Less(i, j int) bool { return h[i].value < h[j].value }
func (h minPath) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minPath) Push(x interface{}) {
	*h = append(*h, x.(path))
}

func (h *minPath) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type hp struct {
	values *minPath
}

func newHeap() *hp {
	return &hp{values: &minPath{}}
}

func (h *hp) push(p path) {
	heap.Push(h.values, p)
}

func (h *hp) pop() path {
	i := heap.Pop(h.values)
	return i.(path)
}
