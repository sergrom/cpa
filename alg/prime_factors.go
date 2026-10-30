package alg

import "math"

// primeFactors returns the distinct prime factors of num.
// For num <= 1 it returns an empty slice.
func primeFactors(num int) []int {
	ans := make([]int, 0)
	if num <= 1 {
		return ans
	}
	if num%2 == 0 {
		ans = append(ans, 2)
		for num%2 == 0 {
			num /= 2
		}
	}
	sqrt := int(math.Sqrt(float64(num)))
	for i := 3; i <= sqrt; i += 2 {
		if num%i == 0 {
			ans = append(ans, i)
			for num%i == 0 {
				num = num / i
			}
		}
	}
	if num > 2 {
		ans = append(ans, num)
	}

	return ans
}
