package alephtau

import (
  "net/http"
  "sync"
)

var handles = make(map[string]http.HandlerFunc)

type ATHandle struct {
  mu sync.RWMutex
}

func (handle *ATHandle) RegisterHandle(pattern string, handleFunc http.HandlerFunc) {
  handle.mu.Lock()
  defer handle.mu.Unlock()
  handles[pattern] = handleFunc
}