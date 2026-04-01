package "alg"

// compress - coordinate compression
func compress(nums []int) map[int]int {
	vals := append([]int{}, nums...)
	sort.Ints(vals)

	uniq := []int{}
	for i, v := range vals {
		if i == 0 || v != vals[i-1] {
			uniq = append(uniq, v)
		}
	}

	comp := make(map[int]int)
	for i, v := range uniq {
		comp[v] = i
	}

	return comp
}
