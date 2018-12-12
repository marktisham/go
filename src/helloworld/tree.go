package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

//
// Exercise: https://tour.golang.org/concurrency/7
//
// Solutions:https://gist.github.com/zyxar/2317744s

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

func EquivTrees() {

	// See https://gist.github.com/zyxar/2317744

	t1 := tree.New(2)
	t2 := tree.New(10)
	if Same(t1, t2) == true {
		fmt.Println("SAME!")
	} else {
		fmt.Println("Different!")
	}
}

//
// Building trees for number sorting
//

type MyTree struct {
	left  *MyTree
	right *MyTree
	val   int
}

func BuildTree() {
	// Arbitrary unsorted array
	a := []int{224, 35, 4, 66, 31, 11, 31, 22, -5, 1201, 18, 99, 18, 222, 123, 43, 12, 44}

	fmt.Printf("Original:\n")
	var root *MyTree
	for i, v := range a {
		fmt.Printf("%d,", v)

		if i == 0 {
			root = &MyTree{nil, nil, v}
		} else {
			InsertTree(root, v)
		}
	}

	ch := make(chan int)
	go TreeInOrder(root, ch)

	var vals []int
	for val := range ch {
		vals = append(vals, val)
	}

	fmt.Printf("\nFrom channel:\n")
	for _, v := range vals {
		fmt.Printf("%d,", v)
	}
	fmt.Printf("\nReversed:\n")
	for i := len(vals) - 1; i >= 0; i-- {
		fmt.Printf("%d,", vals[i])
	}

	const largestN int = 5
	fmt.Printf("\nLargest %d:\n", largestN)
	for i := len(vals) - 1; i >= len(vals)-1-largestN; i-- {
		fmt.Printf("%d,", vals[i])
	}
}

func InsertTree(parent *MyTree, val int) *MyTree {
	if val < parent.val {
		if parent.left != nil {
			return InsertTree(parent.left, val)
		}

		parent.left = &MyTree{}
		parent.left.val = val
		return parent.left
	} else {
		if parent.right != nil {
			return InsertTree(parent.right, val)
		}
		parent.right = &MyTree{}
		parent.right.val = val
		return parent.right
	}
}

func TreeInOrder(parent *MyTree, ch chan int) {
	_TreeInOrder(parent, ch)
	close(ch)
}

func _TreeInOrder(parent *MyTree, ch chan int) {
	if parent == nil {
		return
	}
	_TreeInOrder(parent.left, ch)

	ch <- parent.val

	_TreeInOrder(parent.right, ch)

}
