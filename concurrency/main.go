package main

import (
	// "fmt"
	"log"
	"time"
)

func main() {
	defer timeTrack(time.Now(), "delay test")

	c := make(chan bool, 2)
	go func() {
		// Delay by one second
		time.Sleep(time.Second)
		c <- true
	}()

	go func() {
		// Delay by two seconds
		time.Sleep(2 * time.Second)
		c <- true
	}()

	<-c
	<-c
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
