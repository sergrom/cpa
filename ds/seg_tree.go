package ds

import "math"

// SegTree ...
type SegTree struct {
	n    int
	tree []int
	agg  Aggregator
}

func NewSegTree(arr []int, agg Aggregator) *SegTree {
	n := len(arr)
	tree := make([]int, 2*n)

	for i := 0; i < n; i++ {
		tree[n+i] = arr[i]
	}

	for i := n - 1; i > 0; i-- {
		tree[i] = agg.Aggregate(tree[i<<1], tree[i<<1|1])
	}

	return &SegTree{
		n:    n,
		tree: tree,
		agg:  agg,
	}
}

func (st *SegTree) Update(idx int, val int) {
	st.tree[idx+st.n] = val
	idx += st.n

	for i := idx; i > 1; i >>= 1 {
		parent := i >> 1
		st.tree[parent] = st.agg.Aggregate(st.tree[parent<<1], st.tree[parent<<1|1])
	}
}

// Query returns the aggregate over [l, r). An empty range returns the identity.
func (st *SegTree) Query(l, r int) int {
	left, right := st.agg.GetDefaultValue(), st.agg.GetDefaultValue()

	l, r = l+st.n, r+st.n
	for l < r {
		if l&1 > 0 {
			left = st.agg.Aggregate(left, st.tree[l])
			l++
		}
		if r&1 > 0 {
			r--
			right = st.agg.Aggregate(st.tree[r], right)
		}
		l, r = l>>1, r>>1
	}

	return st.agg.Aggregate(left, right)
}

// ---------

// Aggregator defines an associative operation and its two-sided identity.
// Aggregate need not be commutative; values are combined from left to right.
type Aggregator interface {
	GetDefaultValue() int
	Aggregate(v1, v2 int) int
}

type MinAggregator struct{}

func (s *MinAggregator) GetDefaultValue() int     { return math.MaxInt }
func (s *MinAggregator) Aggregate(v1, v2 int) int { return min(v1, v2) }

// ---------

type MaxAggregator struct{}

func (s *MaxAggregator) GetDefaultValue() int     { return math.MinInt }
func (s *MaxAggregator) Aggregate(v1, v2 int) int { return max(v1, v2) }

// ---------

type SumAggregator struct{}

func (s *SumAggregator) GetDefaultValue() int     { return 0 }
func (s *SumAggregator) Aggregate(v1, v2 int) int { return v1 + v2 }
