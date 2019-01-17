package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

// Paragraph object
type Paragraph struct {
	text string
}

func (p Paragraph) CountWords() int {
	var s scanner.Scanner
	s.Init(strings.NewReader(p.text))
	count := 0
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
		count++
	}
	return count
}

func (p Paragraph) CountUniqueWords() int {
	wordCount := make(map[string]int)

	var s scanner.Scanner
	s.Init(strings.NewReader(p.text))
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		//fmt.Printf("%s: %s\n", s.Position, s.TokenText()
		lower := strings.ToLower(s.TokenText())
		wordCount[lower]++
	}

	fmt.Printf("\nUnique words follow:\n")
	for key, val := range wordCount {
		fmt.Printf("%s: %d\n", key, val)
	}

	return len(wordCount)
}

type Counter interface {
	CountWords() int
	CountUniqueWords() int
}

func CountWords() {
	p := Paragraph{"Once upon a time time was upon it once"}
	var c Counter = p
	wordCount := c.CountWords()
	fmt.Printf("\nWords in %s:\n%d\n", p.text, wordCount)

	uniqueWordCount := c.CountUniqueWords()
	fmt.Printf("\nUnique Words in %s:\n%d\n", p.text, uniqueWordCount)
}
