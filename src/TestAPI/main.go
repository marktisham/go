package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/bananas", GetBananas).Methods("GET")
	router.HandleFunc("/bananas", CreateBanana).Methods("POST")
	router.HandleFunc("/bananas/{id}/{color}", GetBanana).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}