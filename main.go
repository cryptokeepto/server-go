package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type contact struct {
	Name string `json:"name"`
	Tel  string `json:"tel"`
}

func main() {
	http.HandleFunc("/", all)
	http.HandleFunc("/product", product)

	http.HandleFunc("/save", post)

	http.ListenAndServe(":8080", nil)
}

func all(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "all")
}

func product(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "product")
	// c := contact{Name: "sittikiat", Tel: "0992828286"}
	// json.NewEncoder(w).Encode(c)
}

func post(w http.ResponseWriter, r *http.Request) {
	c := contact{}
	json.NewDecoder(r.Body).Decode(&c)
	fmt.Println(c)
}
