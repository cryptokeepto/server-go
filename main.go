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
	http.HandleFunc("/", get)
	http.HandleFunc("/save", post)

	http.ListenAndServe(":8080", nil)
}

func get(w http.ResponseWriter, r *http.Request) {
	c := contact{Name: "sittikiat", Tel: "0992828286"}
	json.NewEncoder(w).Encode(c)
}

func post(w http.ResponseWriter, r *http.Request) {
	c := contact{}
	json.NewDecoder(r.Body).Decode(&c)
	fmt.Println(c)
}
