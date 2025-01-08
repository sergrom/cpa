package graph

func floydWarshall(n int, edges [][]int) [][]int {
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
	}

	for _, edg := range edges {
		i, j, v := edg[0], edg[1], edg[2]
		dist[i][j] = v
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}

	return dist
}
