package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Port         int      `json:"port"`
	AppName      string   `json:"app_name"`
	Toggle       bool     `json:"toggle"`
	WhiteListUrl []string `json:"whitelist_url"`
}

func main() {

	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}

	// Read the json data and convert it to golang type
	var cfg Config
	err = json.NewDecoder(file).Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%#v", cfg)

}
