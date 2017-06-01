package main

import (
	"flag"
	"fmt"
)

func main() {
	var port = flag.Int("port", 8080, "The server port")
	var env = flag.String("env", "dev", "The server environment")
	flag.Parse()
	fmt.Println(*port)
	fmt.Println(*env)
}
