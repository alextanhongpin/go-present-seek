package main

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	Author string `json:"author"`
	Post   string `json:"post"`
}

func main() {
	http.HandleFunc("/books", getHandler)
	http.ListenAndServe(":8080", nil)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var book Book
		book = Book{
			Author: "John Doe",
			Post:   "This is a new post",
		}
		json.NewEncoder(w).Encode(&book)
	}
}
