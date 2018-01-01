package structs

import "strings"

// NaiveTrie impl
type NaiveTrie struct {
	entries []string
}

// NewNaiveTrie constructor
func NewNaiveTrie() *NaiveTrie {
	var trie = new(NaiveTrie)
	trie.entries = make([]string, 0, 16)
	return trie
}

// Add a value to the tree
func (t *NaiveTrie) Add(value string) {
	if value == "" || t.contains(value) {
		return
	}

	t.entries = append(t.entries, value)
}

// Lookup the value from the tree
func (t *NaiveTrie) Lookup(value string) []string {
	var res = []string{}
	for _, e := range t.entries {
		if strings.HasPrefix(e, value) {
			res = append(res, e)
		}
	}
	return res
}

// Count of tree elements
func (t *NaiveTrie) Count() int {
	return len(t.entries)
}

func (t *NaiveTrie) contains(entry string) bool {
	for _, e := range t.entries {
		if entry == e {
			return true
		}
	}
	return false
}
