package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	fmt.Println("listening to port *:8080. press ctrl + c to cancel")
	http.ListenAndServe(":8080", nil)
}

// index route is pointing to /
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `{"message": "hello world"}`)
}
