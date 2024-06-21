package strpatternsearch

func findLongestPalindromicString(text string) string {
	N := len(text)
	if N == 0 {
		return ""
	}

	N = 2*N + 1         // Position count
	L := make([]int, N) // LPS Length Array
	L[0], L[1] = 0, 1
	C, R := 1, 2    // centerPosition, centerRightPosition
	var iMirror int // currentLeftPosition
	maxLPSLength, maxLPSCenterPosition := 0, 0
	start, end, diff := -1, -1, -1

	// Uncomment it to print LPS Length array
	for i := 2; i < N; i++ {
		// get currentLeftPosition iMirror for currentRightPosition i
		iMirror = 2*C - i
		L[i] = 0
		diff = R - i
		// If currentRightPosition i is within centerRightPosition R
		if diff > 0 {
			L[i] = min(L[iMirror], diff)
		}

		// Attempt to expand palindrome centered at currentRightPosition i
		// Here for odd positions, we compare characters and
		// if match then increment LPS Length by ONE
		// If even position, we just increment LPS by ONE without
		// any character comparison
		for ((i+L[i])+1 < N && (i-L[i]) > 0) &&
			(((i+L[i]+1)%2 == 0) ||
				(text[(i+L[i]+1)/2] == text[(i-L[i]-1)/2])) {
			L[i]++
		}

		if L[i] > maxLPSLength { // Track maxLPSLength
			maxLPSLength = L[i]
			maxLPSCenterPosition = i
		}

		// If palindrome centered at currentRightPosition i
		// expand beyond centerRightPosition R,
		// adjust centerPosition C based on expanded palindrome.
		if i+L[i] > R {
			C, R = i, i+L[i]
		}
	}

	start = (maxLPSCenterPosition - maxLPSLength) / 2
	end = start + maxLPSLength

	return text[start:end]
}
