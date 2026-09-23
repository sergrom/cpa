package strpatternsearch

import (
	"math/rand"
	"reflect"
	"strings"
	"testing"
)

func naiveMatches(text, pattern string) []int {
	result := make([]int, 0)
	for i := 0; i+len(pattern) <= len(text); i++ {
		if text[i:i+len(pattern)] == pattern {
			result = append(result, i)
		}
	}
	return result
}

func TestRabinKarpAndZ(t *testing.T) {
	cases := [][2]string{{"abc", "z"}, {"xab", "ab"}, {"$$", "$"}, {"aaaa", "aa"}, {"", ""}, {"abc", ""}, {"a", "long"}, {"\x00\xff$\x00", "\x00"}}
	alphabet := []byte{'a', 'b', '$', 0, 255}
	rng := rand.New(rand.NewSource(10))
	for i := 0; i < 1500; i++ {
		text, pattern := make([]byte, rng.Intn(60)), make([]byte, rng.Intn(15))
		for j := range text {
			text[j] = alphabet[rng.Intn(len(alphabet))]
		}
		for j := range pattern {
			pattern[j] = alphabet[rng.Intn(len(alphabet))]
		}
		if i%3 == 0 && len(text) > 0 {
			a := rng.Intn(len(text))
			pattern = text[a : a+rng.Intn(len(text)-a+1)]
		}
		cases = append(cases, [2]string{string(text), string(pattern)})
	}
	cases = append(cases, [2]string{strings.Repeat("abcdefgh", 1000), strings.Repeat("abcdefgh", 10)})
	for _, tc := range cases {
		want := naiveMatches(tc[0], tc[1])
		for name, fn := range map[string]func(string, string) []int{"RabinKarp": RabinKarpSearch, "Z": ZSearch} {
			if got := fn(tc[0], tc[1]); !reflect.DeepEqual(got, want) {
				t.Fatalf("%s(%q,%q)=%v want %v", name, tc[0], tc[1], got, want)
			}
		}
	}
}

func naiveAho(text string, words []string) map[string][][2]int {
	want := make(map[string][][2]int)
	seen := make(map[string]bool)
	for _, word := range words {
		if seen[word] {
			continue
		}
		seen[word] = true
		for _, i := range naiveMatches(text, word) {
			want[word] = append(want[word], [2]int{i, i + len(word) - 1})
		}
	}
	return want
}

func TestAhoCorasick(t *testing.T) {
	cases := []struct {
		text  string
		words []string
	}{
		{"a", []string{"a"}}, {"", nil}, {"abc", nil}, {"", []string{""}},
		{"ahishers", []string{"he", "she", "hers", "his"}},
		{"aaaa", []string{"a", "aa", "aaa", "a", ""}},
		{"ab#aba", []string{"a", "ab", "aba", "ba", ""}},
	}
	words := make([]string, 100)
	for i := range words {
		words[i] = string([]byte{'a' + byte(i/26), 'a' + byte(i%26)})
	}
	cases = append(cases, struct {
		text  string
		words []string
	}{strings.Join(words, "#"), words})
	rng := rand.New(rand.NewSource(11))
	for i := 0; i < 300; i++ {
		text := make([]byte, rng.Intn(60))
		for j := range text {
			text[j] = 'a' + byte(rng.Intn(3))
		}
		patterns := make([]string, rng.Intn(20))
		for j := range patterns {
			b := make([]byte, rng.Intn(7))
			for k := range b {
				b[k] = 'a' + byte(rng.Intn(3))
			}
			patterns[j] = string(b)
		}
		cases = append(cases, struct {
			text  string
			words []string
		}{string(text), patterns})
	}
	for _, tc := range cases {
		want := naiveAho(tc.text, tc.words)
		if got := AhoCorasickSearch(tc.text, tc.words); !reflect.DeepEqual(got, want) {
			t.Fatalf("Aho(%q,%v)=%v want %v", tc.text, tc.words, got, want)
		}
	}
}

func TestAhoInvalidPattern(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("expected invalid pattern panic")
		}
	}()
	AhoCorasickSearch("abc", []string{"A"})
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
func TestLongestPalindrome(t *testing.T) {
	cases := []string{"", "a", "ab", "abba", "babad", "\x00a\x00"}
	for n := 1; n <= 9; n++ {
		for mask := 0; mask < 1<<n; mask++ {
			b := make([]byte, n)
			for i := range b {
				b[i] = 'a' + byte((mask>>i)&1)
			}
			cases = append(cases, string(b))
		}
	}
	for _, text := range cases {
		longest := 0
		for i := 0; i < len(text); i++ {
			for j := i + 1; j <= len(text); j++ {
				if isPalindrome(text[i:j]) {
					longest = max(longest, j-i)
				}
			}
		}
		got := findLongestPalindromicString(text)
		if !strings.Contains(text, got) || !isPalindrome(got) || len(got) != longest {
			t.Fatalf("palindrome(%q)=%q, want length %d", text, got, longest)
		}
	}
}
