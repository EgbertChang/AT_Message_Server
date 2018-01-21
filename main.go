package main

import (
  "fmt"
  "net/http"
  "log"
  "time"
)


func main() {
  fmt.Println("Alephtau Server is starting")
  
  server := &http.Server{
    Addr: ":9090",
    Handler: &AlephTauServer{},    // type Handler interface {}，handler to invoke, http.DefaultServeMux if nil
    ReadTimeout: 3 * time.Minute,
    WriteTimeout: 3 * time.Minute,
    MaxHeaderBytes: 1 << 20,    // 这里依然使用右移的方式来表示存储空间
  }
  log.Fatal(server.ListenAndServe())
}
