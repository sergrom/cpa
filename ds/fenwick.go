package ds

// Fenwick implements Binary Indexed Tree for sums.
type Fenwick struct {
	tree []int
	n    int
}

// NewFenwickFrom
func NewFenwick(arr []int) *Fenwick {
	n := len(arr)
	f := &Fenwick{
		tree: make([]int, n+1),
		n:    n,
	}

	// копируем массив (1-based)
	for i := 1; i <= n; i++ {
		f.tree[i] = arr[i-1]
	}

	// строим дерево
	for i := 1; i <= n; i++ {
		j := i + (i & -i)
		if j <= n {
			f.tree[j] += f.tree[i]
		}
	}

	return f
}

// Add adds delta to element at index i (0-based index).
func (f *Fenwick) Add(i int, delta int) {
	i++ // switch to 1-based index
	for i <= f.n {
		f.tree[i] += delta
		i += i & -i
	}
}

// PrefixSum returns sum of elements in range [0..i] (0-based index).
func (f *Fenwick) PrefixSum(i int) int {
	if i < 0 {
		return 0
	}

	i++ // switch to 1-based index
	sum := 0
	for i > 0 {
		sum += f.tree[i]
		i -= i & -i
	}
	return sum
}

// RangeSum returns sum of elements in range [l..r] (0-based index).
func (f *Fenwick) RangeSum(l, r int) int {
	if l > r {
		return 0
	}
	return f.PrefixSum(r) - f.PrefixSum(l-1)
}

// Set sets nums[i] = value.
// Since Fenwick stores aggregated data, we first find current value.
func (f *Fenwick) Set(i int, value int) {
	current := f.RangeSum(i, i)
	f.Add(i, value-current)
}
