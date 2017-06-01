package main

// START OMIT
import (
  "flag"
  "fmt"
  "net/http"

  "github.com/julienschmidt/httprouter"
)

func main() {
  var port = flag.Int("port", 8080, "The server port")
  flag.Parse()
  router := httprouter.New()

  router.GET("/api/jobs", getJobsHandler)

  fmt.Printf("listening to port *:%d. press ctrl + c to cancel", *port)
  http.ListenAndServe(fmt.Sprintf(":%d", *port), router)
}

// ... go run main.go -port 4000

// END OMIT

// Get a list of job
func getJobsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  w.Write([]byte(`{"message": "hello world!"}`))
}
