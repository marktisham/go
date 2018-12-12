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
	/*
		PlayWithArrays()
		PlayWithStructs()
		PlayWithMaps()
		UseAnInterace()
	*/
	//PlayWithGoRoutines()
	//PlayWithChannels()
	//BufferChannels()
	TreeStuff()

	// Continue here: https://tour.golang.org
	// go builder:
	// cmd-9
	// go build
	// or
	// go hello.go goroutines.go

	// to import an external lib you need to run
	// go get
}
