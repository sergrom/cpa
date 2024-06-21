package strpatternsearch

// Aho-Corasick Algorithm for Pattern Searching
// Time Complexity: O(n + l + z), where 'n' is the length of the text, 'l' is the length of keywords, and 'z' is the number of matches.
// Auxiliary Space: O(l * q), where 'q' is the length of the alphabet since that is the maximum number of children a node can have.

// searchWords - This function finds all occurrences of all array words in text.
// returns a map, key is word, value is slice of all occurences of the word [2]int{fromIdx, toIdx}
func AhoCorasickSearch(text string, words []string) map[string][][2]int {
	// Preprocess patterns.
	// Build machine with goto, failure and output functions
	_, out, f, g := buildMatchingMachine(words)

	// Initialize current state
	currentState := 0

	ans := make(map[string][][2]int)

	// Traverse the text through the built machine to find all occurrences of words in arr[]
	for i := 0; i < len(text); i++ {
		currentState = findNextState(currentState, text[i], f, g)

		// If match not found, move to next state
		if out[currentState] == 0 {
			continue
		}

		// Collect all matching words of words[] using output function.
		for j := 0; j < len(words); j++ {
			if out[currentState]&(1<<j) > 0 {
				ans[words[j]] = append(ans[words[j]], [2]int{i - len(words[j]) + 1, i})
			}
		}
	}

	return ans
}

// Builds the string matching machine.
// arr -   array of words. The index of each keyword is important:
//
//	"out[state] & (1 << i)" is > 0 if we just found word[i]
//	in the text.
//
// Returns the number of states that the built machine has.
// States are numbered 0 up to the return value - 1, inclusive.
//
// out - OUTPUT FUNCTION IS IMPLEMENTED USING out[]
// Bit i in this mask is one if the word with index i
// appears when the machine enters this state.
//
// f - FAILURE FUNCTION IS IMPLEMENTED USING f[]
//
// g - GOTO FUNCTION (OR TRIE) IS IMPLEMENTED USING g[][]
func buildMatchingMachine(words []string) (int, []int, []int, [][26]int) {
	maxs := 0
	for _, word := range words {
		maxs += len(word)
	}

	out := make([]int, maxs)

	f := make([]int, maxs)

	g := make([][26]int, maxs)

	// Initialize all values in goto function as -1.
	for i := range g {
		for j := range g[i] {
			g[i][j] = -1
		}
	}

	// Initially, we just have the 0 state
	states := 1

	// Construct values for goto function, i.e., fill g[][]
	// This is same as building a Trie for arr[]
	for i := 0; i < len(words); i++ {
		word := words[i]
		currentState := 0

		// Insert all characters of current word in arr[]
		for j := 0; j < len(word); j++ {
			ch := word[j] - 'a'

			// Allocate a new node (create a new state) if a
			// node for ch doesn't exist.
			if g[currentState][ch] == -1 {
				states++
				g[currentState][ch] = states
			}

			currentState = g[currentState][ch]
		}

		// Add current word in output function
		out[currentState] |= (1 << i)
	}

	// For all characters which don't have an edge from
	// root (or state 0) in Trie, add a goto edge to state 0 itself
	for ch := 0; ch < 26; ch++ {
		if g[0][ch] == -1 {
			g[0][ch] = 0
		}
	}

	// Now, let's build the failure function
	// Initialize values in fail function
	for i := range f {
		f[i] = -1
	}

	// Failure function is computed in breadth first order using a queue
	q := make([]int, 0)

	// Iterate over every possible input
	for ch := 0; ch < 26; ch++ {
		// All nodes of depth 1 have failure function value
		// as 0. For example, in above diagram we move to 0 from states 1 and 3.
		if g[0][ch] != 0 {
			f[g[0][ch]] = 0
			q = append(q, g[0][ch])
		}
	}

	// Now queue has states 1 and 3
	for len(q) > 0 {
		// Remove the front state from queue
		state := q[0]
		q = q[1:]

		// For the removed state, find failure function for
		// all those characters for which goto function is not defined.
		for ch := 0; ch < 26; ch++ {
			// If goto function is defined for character 'ch' and 'state'
			if g[state][ch] != -1 {
				// Find failure state of removed state
				failure := f[state]

				// Find the deepest node labeled by proper
				// suffix of string from root to current state.
				for g[failure][ch] == -1 {
					failure = f[failure]
				}

				failure = g[failure][ch]
				f[g[state][ch]] = failure

				// Merge output values
				out[g[state][ch]] |= out[failure]

				// Insert the next level node (of Trie) in Queue
				q = append(q, g[state][ch])
			}
		}
	}

	return states, out, f, g
}

// findNextState - Returns the next state the machine will transition to using goto
// and failure functions.
// currentState - The current state of the machine. Must be between
// 0 and the number of states - 1, inclusive.
// nextInput - The next character that enters into the machine.
func findNextState(currentState int, nextInput byte, f []int, g [][26]int) int {
	answer := currentState
	ch := nextInput - 'a'

	// If goto is not defined, use failure function
	for g[answer][ch] == -1 {
		answer = f[answer]
	}

	return g[answer][ch]
}
