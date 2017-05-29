package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Author string `json:"author"`
	Post   string `json:"post"`
}

func main() {

	// Decode json data
	b := Book{
		Author: "Someone",
		Post:   "Something",
	}

	jsonOutput, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jsonOutput)

	// Encode json data
	jsonInput := []byte(`{"author": "john", "post": "doe"}`)

	var book Book
	err = json.Unmarshal(jsonInput, &book)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(book)

}
