package main

import (
	"fmt"
)

func addition(x, y int) int {
	return x + y
}

func main() {
	fmt.Printf("1 + 2 = %d\n", addition(1, 2))
	// This will throw an error
	// addition(false, true)
}
