package trie

type Trie struct {
	char     rune
	children []*Trie
	isWord   bool
}

func NewTrie() *Trie {
	return &Trie{
		children: []*Trie{},
	}
}

func NewTrieWithChar(char rune) *Trie {
	return &Trie{
		char:     char,
		children: []*Trie{},
	}
}

func (tr *Trie) getChild(char rune) (*Trie, bool) {
	for _, child := range tr.children {
		if child.char == char {
			return child, true
		}
	}
	return nil, false
}

func (tr *Trie) addChild(char rune) *Trie {
	child := NewTrieWithChar(char)
	tr.children = append(tr.children, child)
	return child
}

func (tr *Trie) Count() int {
	count := 0
	for _, child := range tr.children {
		if child.isWord {
			count++
		}
		count += child.Count()
	}
	return count
}

func (tr *Trie) Add(word string) *Trie {
	chars := []rune(word)
	curr := tr

	for i := 0; i < len(chars); i++ {
		child, ok := curr.getChild(chars[i])
		if ok {
			curr = child
		} else {
			curr = curr.addChild(chars[i])
		}
	}

	curr.isWord = true

	return curr
}

func (tr *Trie) FindNode(word string) *Trie {
	chars := []rune(word)
	curr := tr

	for i := 0; i < len(chars); i++ {
		child, ok := curr.getChild(chars[i])
		if ok {
			curr = child
		} else {
			return nil
		}
	}

	return curr
}

func (tr *Trie) Find(word string) *Trie {
	node := tr.FindNode(word)
	if node != nil && node.isWord {
		return node
	}
	return nil
}

func (tr *Trie) FindWordsWithPrefix(prefix string) []string {
	node := tr.FindNode(prefix)
	if node == nil {
		return []string{}
	}

	words := make([]string, 0)
	for _, child := range node.children {
		child.words([]rune(prefix), &words)
	}
	return words
}

func (tr *Trie) words(prefix []rune, words *[]string) {
	prefix = append(prefix, tr.char)
	if tr.isWord {
		*words = append(*words, string(prefix))
	}
	for _, child := range tr.children {
		child.words(prefix, words)
	}
}
