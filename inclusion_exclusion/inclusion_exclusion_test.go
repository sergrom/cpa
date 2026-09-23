package inclusionexclusion

import "testing"

func TestCoprimeCount(t *testing.T) {
	for n := 1; n <= 120; n++ {
		for r := 0; r <= 120; r++ {
			want := 0
			for x := 1; x <= r; x++ {
				a, b := n, x
				for b != 0 {
					a, b = b, a%b
				}
				if a == 1 {
					want++
				}
			}
			if got := solve(n, r); got != want {
				t.Fatalf("solve(%d,%d)=%d want %d", n, r, got, want)
			}
		}
	}
}
