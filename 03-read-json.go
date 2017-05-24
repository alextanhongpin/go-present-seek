package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Book struct {
	Author string `json:"author"`
	Post   string `json:"post"`
}

func main() {
	http.HandleFunc("/books", postHandler)
	http.ListenAndServe(":8080", nil)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var book Book
		json.NewDecoder(r.Body).Decode(&book)

		fmt.Println(book)
		// Do something with book
	} else if r.Method == "GET" {
		queries := r.URL.Query()
		fmt.Println(queries.Get("name"))
	}
}
