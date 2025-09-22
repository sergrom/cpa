package ds

type FenwickTree struct {
	tree []int
}

func fenwickTree(size int) *FenwickTree {
	return &FenwickTree{tree: make([]int, size+1)}
}

func (ft *FenwickTree) update(index, delta int) {
	index++
	for index < len(ft.tree) {
		ft.tree[index] += delta
		index += index & -index
	}
}

func (ft *FenwickTree) query(index int) int {
	index++
	res := 0
	for index > 0 {
		res += ft.tree[index]
		index -= index & -index
	}
	return res
}
