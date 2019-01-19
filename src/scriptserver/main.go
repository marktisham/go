package main

import (
	"fmt"
	"log"
	"net/http"
)

var loadcount int

func htmlBody(w http.ResponseWriter, r *http.Request) {
	loadcount++

	// This simulates the one-time include the customer would add to their head
	// This include has short caching and injects an include to a bigger/longer
	// cached include. This allows us to version bump without customer change.
	str := "<html><head><script src=\"customerinclude.js\"></script></head><body><h1>Hello There</h1></body></html>"
	fmt.Fprintf(w, str)
}

// This would be the super lightweight JS include the customer would
// make. Little to no caching on this one, but it injects a bigger include
// that would use caching.
func customerInclude(w http.ResponseWriter, r *http.Request) {
	str := "console.log(\"customer include loading...\");"
	str += "console.log(\"loadcount=" + fmt.Sprintf("%d", loadcount) + "\");"
	str += "var s=document.createElement(\"script\");"

	// Simulate a cache busting version change, when # of loads reaches
	// a level, bump to a new include
	if loadcount >= 5 {
		str += "s.setAttribute(\"src\",\"biginclude2.js\");"
	} else {
		str += "s.setAttribute(\"src\",\"biginclude1.js\");"
	}

	str += "var h=document.getElementsByTagName(\"head\");"
	str += "h[0].appendChild(s);"

	str += "console.log(\"customer include loaded\");"
	//w.Header().Add("cache-control", "max-age=5")
	fmt.Fprintf(w, str)
}

// This is the original "v1" large include that would be cached indefinitely
func biginclude1(w http.ResponseWriter, r *http.Request) {
	str := "console.log(\"big include 1 loaded\");"
	w.Header().Add("cache-control", "max-age=6000")
	fmt.Fprintf(w, str)
}

// And here's the "v2" large include that would later be pulled by the common include
func biginclude2(w http.ResponseWriter, r *http.Request) {
	str := "console.log(\"big include 2 loaded\");"
	w.Header().Add("cache-control", "max-age=6000")
	fmt.Fprintf(w, str)
}

func main() {
	loadcount = 0
	mux := http.NewServeMux()
	mux.HandleFunc("/", htmlBody)
	mux.HandleFunc("/customerinclude.js", customerInclude)
	mux.HandleFunc("/biginclude1.js", biginclude1)
	mux.HandleFunc("/biginclude2.js", biginclude2)

	log.Fatal(http.ListenAndServe(":8081", mux))
}
