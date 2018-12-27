package main

func main() {

	chItems := make(chan Resource, 5)
	go loadItems(chItems)
	AnalyzeItems(chItems)
}
