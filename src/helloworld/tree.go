package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int, sender bool) bool {
	if t == nil {
		return true
	}
	if Walk(t.Left, ch, sender) == false {
		return false
	}

	if sender == true {
		ch <- t.Value
	} else {
		testVal := <-ch
		if testVal != t.Value {
			return false
		}
	}

	if Walk(t.Right, ch, sender) == false {
		return false
	}

	return true
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch := make(chan int)
	go Walk(t1, ch, true)

	return Walk(t2, ch, false)
}

func TreeStuff() {

	t1 := tree.New(2)
	t2 := tree.New(10)
	if Same(t1, t2) == true {
		fmt.Println("SAME!")
	} else {
		fmt.Println("Different!")
	}
}
