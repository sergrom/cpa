package alg

func subsets(nums []int) [][]int {
	subsets := [][]int{{}}
	for _, num := range nums {
		for _, subset := range subsets {
			subsets = append(subsets, append([]int{}, append(subset, num)...))
		}
	}
	return subsets
}

// subsets([]int{1,2,3})
// outputs: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
