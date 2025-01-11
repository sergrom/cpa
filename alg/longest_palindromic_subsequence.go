package alg

// LPS - longest palindromic subsequence
// Bottom-up DP solution
// Time Complexity: O(n2)
// Auxiliary Space: O(n2)
// The final answer is stored in dp[0][n-1]

func lpsArr(str string) [][]int {
	n := len(str)

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if str[i] == str[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp
}
