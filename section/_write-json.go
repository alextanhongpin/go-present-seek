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
	http.HandleFunc("/books", bookHandler)
	fmt.Println("listening to port *:8080. press ctrl + c to cancel")
	http.ListenAndServe(":8080", nil)
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var book Book
		book = Book{
			Author: "John Doe",
			Post:   "This is a new post",
		}
		json.NewEncoder(w).Encode(&book)
	} else if r.Method == "POST" {
		var book Book
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(book)
		fmt.Fprint(w, `{"message": "successfully decode json"}`)
	} else {
		http.Error(w, "Invalid request method", 405)
	}
}
