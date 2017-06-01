
// Not shown...

func getJobsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  rows, err := env.DB.Query("SELECT id, name, created_at FROM job")
  defer rows.Close()

  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  var jobs []Job
  // Iterate through each rows
  for rows.Next() {
    var job Job
    err := rows.Scan(&job.ID, &job.Name, &job.CreatedAt)
    if err != nil {
      http.Error(w, err.Error(), http.StatusBadRequest)
      return
    }
    jobs = append(jobs, job)
  }

  if err = rows.Err(); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(jobs)
}
