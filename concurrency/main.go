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
		time.Sleep(time.Second)
		c <- true
	}()

	go func() {
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
