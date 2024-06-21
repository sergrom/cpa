package alg

// You are given an integer array coins representing coins of different denominations and an integer k.
// You have an infinite number of coins of each denomination.
// However, you are not allowed to combine coins of different denominations.
// Return the kth smallest amount that can be made using these coins.

// Input: coins = [3,6,9], k = 3

// Output: 9

// Explanation: The given coins can make the following amounts:
// Coin 3 produces multiples of 3: 3, 6, 9, 12, 15, etc.
// Coin 6 produces multiples of 6: 6, 12, 18, 24, etc.
// Coin 9 produces multiples of 9: 9, 18, 27, 36, etc.
// All of the coins combined produce: 3, 6, 9, 12, 15, etc.

// getK returns how many different amounts less or equal than val can be made with coins
func getK(coins []int, val int) int {
	var ans int

	for msk := 1; msk < (1 << len(coins)); msk++ {
		lc, bits := 1, 0
		for i := 0; i < len(coins); i++ {
			if msk&(1<<i) > 0 {
				bits++
				lc = lcm(lc, coins[i])
			}
		}

		cur := val / lc
		if bits%2 == 1 {
			ans += cur
		} else {
			ans -= cur
		}
	}

	return ans
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
