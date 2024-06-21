package strpatternsearch

// KMP (Knuth Morris Pratt) Pattern Searching
// Finds substring occurences with overlapping

// var pattern = "qweqwe"
// var text = "qweqweqwerrqweqwe"
//     found:  |  |       |        [0 3 11]

func KMPSearch(text, pattern string) []int {
	var M = len(pattern)
	var N = len(text)

	ans := make([]int, 0)

	// create lps[] that will hold the longest
	// prefix suffix values for pattern
	var j = 0 // index for pat[]

	// Preprocess the pattern (calculate lps[]
	// array)
	lps := computeLPSArray(pattern)

	var i = 0 // index for txt[]
	for N-i >= M-j {
		if pattern[j] == text[i] {
			j++
			i++
		}
		if j == M {
			ans = append(ans, i-j)
			// fmt.Println("Found pattern " + "at index " + strconv.Itoa(i-j))
			// document.write("Found pattern " + "at index " + (i - j) + "\n");
			j = lps[j-1]
		} else if i < N && pattern[j] != text[i] { // mismatch after j matches
			// Do not match lps[0..lps[j-1]] characters,
			// they will match anyway
			if j != 0 {
				j = lps[j-1]
			} else {
				i = i + 1
			}
		}
	}

	return ans
}

func computeLPSArray(pattern string) []int {
	// length of the previous longest prefix suffix
	var l = 0
	var i = 1
	lps := make([]int, len(pattern))
	lps[0] = 0 // lps[0] is always 0

	// the loop calculates lps[i] for i = 1 to M-1
	for i < len(pattern) {
		if pattern[i] == pattern[l] {
			l++
			lps[i] = l
			i++
		} else { // (pat[i] != pat[len])
			// This is tricky. Consider the example.
			// AAACAAAA and i = 7. The idea is similar
			// to search step.
			if l != 0 {
				l = lps[l-1]
				// Also, note that we do not increment
				// i here
			} else { // if (len == 0)
				lps[i] = l
				i++
			}
		}
	}

	return lps
}
