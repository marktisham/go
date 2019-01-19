package main

import (
	"fmt"
	"log"
	"net/http"
)

func htmlBody(w http.ResponseWriter, r *http.Request) {
	str := "<html><head><script src=\"script1.js\"></script></head><body><h1>Hello There</h1></body></html>"
	fmt.Fprintf(w, str)
}

func script1(w http.ResponseWriter, r *http.Request) {
	// TODO...
	vers := r.URL.Query().Get("v")

	str := "console.log(\"script 1 loading\");"
	str += "//" + vers + "\n"
	str += "var s=document.createElement(\"script\");"

	if vers == "3" {
		str += "s.setAttribute(\"src\",\"script3.js\");"
	} else {
		str += "s.setAttribute(\"src\",\"script2.js\");"
	}

	str += "var h=document.getElementsByTagName(\"head\");"
	str += "h[0].appendChild(s);"

	str += "console.log(\"script 1 loaded\");"
	//w.Header().Add("cache-control", "max-age=5")
	fmt.Fprintf(w, str)
}

func script2(w http.ResponseWriter, r *http.Request) {
	str := "console.log(\"script 2 loaded\");"

	w.Header().Add("cache-control", "max-age=60")
	fmt.Fprintf(w, str)
}

func script3(w http.ResponseWriter, r *http.Request) {
	str := "console.log(\"script 3 loaded\");"

	w.Header().Add("cache-control", "max-age=60")
	fmt.Fprintf(w, str)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", htmlBody)
	mux.HandleFunc("/script1.js", script1)
	mux.HandleFunc("/script2.js", script2)
	mux.HandleFunc("/script3.js", script3)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
