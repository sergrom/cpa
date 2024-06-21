package graph

func adjMap(edges [][]int) map[int]map[int]bool {
	edgMap := make(map[int]map[int]bool)
	for _, pr := range edges {
		from, to := pr[0], pr[1]
		if edgMap[from] == nil {
			edgMap[from] = make(map[int]bool)
		}
		edgMap[from][to] = true
	}

	return edgMap
}
