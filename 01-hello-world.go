package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

// index route is pointing to /
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `{"message": "hello world"}`)
}
