package strpatternsearch

// Finite Automata algorithm for Pattern Searching
// Time Complexity: O(m^2) + O(n),
// Auxiliary Space: O(m),
// where m = len(pattern), n = len(text)

// Prints all occurrences of pat in txt
func FSSearch(text, pattern string) []int {
	M := len(pattern)
	N := len(text)

	TF := computeTF(pattern, M)

	ans := make([]int, 0)

	// Process txt over FA.
	state := 0
	for i := 0; i < N; i++ {
		state = TF[state][text[i]]
		if state == M {
			ans = append(ans, i-M+1)
		}
	}

	return ans
}

func getNextState(pattern string, M int, state int, x byte) int {
	// If the character c is same as next character
	// in pattern,then simply increment state
	if state < M && x == pattern[state] {
		return state + 1
	}

	// ns stores the result which is next state

	// ns finally contains the longest prefix
	// which is also suffix in "pat[0..state-1]c"

	// Start from the largest possible value
	// and stop when you find a prefix which
	// is also suffix
	for ns := state; ns > 0; ns-- {
		if pattern[ns-1] == x {
			var i int
			for i = 0; i < ns-1; i++ {
				if pattern[i] != pattern[state-ns+1+i] {
					break
				}
			}
			if i == ns-1 {
				return ns
			}
		}
	}

	return 0
}

// computeTF - This function builds the TF table which represents Finite Automata for a given pattern
func computeTF(pat string, M int) [][256]int {
	TF := make([][256]int, M+1)
	for state := 0; state <= M; state++ {
		for x := 0; x < 256; x++ {
			TF[state][x] = getNextState(pat, M, state, byte(x))
		}
	}
	return TF
}
