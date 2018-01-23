package alephtau

import (
  "net/http"
)

func multiplexer() *http.ServeMux {
  mux := http.NewServeMux()
  for k, v := range handles {
    mux.HandleFunc(k, v)
  }
  return mux
}
