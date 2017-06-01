package main

import (
  "database/sql"
  "fmt"
  "github.com/go-sql-driver/mysql"
)

// Config schema for our app
type Config struct {
  // DSN is the data source name for our db connection
  DSN string `json:"dsn"`
}

func main() {
  // load config...
  db := setupDB(config.DSN)
  defer db.Close()
}

// setupDB prepares the db for connections
func setupDB(dataSourceName string) *sql.DB {
  var db *sql.DB

  db, err := sql.Open("mysql", dataSourceName)
  if err != nil {
    log.Fatal(err)
  }
  // Limit to 150 connections, default is unlimited
  db.SetMaxOpenConns(150)

  // Ping to ensure connection is available
  err = db.Ping()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("connection to database open")
  return db
}
