package alg

import (
	"math"
	"math/big"
)

// Binomial coefficients are the number of ways
// to select a set of 'k' elements from 'n' different elements
// without taking into account the order
// of arrangement of these elements (i.e., the number of unordered sets).

// C returns the exact binomial coefficient. Invalid n or k returns zero.
// It panics if the result cannot be represented as an int.
func C(n, k int) int {
	if n < 0 || k < 0 || k > n {
		return 0
	}
	value := new(big.Int).Binomial(int64(n), int64(k))
	if !value.IsInt64() || value.Int64() > int64(math.MaxInt) {
		panic("C: result overflows int")
	}
	return int(value.Int64())
}

func binomialCoeff(n, r int) int {
	if r > n {
		return 0
	}

	m := 1_000_000_007
	inv := make([]int, r+1)
	inv[0] = 1
	if r+1 >= 2 {
		inv[1] = 1
	}

	// Getting the modular inversion
	// for all the numbers
	// from 2 to r with respect to m
	// here m = 1000000007
	for i := 2; i <= r; i++ {
		inv[i] = m - (m/i)*inv[m%i]%m
	}

	ans := 1

	// for 1/(r!) part
	for i := 2; i <= r; i++ {
		ans = ((ans % m) * (inv[i] % m)) % m
	}

	// for (n)*(n-1)*(n-2)*...*(n-r+1) part
	for i := n; i >= (n - r + 1); i-- {
		ans = ((ans % m) * (i % m)) % m
	}
	return ans
}
