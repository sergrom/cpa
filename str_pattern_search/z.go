package strpatternsearch

// Z algorithm (Linear time pattern searching Algorithm)
// Time complexity: O(m + n)
// Space complexity: O(m + n)

// ZSearch returns byte offsets of all matches, including overlaps.
// An empty pattern matches every boundary, including len(text).
func ZSearch(text, pattern string) []int {
	ans := make([]int, 0)

	if len(pattern) == 0 {
		for i := 0; i <= len(text); i++ {
			ans = append(ans, i)
		}
		return ans
	}
	// The separator may also occur in the input: only text offsets are checked,
	// and matching at least len(pattern) bytes is sufficient.
	// Construct Z array
	Z := GetZarr(pattern + "$" + text)

	// now looping through Z array for matching condition
	for i := len(pattern) + 1; i < len(Z); i++ {
		// Longer prefix matches may cross the separator; they still match pattern.
		if Z[i] >= len(pattern) {
			ans = append(ans, i-len(pattern)-1)
		}
	}

	return ans
}

// Fills Z array for given string str[]
func GetZarr(str string) []int {
	n := len(str)
	var L, R, k int

	zarr := make([]int, n)

	// [L,R] make a window which matches with prefix of s
	L, R = 0, 0
	for i := 1; i < n; i++ {
		// if i>R nothing matches so we will calculate.
		// Z[i] using naive way.
		if i > R {
			L, R = i, i

			// R-L = 0 in starting, so it will start
			// checking from 0'th index. For example,
			// for "ababab" and i = 1, the value of R
			// remains 0 and Z[i] becomes 0. For string
			// "aaaaaa" and i = 1, Z[i] and R become 5
			for R < n && str[R-L] == str[R] {
				R++
			}
			zarr[i] = R - L
			R--
		} else {
			// k = i-L so k corresponds to number which
			// matches in [L,R] interval.
			k = i - L

			// if Z[k] is less than remaining interval
			// then Z[i] will be equal to Z[k].
			// For example, str = "ababab", i = 3, R = 5
			// and L = 2
			if zarr[k] < R-i+1 {
				zarr[i] = zarr[k]
			} else {
				// For example str = "aaaaaa" and i = 2, R is 5,
				// L is 0

				// else start from R and check manually
				L = i
				for R < n && str[R-L] == str[R] {
					R++
				}
				zarr[i] = R - L
				R--
			}
		}
	}

	return zarr
}
