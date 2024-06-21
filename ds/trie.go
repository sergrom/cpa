package ds

// Префиксное дерево

const (
	// ALBHABET_SIZE total characters in english alphabet
	ALBHABET_SIZE = 26
)

type trieNode struct {
	childrens [ALBHABET_SIZE]*trieNode
	isWordEnd bool
}

type trie struct {
	root *trieNode
}

func initTrie() *trie {
	return &trie{
		root: &trieNode{},
	}
}

func (t *trie) Insert(word string) {
	wordLength := len(word)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.childrens[index] == nil {
			current.childrens[index] = &trieNode{}
		}
		current = current.childrens[index]
	}
	current.isWordEnd = true
}

func (t *trie) Search(word string) bool {
	wordLength := len(word)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.childrens[index] == nil {
			return false
		}
		current = current.childrens[index]
	}
	return current.isWordEnd
}

func (t *trie) StartsWith(prefix string) bool {
	wordLength := len(prefix)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := prefix[i] - 'a'
		if current.childrens[index] == nil {
			return false
		}
		current = current.childrens[index]
	}

	return true
}
