package alephtau

import (
  "net/http"
)

type Server struct {}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  mux := multiplexer()
  mux.ServeHTTP(w, r)
}
