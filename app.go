package main

import (
  "log"
  "net/http"
)

func main() {
  assets := http.StripPrefix("/", http.FileServer(http.Dir("static/")))
  http.Handle("/", assets)

  log.Println("Listening at port 4000")
  log.Fatal(http.ListenAndServe(":4000", nil))
}
