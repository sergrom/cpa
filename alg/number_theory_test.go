package alg

import (
	"reflect"
	"testing"
)

func TestIsPrime(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want bool
	}{
		{-7, false}, {-1, false}, {0, false}, {1, false}, {2, true}, {3, true}, {4, false}, {49, false}, {97, true},
	} {
		if got := isPrime(tc.n); got != tc.want {
			t.Errorf("isPrime(%d)=%v, want %v", tc.n, got, tc.want)
		}
	}
}

func TestPrimeFactors(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want []int
	}{
		{-7, []int{}}, {0, []int{}}, {1, []int{}}, {2, []int{2}}, {8, []int{2}}, {45, []int{3, 5}}, {60, []int{2, 3, 5}}, {97, []int{97}},
	} {
		if got := primeFactors(tc.n); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("primeFactors(%d)=%v, want %v", tc.n, got, tc.want)
		}
	}
}
