package main

import (
	// "fmt"
	"log"
	"time"
)

func delay100() bool {
	time.Sleep(1 * time.Second)
	return true
}

func delay200() bool {
	time.Sleep(2 * time.Second)
	return true
}

func main() {
	defer timeTrack(time.Now(), "delay test")
	c := make(chan bool, 2)
	go func() {
		c <- delay100()
	}()
	go func() {
		c <- delay200()
	}()

	<-c
	<-c
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
