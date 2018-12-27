package main

var mimeTypes = []string{"html", "css", "js", "png"}

type Resource struct {
	url      string
	size     int
	mimetype string
}
