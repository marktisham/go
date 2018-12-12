package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)

	val := t.Value
	//ch <- val
	fmt.Printf("%d,", val)

	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return false
}

func TreeStuff() {
	ch := make(chan int)
	t1 := tree.New(10)
	t2 := tree.New(10)
	Walk(t1, ch)
	fmt.Printf("\n")
	Walk(t2, ch)
}
