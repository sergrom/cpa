package graph

import "math"

// Bellman-Ford Algorithm - O(V*E) Time and O(V) Space

// BellmanFord Bellman-Ford is a single source shortest path algorithm.
// It effectively works in the cases of negative edges and is able to detect
// negative cycles as well. It works on the principle of relaxation of the edges.
func BellmanFord(nodesCnt int, edges [][]int, src int) []int {
	// Initially distance from source to all
	// other vertices is not known(Infinite).
	dist := make([]int, nodesCnt)
	for i := range dist {
		dist[i] = math.MaxInt
	}
	dist[src] = 0

	// Relaxation of all the edges V times, not (V - 1) as we
	// need one additional relaxation to detect negative cycle
	for i := 0; i < nodesCnt; i++ {
		for _, edg := range edges {
			u, v, wt := edg[0], edg[1], edg[2]
			if dist[u] != math.MaxInt && dist[u]+wt < dist[v] {
				// If this is the Vth relaxation, then there is
				// a negative cycle
				if i == nodesCnt-1 {
					return nil
				}

				// Update shortest distance to node v
				dist[v] = dist[u] + wt
			}
		}
	}

	return dist
}

// func main() {
// 	V := 5
// 	edges := [][]int{{1, 3, 2}, {4, 3, -1}, {2, 4, 1}, {1, 2, 1}, {0, 1, 5}}
// 	src := 0
// 	ans := BellmanFord(V, edges, src)
// 	for n, dist := range ans {
// 		fmt.Println(n, dist)
// 	}
// }
