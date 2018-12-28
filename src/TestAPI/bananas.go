package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Banana struct {
	// Parm names need to be capitalized for json marshalling!
	// (e.g. public)
	// see https://stackoverflow.com/questions/26327391/json-marshalstruct-returns
	Number int
	Color string
}

func GetBananas(w http.ResponseWriter, r *http.Request) {
	var bananas []Banana
	params := mux.Vars(r)
	lc:=r.FormValue("lc")
	notexist:=params["no"]	// returns as null
	fmt.Printf("%s %s",lc, notexist)

	b:=Banana{1,"green"}
	bananas=append(bananas,b)
	b=Banana{2,"yellow"}
	bananas=append(bananas,b)

	json.NewEncoder(w).Encode(bananas)
}

func GetBanana(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	idv:=params["id"]
	clr:=params["color"]
	num,_:=strconv.Atoi(idv)

	b:=Banana{Number:num,Color:clr}

	// Response as JSON payload
	json.NewEncoder(w).Encode(b)
}

func CreateBanana(w http.ResponseWriter, r *http.Request) {
	var banana Banana
	_ = json.NewDecoder(r.Body).Decode(&banana)

	banana.Number*=2
	banana.Color=banana.Color + " (Doubled)"

	json.NewEncoder(w).Encode(banana)
}