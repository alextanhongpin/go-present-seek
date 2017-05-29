package main

import (
	// "encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("listening to port *:8080. press ctrl + c to cancel")

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := make(chan int, 2)

	go func() {
		c <- fib(10)
	}()

	go func() {
		c <- fib(20)
	}()

	fmt.Fprintf(w, `{"fib10": %d, "fib20": %d}`, <-c, <-c)
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
