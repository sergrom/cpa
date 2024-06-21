package graph

func hasCicles(edges [][]int) bool {
	vertexes, vtxMap := make([]int, 0), make(map[int]bool)
	edgMap := make(map[int]map[int]bool)
	for _, edg := range edges {
		from, to := edg[0], edg[1]
		if edgMap[from] == nil {
			edgMap[from] = make(map[int]bool)
			vertexes = append(vertexes, from)
			vtxMap[from] = true
		}
		edgMap[from][to] = true
		if !vtxMap[to] {
			vertexes = append(vertexes, to)
			vtxMap[to] = true
		}
	}

	visited := make(map[int]bool)
	recStack := make(map[int]bool)

	for _, vtx := range vertexes {
		if !visited[vtx] && isCyclicVtx(vtx, edgMap, visited, recStack) {
			return true
		}
	}

	return false
}

func isCyclicVtx(vtx int, edgMap map[int]map[int]bool, visited, recStack map[int]bool) bool {
	if !visited[vtx] {
		visited[vtx], recStack[vtx] = true, true

		for vtxNext := range edgMap[vtx] {
			if !visited[vtxNext] && isCyclicVtx(vtxNext, edgMap, visited, recStack) {
				return true
			} else if recStack[vtxNext] {
				return true
			}
		}
	}

	recStack[vtx] = false
	return false
}
