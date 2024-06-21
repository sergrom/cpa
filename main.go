package main

import "fmt"

func countOfPeaks(nums []int, queries [][]int) []int {

	arr := make([]int, len(nums))
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			arr[i] = 1
		}
	}

	st := NewSegTree(arr)
	ans := make([]int, 0)
	for _, q := range queries {
		if q[0] == 1 {
			v := st.Query(q[1], q[2]+1)
			if arr[q[1]] == 1 {
				v--
			}
			if arr[q[2]] == 1 && q[1] != q[2] {
				v--
			}
			ans = append(ans, v)
		} else {
			idx, val := q[1], q[2]
			nums[idx] = val

			start := idx - 1
			if start < 0 {
				start = 0
			}
			for i := start; i <= idx+1 && i < len(nums)-1; i++ {
				isPeak := i > 0 && i < len(nums)-1 && nums[i] > nums[i-1] && nums[i] > nums[i+1]
				if !isPeak && arr[i] == 1 {
					st.Update(i, 0)
					arr[i] = 0
				}
				if isPeak && arr[i] == 0 {
					st.Update(i, 1)
					arr[i] = 1
				}
			}
		}
	}

	// fmt.Println(arr)

	return ans
}

func main() {
	// 4, 9, 4, 10, 7 -> 4, 3, 4, 3, 7

	// arr := []int{0, 0, 1, 0, 0} //
	// st := NewSegTree(arr)

	// st.Update(3, 2)
	// st.Update(1, 3)

	// fmt.Println(st.Query(2, 3))

	fmt.Println(countOfPeaks([]int{3, 1, 4, 2, 5}, [][]int{{2, 3, 4}, {1, 0, 4}}))               // [0]
	fmt.Println(countOfPeaks([]int{4, 1, 4, 2, 1, 5}, [][]int{{2, 2, 4}, {1, 0, 2}, {1, 0, 4}})) // [0,1]
	fmt.Println(countOfPeaks([]int{3, 6, 9}, [][]int{{1, 1, 1}, {1, 2, 2}, {2, 2, 3}}))          // [0,1]
	fmt.Println(countOfPeaks([]int{7, 10, 7}, [][]int{{1, 2, 2}, {2, 0, 6}, {1, 0, 2}}))         // [0,1]
	fmt.Println(countOfPeaks([]int{4, 9, 4, 10, 7}, [][]int{{2, 3, 2}, {2, 1, 3}, {1, 2, 3}}))   // [0]
	// 4, 3, 4, 2, 7
}

// SegTree ...
type SegTree struct {
	n    int
	tree []int
}

func NewSegTree(arr []int) *SegTree {
	n := len(arr)
	tree := make([]int, 2*n)

	for i := 0; i < n; i++ {
		tree[n+i] = arr[i]
	}

	for i := n - 1; i > 0; i-- {
		tree[i] = tree[i<<1] + tree[i<<1|1]
	}

	return &SegTree{
		n:    n,
		tree: tree,
	}
}

func (st *SegTree) Update(idx int, val int) {
	st.tree[idx+st.n] = val
	idx += st.n

	for i := idx; i > 1; i >>= 1 {
		st.tree[i>>1] = st.tree[i] + st.tree[i^1]
	}
}

func (st *SegTree) Query(l, r int) int {
	res := 0

	l, r = l+st.n, r+st.n
	for l < r {
		if l&1 > 0 {
			res += st.tree[l]
			l++
		}
		if r&1 > 0 {
			r--
			res += st.tree[r]
		}
		l, r = l>>1, r>>1
	}

	return res
}
