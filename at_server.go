package main

import (
  "net/http"
  "fmt"
)

type AlephTauServer struct {}

func (s *AlephTauServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Println(r.URL.Path)
}
