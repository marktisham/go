package main

import (
	"fmt"
	"strconv"
)

type MyStuff struct {
	name  string
	value int
}

// Play with methods
func (m MyStuff) DoubleVal() int {
	return m.value * 2
}

func (m *MyStuff) DoubleIt() {
	m.value *= 2
	m.name += " doubled"
}

//
// Play with interfaces
//
type Doubler interface {
	DoubleVal() int
}

type ASecondStruct struct {
	myval int
}
type AStructWithoutIntfc struct {
	otherval int
}

func (s ASecondStruct) DoubleVal() int {
	return s.myval * 2
}

func UseAnInterace() {
	m := MyStuff{"mystuff via interface", 12}
	var myintfc Doubler

	myintfc = m
	val := myintfc.DoubleVal()
	fmt.Printf("MyStuff via interface: " + strconv.Itoa(val) + "\n")

	a := ASecondStruct{55}
	myintfc = a
	val = myintfc.DoubleVal()
	fmt.Printf("ASecondStruct via interface: " + strconv.Itoa(val) + "\n")

	// A struct that does not support the interface should fail build:
	/*
		s3:=AStructWithoutIntfc{12}
		myintfc=s3	// this should fail
	*/
}

// Play with functions
func PrintMyStuff(m MyStuff) {
	fmt.Printf(m.name + ": " + strconv.Itoa(m.value) + "\n")
}

func addStuff(a int, b int) (int, string) {
	return a + b, "Behold!"
}

func noChangeStuff(m MyStuff) {
	// Pass by copy, no change
	m.name = "NOCHANGE"
}

func doChangeStuff(m *MyStuff) {
	// Pass by ref, yes change
	m.name = "NEWVALUE"
}

func PlayWithStructs() {
	fmt.Printf("\nStructs!\n")

	addedVal, response := addStuff(5, 3)
	mystuff := MyStuff{response, addedVal}
	PrintMyStuff(mystuff)

	noChangeStuff(mystuff)
	PrintMyStuff(mystuff)

	doChangeStuff(&mystuff)
	PrintMyStuff(mystuff)

	fmt.Printf("doubling it...\n")
	fmt.Printf(strconv.Itoa(mystuff.DoubleVal()) + "\n")
	mystuff.DoubleIt()
	PrintMyStuff(mystuff)
}

// Maps
func PlayWithMaps() {
	mymap := make(map[int]MyStuff)
	mymap[3] = MyStuff{"twenty at 3", 20}
	mymap[12] = MyStuff{"fifty at twelve", 50}

	for i, v := range mymap {
		fmt.Printf("Index: " + strconv.Itoa(i) + "\n")
		fmt.Printf("Value: " + v.name + ", " + strconv.Itoa(v.value) + "\n")
	}
}
