package main

// See https://tour.golang.org/list

import (
"fmt"
"strconv"
)

type MyStuff struct {
	name string
	value int
}

func PrintMyStuff(m MyStuff) {
	fmt.Printf(m.name + ": " + strconv.Itoa(m.value) + "\n")
}

func addStuff(a int,b int) (int,string) {
	return a+b,"Behold!"
}

func noChangeStuff(m MyStuff) {
	// Pass by copy, no change
	m.name="NOCHANGE"
}

func doChangeStuff(m *MyStuff) {
	// Pass by ref, yes change
	m.name="NEWVALUE"
}

func PlayWithStructs() {
	fmt.Printf("\nStructs!\n")

	addedVal,response := addStuff(5,3)
	mystuff:=MyStuff { response, addedVal }
	PrintMyStuff(mystuff)

	noChangeStuff(mystuff)
	PrintMyStuff(mystuff)

	doChangeStuff(&mystuff)
	PrintMyStuff(mystuff)
}

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
	PlayWithStructs()
	PlayWithArrays()

	// Continue here: https://tour.golang.org/moretypes/19

}

