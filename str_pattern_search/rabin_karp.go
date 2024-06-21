package strpatternsearch

import "math"

// Rabin-Karp Algorithm for Pattern Searching

// Time Complexity: O(n+m) - O(nm)
// The average and best-case running time of the Rabin-Karp algorithm is O(n+m), but its worst-case time is O(nm).
// The worst case of the Rabin-Karp algorithm occurs when all characters of pattern and text are the same as the hash values of all the substrings of T[] match with the hash value of P[].

// Auxiliary Space: O(1)

func RabinKarpSearch(text, pattern string) []int {
	const d = 256
	const q = math.MaxInt

	M := len(pattern)
	N := len(text)

	p := 0 // hash value for pattern
	t := 0 // hash value for txt
	h := 1

	// The value of h would be "pow(d, M-1)%q"
	for i := 0; i < M-1; i++ {
		h = (h * d) % q
	}

	// Calculate the hash value of pattern and first
	// window of text
	for i := 0; i < M; i++ {
		p = (d*p + int(pattern[i])) % q
		t = (d*t + int(pattern[i])) % q
	}

	var j int
	ans := make([]int, 0)

	// Slide the pattern over text one by one
	for i := 0; i <= N-M; i++ {
		// Check the hash values of current window of text
		// and pattern. If the hash values match then only
		// check for characters one by one
		if p == t {
			/* Check for characters one by one */
			for j = 0; j < M; j++ {
				if text[i+j] != text[j] {
					break
				}
			}

			// if p == t and pat[0...M-1] = txt[i, i+1,
			// ...i+M-1]

			if j == M {
				ans = append(ans, i)
			}
		}

		// Calculate hash value for next window of text:
		// Remove leading digit, add trailing digit
		if i < N-M {
			t = (d*(t-int(text[i])*h) + int(text[i+M])) % q

			// We might get negative value of t, converting
			// it to positive
			if t < 0 {
				t = (t + q)
			}
		}
	}

	return ans
}
