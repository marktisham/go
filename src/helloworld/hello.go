package main

// See https://tour.golang.org/list

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func TreeStuff2() {
	t := tree.New(1)
	fmt.Println(t.Value)
	// https://tour.golang.org/concurrency/8
}

func main() {
	//PlayWithArrays()
	//PlayWithStructs()
	//PlayWithMaps()
	//UseAnInterace()
	//PlayWithGoRoutines()
	//PlayWithChannels()
	//BufferChannels()
	//EquivTrees()
	BuildTree()

	// https://tour.golang.org
	// go build
	// or run specific files:
	// go hello.go goroutines.go
	// go fmt = reformat all code

	// to import an external lib you need to run this first:
	// go get
}
