package strpatternsearch

// RabinKarpSearch returns byte offsets of all matches, including overlaps.
// An empty pattern matches every boundary, including len(text).
// Expected time O(n+m), worst case O(n*m); auxiliary space O(1), excluding results.
func RabinKarpSearch(text, pattern string) []int {
	const base, modulus uint64 = 256, 1_000_000_007
	n, m := len(text), len(pattern)
	matches := make([]int, 0)
	if m > n {
		return matches
	}
	if m == 0 {
		for i := 0; i <= n; i++ {
			matches = append(matches, i)
		}
		return matches
	}
	var patternHash, windowHash uint64
	power := uint64(1)
	for i := 0; i < m; i++ {
		patternHash = (patternHash*base + uint64(pattern[i])) % modulus
		windowHash = (windowHash*base + uint64(text[i])) % modulus
		if i > 0 {
			power = power * base % modulus
		}
	}
	for i := 0; i <= n-m; i++ {
		if patternHash == windowHash && text[i:i+m] == pattern {
			matches = append(matches, i)
		}
		if i < n-m {
			removed := uint64(text[i]) * power % modulus
			windowHash = ((windowHash+modulus-removed)*base + uint64(text[i+m])) % modulus
		}
	}
	return matches
}
