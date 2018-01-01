package structs_test

import (
	"testing"

	"fmt"

	"github.com/andrievsky/sunday/structs"
	"github.com/stretchr/testify/assert"
)

func setupTries() []structs.Trie {
	return []structs.Trie{structs.NewNaiveTrie(), structs.NewSimpleTrie()}
}

func TestAddEmptyItem(t *testing.T) {
	var tries = setupTries()
	for _, trie := range tries {
		trie.Add("")
		assert.Equal(t, 0, trie.Count())
	}
}

func TestAddItem(t *testing.T) {
	var tries = setupTries()
	for _, trie := range tries {
		trie.Add("a")
		assert.Equal(t, 1, trie.Count())
	}
}

func TestLookupItem(t *testing.T) {
	var tries = setupTries()
	for _, trie := range tries {
		trie.Add("a")
		assert.Equal(t, "a", trie.Lookup("a")[0], fmt.Sprintf("Trie type %T", trie))
	}
}

func TestAddSameItems(t *testing.T) {
	var tries = setupTries()
	for _, trie := range tries {
		trie.Add("a")
		trie.Add("a")
		assert.Equal(t, 1, trie.Count())
		trie.Add("b")
		trie.Add("b")
		assert.Equal(t, 2, trie.Count())
	}
}

func TestLookupItems(t *testing.T) {
	var tries = setupTries()
	for _, trie := range tries {
		trie.Add("one")
		trie.Add("two")
		trie.Add("three")
		trie.Add("four")
		trie.Add("five")
		assert.Equal(t, 5, trie.Count())
		var res = trie.Lookup("f")
		assert.Equal(t, "four", res[0])
		assert.Equal(t, "five", res[1])
	}
}
