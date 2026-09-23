package strpatternsearch

// AhoCorasickSearch finds all occurrences, including overlaps, of lowercase
// ASCII patterns. Each result is an inclusive pair of byte offsets [start, end].
// Text bytes outside a-z break matches. Invalid pattern bytes cause a panic.
// Duplicate patterns are reported once; an empty pattern matches every boundary
// as [i, i-1]. There is no machine-word limit on the number of patterns.
// Time O(total pattern length * 26 + len(text) + matches), space O(states * 26),
// excluding the result. Output links avoid copying suffix matches into each node.
func AhoCorasickSearch(text string, words []string) map[string][][2]int {
	nodes := []ahoNode{newAhoNode()}
	for _, word := range words {
		state := 0
		for i := range len(word) {
			if word[i] < 'a' || word[i] > 'z' {
				panic("AhoCorasickSearch: patterns must contain only a-z")
			}
			ch := word[i] - 'a'
			if nodes[state].next[ch] == -1 {
				nodes[state].next[ch] = len(nodes)
				nodes = append(nodes, newAhoNode())
			}
			state = nodes[state].next[ch]
		}
		nodes[state].terminal = true
		nodes[state].word = word
	}
	queue := make([]int, 0)
	for ch := range 26 {
		if next := nodes[0].next[ch]; next != -1 {
			queue = append(queue, next)
		} else {
			nodes[0].next[ch] = 0
		}
	}
	for head := 0; head < len(queue); head++ {
		state := queue[head]
		failure := nodes[state].failure
		if failure != 0 && nodes[failure].terminal {
			nodes[state].output = failure
		} else {
			nodes[state].output = nodes[failure].output
		}
		for ch := range 26 {
			next := nodes[state].next[ch]
			if next == -1 {
				nodes[state].next[ch] = nodes[failure].next[ch]
			} else {
				nodes[next].failure = nodes[failure].next[ch]
				queue = append(queue, next)
			}
		}
	}
	matches := make(map[string][][2]int)
	if nodes[0].terminal {
		matches[""] = append(matches[""], [2]int{0, -1})
	}
	state := 0
	for i := range len(text) {
		if text[i] < 'a' || text[i] > 'z' {
			state = 0
		} else {
			state = nodes[state].next[text[i]-'a']
		}
		for node := state; node > 0; node = nodes[node].output {
			if nodes[node].terminal {
				word := nodes[node].word
				matches[word] = append(matches[word], [2]int{i - len(word) + 1, i})
			}
		}
		if nodes[0].terminal {
			matches[""] = append(matches[""], [2]int{i + 1, i})
		}
	}
	return matches
}

type ahoNode struct {
	next     [26]int
	failure  int
	output   int
	terminal bool
	word     string
}

func newAhoNode() ahoNode {
	node := ahoNode{output: -1}
	for i := range node.next {
		node.next[i] = -1
	}
	return node
}
