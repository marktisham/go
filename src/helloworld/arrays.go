package main

import (
	"fmt"
	"sort"
)

// Array Fun

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func PlayWithArrays() {
	fmt.Printf("\nArrays!\n")
	var myslice []int
	printSlice(myslice)
	myslice = append(myslice, 123)
	printSlice(myslice)
	myslice = append(myslice, 12345)
	printSlice(myslice)

	fmt.Printf("\nOutputing the array\n")
	for i, v := range myslice {
		fmt.Printf("index %d = %d\n", i, v)
	}
}

func DeleteRecurringChars(str string) string {
	fmt.Println(str)
	var lookup = make(map[uint8]bool)
	var result = ""
	for i := 0; i < len(str); i++ {
		c := str[i]
		if lookup[c] == false {
			result += string(c)
			lookup[c] = true
		}
	}

	return result
}

func SortAndSearch() {
	a := []int{12, 222, 23, 11, 121}
	for _, v := range a {
		fmt.Printf("%d,", v)
	}
	fmt.Println()

	sort.Ints(a)
	for _, v := range a {
		fmt.Printf("%d,", v)
	}
	fmt.Println()

	res := sort.SearchInts(a, 121)
	fmt.Println(res)
}
