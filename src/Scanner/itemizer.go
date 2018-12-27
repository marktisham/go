package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func doHttpStuff() {
	resp, err := http.Get("https://rigor.com")
	if err != nil {
		fmt.Printf("Bad request! %s\n", err)
		return
	}
	defer resp.Body.Close()

	// "go get" to get the ioutil package
	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}
}

func loadItems(ch chan Resource) {
	doHttpStuff()
	return

	// Simulating loading an HTML file
	const resourceCount int = 10

	for i := 0; i < resourceCount; i++ {
		for _, mimeType := range mimeTypes {
			loadSpecificItem(ch, i, mimeType)
		}
	}
	close(ch)
}

func loadSpecificItem(ch chan Resource, resource_id int, mimeType string) {
	const resourceSize int = 112000

	url := fmt.Sprintf("https://test.com/resources/%d.%s",
		resource_id, mimeType)
	res := Resource{url, resourceSize, mimeType}
	fmt.Printf("LOADER: Found resource: %s\n", url)

	ch <- res

	time.Sleep(100 * time.Millisecond)
}
