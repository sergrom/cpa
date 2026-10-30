package ds

import "math/bits"

type SpTab struct {
	lookup [][]int
}

func NewSparseTable(arr []int) *SpTab {
	n := len(arr)
	levels := bits.Len(uint(n))
	lookup := make([][]int, n)
	for i := range lookup {
		lookup[i] = make([]int, levels)
		lookup[i][0] = arr[i]
	}

	for j := 1; (1 << j) <= n; j++ {
		for i := 0; (i + (1 << j) - 1) < n; i++ {
			if lookup[i][j-1] < lookup[i+(1<<(j-1))][j-1] {
				lookup[i][j] = lookup[i][j-1]
			} else {
				lookup[i][j] = lookup[i+(1<<(j-1))][j-1]
			}
		}
	}

	return &SpTab{
		lookup: lookup,
	}
}

func (st *SpTab) Query(l, r int) int {
	j := bits.Len(uint(r-l+1)) - 1
	if st.lookup[l][j] <= st.lookup[r-(1<<j)+1][j] {
		return st.lookup[l][j]
	}
	return st.lookup[r-(1<<j)+1][j]
}
