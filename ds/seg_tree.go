package ds

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
