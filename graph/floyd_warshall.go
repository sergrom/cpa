package graph

import "math"

// floydWarshall computes shortest paths in a directed graph in O(n^3) time.
// Edges are {from, to, weight}; vertices are numbered 0..n-1.
// math.MaxInt denotes unreachable vertices. Parallel edges use the minimum weight.
// A negative cycle returns nil. Finite distances and intermediate sums must fit
// in int and be less than math.MaxInt; an unrepresentable sum causes a panic.
func floydWarshall(n int, edges [][]int) [][]int {
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt
		}
		dist[i][i] = 0
	}
	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		if weight == math.MaxInt {
			panic("floydWarshall: edge weight conflicts with infinity")
		}
		dist[from][to] = min(dist[from][to], weight)
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			if dist[i][k] == math.MaxInt {
				continue
			}
			for j := 0; j < n; j++ {
				if dist[k][j] == math.MaxInt {
					continue
				}
				a, b := dist[i][k], dist[k][j]
				if (b > 0 && a >= math.MaxInt-b) || (b < 0 && a < math.MinInt-b) {
					panic("floydWarshall: distance overflows finite int range")
				}
				dist[i][j] = min(dist[i][j], a+b)
				if i == j && dist[i][j] < 0 {
					return nil
				}
			}
		}
	}
	return dist
}
