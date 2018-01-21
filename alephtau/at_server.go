package alephtau

import (
  "net/http"
  "fmt"
)

type Server struct {}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Println(r.URL.Path)

  mux := http.NewServeMux()
  mux.HandleFunc("/path1", path1)
  mux.HandleFunc("/path2", path2)
  mux.HandleFunc("/path3", path3)

  mux.ServeHTTP(w, r)
}

func path1(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Enter path1")
  fmt.Println(r.URL.Path)

  w.Write([]byte("This is a path3 ..."))
}

func path2(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Enter path2")
  fmt.Println(r.URL.Path)

  w.Write([]byte("This is a path3 ..."))
}

func path3(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Enter path3")
  fmt.Println(r.URL.Path)

  w.Write([]byte("This is a path3 ..."))
}

