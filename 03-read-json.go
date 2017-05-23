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
	http.HandleFunc("/books", getHandler)
	http.ListenAndServe(":8080", nil)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var book Book
		json.NewDecoder(r.Body).Decode(&book)

		fmt.Println(book)

		json.NewEncoder(w).Encode(book)
	}
}
