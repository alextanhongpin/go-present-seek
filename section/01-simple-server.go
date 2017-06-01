package main

import (
  "fmt"
  "net/http"

  "github.com/julienschmidt/httprouter"
)

func main() {
  port := 8080
  router := httprouter.New()

  router.GET("/api/jobs", getJobsHandler)

  fmt.Printf("listening to port *:%d. press ctrl + c to cancel", port)
  http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}

// Get a list of job
func getJobsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  w.Write([]byte(`{"message": "hello world!"}`))
}
