package alg

import "math/bits"

// modPow returns x^e modulo m in [0, m). Requires e >= 0 and m > 0.
// Intermediate products do not overflow, even for a modulus near MaxInt.
func modPow(x, e, m int) int {
	return int(modPow64(int64(x), int64(e), int64(m)))
}

// modPow64 returns x^e modulo m in O(log e) time.
// It panics if e < 0 or m <= 0. In particular, x^0 mod 1 is zero.
func modPow64(x, e, m int64) int64 {
	if e < 0 || m <= 0 {
		panic("modPow64: requires e >= 0 and m > 0")
	}
	x %= m
	if x < 0 {
		x += m
	}
	base, modulus := uint64(x), uint64(m)
	result := uint64(1) % modulus
	for e > 0 {
		if e&1 != 0 {
			result = mulMod64(result, base, modulus)
		}
		base = mulMod64(base, base, modulus)
		e >>= 1
	}
	return int64(result)
}

// a and b must be less than m. Then the high word of a*b is also less
// than m, so Div64's quotient fits in uint64.
func mulMod64(a, b, m uint64) uint64 {
	hi, lo := bits.Mul64(a, b)
	_, remainder := bits.Div64(hi, lo, m)
	return remainder
}
