package main

import (
	"fmt"

	"github.com/andrievsky/sunday/structs"
)

func main() {
	// tries performance test
	var trie = structs.NewSimpleTrie()
	trie.Add("a")
	fmt.Println(trie.Count())
}
