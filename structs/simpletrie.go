package structs

type node struct {
	Children map[rune]*node
	source   string
}

func newNode() *node {
	return &node{Children: make(map[rune]*node, 1)}
}

// SimpleTrie implementation
type SimpleTrie struct {
	root  *node
	count int
}

// NewSimpleTrie constructor
func NewSimpleTrie() *SimpleTrie {
	var trie = new(SimpleTrie)
	trie.root = newNode()
	return trie
}

// Add a value to the tree
func (t *SimpleTrie) Add(value string) {
	if value == "" {
		return
	}
	var current = t.root
	for _, char := range value {
		if _, ok := current.Children[char]; !ok {
			current.Children[char] = newNode()
		}
		current = current.Children[char]
	}
	if current.source != "" {
		return
	}
	current.source = value
	t.count++
}

// Lookup the value from the tree
func (t *SimpleTrie) Lookup(value string) []string {
	var res = make([]string, 0)
	var node = t.root
	var ok bool
	for _, char := range value {
		if node, ok = node.Children[char]; !ok {
			return res
		}
	}
	res = getWords(node, res)

	return res
}

func getWords(node *node, res []string) []string {
	if node.source != "" {
		res = append(res, node.source)
	}

	for _, value := range node.Children {
		res = getWords(value, res)
	}
	return res
}

// Count of tree elements
func (t *SimpleTrie) Count() int {
	return t.count
}
