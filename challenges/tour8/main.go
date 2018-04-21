package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go WalkAndClose(t1, ch1)
	go WalkAndClose(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if !ok1 && !ok2 {
			return true
		}
		if v1 != v2 {
			return false
		}
	}
}

func WalkAndClose(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

func main() {
	tree1 := tree.New(1)
	tree2 := tree.New(1)
	fmt.Println(Same(tree1, tree2))

	tree1 = tree.New(1)
	tree2 = tree.New(2)
	fmt.Println(Same(tree1, tree2))

}
