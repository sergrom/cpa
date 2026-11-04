package ds

import (
	"math"
	"math/rand"
	"testing"
)

// Selecting the first nonzero value is associative with identity zero,
// but changing the argument order can change the result.
type firstNonZeroAggregator struct{}

func (firstNonZeroAggregator) GetDefaultValue() int { return 0 }
func (firstNonZeroAggregator) Aggregate(a, b int) int {
	if a != 0 {
		return a
	}
	return b
}

func TestSegTreeMaxMinInt(t *testing.T) {
	st := NewSegTree([]int{math.MinInt}, &MaxAggregator{})
	if got := st.Query(0, 1); got != math.MinInt {
		t.Fatalf("got %d, want %d", got, math.MinInt)
	}
}

func TestSegTreeAgainstSlice(t *testing.T) {
	for name, agg := range map[string]Aggregator{
		"min": &MinAggregator{}, "max": &MaxAggregator{},
		"sum": &SumAggregator{}, "firstNonZero": firstNonZeroAggregator{},
	} {
		t.Run(name, func(t *testing.T) {
			rng := rand.New(rand.NewSource(42))
			for n := 0; n <= 20; n++ {
				arr := make([]int, n)
				for i := range arr {
					arr[i] = rng.Intn(11) - 5
				}
				st := NewSegTree(arr, agg)
				for step := 0; step < 20; step++ {
					for l := 0; l <= n; l++ {
						want := agg.GetDefaultValue()
						for r := l; r <= n; r++ {
							if r > l {
								want = agg.Aggregate(want, arr[r-1])
							}
							if got := st.Query(l, r); got != want {
								t.Fatalf("n=%d step=%d Query(%d,%d)=%d, want %d; arr=%v", n, step, l, r, got, want, arr)
							}
						}
					}
					if n > 0 {
						idx, value := rng.Intn(n), rng.Intn(11)-5
						arr[idx] = value
						st.Update(idx, value)
					}
				}
			}
		})
	}
}
