package main

// See https://tour.golang.org/list

import (
"fmt"
//"strconv"
)

// Array Fun

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func PlayWithArrays() {
	fmt.Printf("\nArrays!\n")
	var myslice []int
	printSlice(myslice)
	myslice=append(myslice,123)
	printSlice(myslice)
	myslice=append(myslice,12345)
	printSlice(myslice)

	fmt.Printf("\nOutputing the array\n")
	for i, v := range myslice {
		fmt.Printf("index %d = %d\n", i, v)
	}
}

func main() {
	//PlayWithStructs()
	//PlayWithArrays()
	//PlayWithMaps()
	//UseAnInterace()

	// Continue here: https://tour.golang.org

}

