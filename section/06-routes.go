
func main() {
  // ...not shown

  // Setup router
  router := httprouter.New()

  // Setup routes
  router.GET("/api/jobs", getJobsHandler)
  router.GET("/api/jobs/:id", getJobHandler)
  router.POST("/api/jobs", createJobHandler)
  router.DELETE("/api/jobs/:id", deleteJobHandler)
  router.PUT("/api/jobs/:id", updateJobHandler)

  // Start server
  fmt.Printf("listening to port *:%d. press ctrl + c to cancel", *port)
  http.ListenAndServe(fmt.Sprintf(":%d", *port), router)
}