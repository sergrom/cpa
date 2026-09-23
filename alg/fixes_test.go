package alg

import (
	"math"
	"math/big"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func mustPanic(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if recover() == nil {
			t.Error("expected panic")
		}
	}()
	fn()
}

func TestBinomialExact(t *testing.T) {
	for n := 0; n <= 70; n++ {
		for k := 0; k <= n; k++ {
			want := new(big.Int).Binomial(int64(n), int64(k))
			if !want.IsInt64() || want.Int64() > int64(math.MaxInt) {
				mustPanic(t, func() { C(n, k) })
			} else if got := C(n, k); int64(got) != want.Int64() {
				t.Fatalf("C(%d,%d)=%d want %v", n, k, got, want)
			}
		}
	}
	for _, pair := range [][2]int{{-1, 0}, {2, -1}, {2, 3}} {
		if C(pair[0], pair[1]) != 0 {
			t.Fatal(pair)
		}
	}
	if C(math.MaxInt, 1) != math.MaxInt {
		t.Fatal("large n, k=1")
	}
}

func TestModPow(t *testing.T) {
	cases := [][3]int64{
		{10, 0, 1}, {0, 0, 7}, {4000000000, 2, 5000000000},
		{math.MinInt64, 3, math.MaxInt64}, {math.MaxInt64 - 1, math.MaxInt64, math.MaxInt64},
		{-2, 3, 5},
	}
	rng := rand.New(rand.NewSource(7))
	for i := 0; i < 1000; i++ {
		cases = append(cases, [3]int64{int64(rng.Uint64()), rng.Int63(), rng.Int63n(math.MaxInt64) + 1})
	}
	for _, tc := range cases {
		x, e, m := tc[0], tc[1], tc[2]
		want := new(big.Int).Exp(big.NewInt(x), big.NewInt(e), big.NewInt(m)).Int64()
		if got := modPow64(x, e, m); got != want {
			t.Fatalf("modPow64(%d,%d,%d)=%d want %d", x, e, m, got, want)
		}
		if x >= int64(math.MinInt) && x <= int64(math.MaxInt) && e <= int64(math.MaxInt) && m <= int64(math.MaxInt) {
			if got := modPow(int(x), int(e), int(m)); int64(got) != want {
				t.Fatal("modPow", tc, got, want)
			}
		}
	}
	for _, tc := range [][3]int64{{1, -1, 2}, {1, 2, 0}, {1, 2, -3}} {
		mustPanic(t, func() { modPow64(tc[0], tc[1], tc[2]) })
	}
}

func TestSuffixAndCyclicOrder(t *testing.T) {
	cases := []string{"", "a", "baa", "abab", "banana", "aaaa", "\x00\xff\x00$"}
	rng := rand.New(rand.NewSource(8))
	for i := 0; i < 500; i++ {
		b := make([]byte, rng.Intn(80))
		rng.Read(b)
		cases = append(cases, string(b))
	}
	for n := 1; n <= 7; n++ {
		for mask := 0; mask < 1<<n; mask++ {
			b := make([]byte, n)
			for i := range b {
				b[i] = byte((mask >> i) & 1)
			}
			cases = append(cases, string(b))
		}
	}
	for _, s := range cases {
		want := make([]int, len(s))
		for i := range want {
			want[i] = i
		}
		sort.Slice(want, func(i, j int) bool { return s[want[i]:] < s[want[j]:] })
		if got := suffixArr(s); !reflect.DeepEqual(got, want) {
			t.Fatalf("suffixArr(%q)=%v want %v", s, got, want)
		}
		p := sortCyclicShifts(s)
		seen := make(map[int]bool)
		if len(p) != len(s) {
			t.Fatal("rotation count")
		}
		for i, pos := range p {
			if pos < 0 || pos >= len(s) || seen[pos] {
				t.Fatal("invalid permutation", p)
			}
			seen[pos] = true
			if i > 0 {
				prev := p[i-1]
				if s[prev:]+s[:prev] > s[pos:]+s[:pos] {
					t.Fatal("unsorted rotations", s, p)
				}
			}
		}
	}
}
