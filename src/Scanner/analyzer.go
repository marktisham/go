package main

import (
	"fmt"
	"sync"
	"time"
)

func AnalyzeItems(ch chan Resource) {
	var wg sync.WaitGroup

	for item := range ch {
		wg.Add(1)
		go _AnalyzeItem(item, &wg)
	}

	fmt.Printf("ANALYZER: No more new items.\n")
	wg.Wait()
	fmt.Printf("ANALYZER: Done analyzing items\n")
}

func _AnalyzeItem(item Resource, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("ANALYZER: Analyzing: %s...\n", item.url)
	time.Sleep(1 * time.Second)
	fmt.Printf("ANALYZER: Done analyzing: %s\n", item.url)
}
