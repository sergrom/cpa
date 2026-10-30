package ds

import (
	"math/bits"
	"math/rand"
	"testing"
)

func TestSparseTable(t *testing.T) {
	rng := rand.New(rand.NewSource(7))
	for n := 0; n <= 100; n++ {
		arr := make([]int, n)
		for i := range arr {
			arr[i] = rng.Intn(200) - 100
		}
		st := NewSparseTable(arr)
		for _, row := range st.lookup {
			if len(row) != bits.Len(uint(n)) {
				t.Fatal("incorrect table size")
			}
		}
		for l := range arr {
			want := arr[l]
			for r := l; r < n; r++ {
				want = min(want, arr[r])
				if got := st.Query(l, r); got != want {
					t.Fatalf("Query(%d,%d)=%d, want %d", l, r, got, want)
				}
			}
		}
	}
}
