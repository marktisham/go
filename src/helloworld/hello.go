package main

import "fmt"

// See https://tour.golang.org/list

func PanicAndRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered!", r)
		}
	}()
	panic("Uh Oh I panicked!")
}

func main() {
	/*
		PlayWithArrays()
		PlayWithStructs()
		PlayWithMaps()
		UseAnInterace()
		PlayWithGoRoutines()
		PlayWithChannels()
		PlayWithChannels2()
		BufferChannels()
		EquivTrees()
		BuildTree()
		CountWords()
		str := DeleteRecurringChars("aaaabdbbdbdbabdbcbbbdbcbdbcbab")
		fmt.Println(str)

		PanicAndRecover()
	*/

	SortAndSearch()

	// https://tour.golang.org
	// go build
	// or run specific files:
	// go hello.go goroutines.go
	// go fmt = reformat all code

	// to import an external lib you need to run this first:
	// go get
}
