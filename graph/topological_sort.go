package graph

// Kahn’s algorithm for Topological Sorting
// https://www.geeksforgeeks.org/topological-sorting-indegree-based-solution/

func topologicalSort(adj map[int]map[int]bool, n int) []int {
	// Vector to store indegree of each vertex
	indegree := make(map[int]int)
	for i := 0; i < n; i++ {
		for j := range adj[i] {
			indegree[j]++
		}
	}

	// Queue to store vertices with indegree 0
	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	result := make([]int, 0, n)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)

		// Decrease indegree of adjacent vertices as the
		// current node is in topological order
		for j := range adj[node] {
			indegree[j]--
			if indegree[j] == 0 {
				queue = append(queue, j)
			}
		}
	}

	if len(result) != n {
		return []int{}
	}

	return result
}
