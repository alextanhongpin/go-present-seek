package main

import (
  "encoding/json"
  "flag"
  "fmt"
  "log"
  "os"
)

// Config schema for our app
type Config struct {
  // DSN is the data source name for our db connection
  DSN string `json:"dsn"`
}

func main() {
  var cfg = flag.String("config", "config.json", "Path to a config file.")
  flag.Parse()
  config := setupConfig(*cfg)
  fmt.Println(config.DSN)
}

// setupConfig takes a pointer reference and set the config loaded
func setupConfig(path string) Config {
  var config Config

  file, err := os.Open(path)
  defer file.Close()
  if err != nil {
    log.Fatal(err)
  }

  err = json.NewDecoder(file).Decode(&config)
  if err != nil {
    log.Fatal(err)
  }

  return config
}
