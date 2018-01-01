package structs

// Trie iterface
type Trie interface {
	Add(value string)
	Lookup(value string) []string
	Count() int
}
