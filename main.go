package main

import (
  "fmt"
  "net/http"
  "log"
  "time"
  "AT_Message_Server/alephtau"
  "AT_Message_Server/handle"
  "runtime"
  _ "github.com/go-sql-driver/mysql"
  "AT_Message_Server/socket"
)

func init() {
  runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
  fmt.Println("Alephtau Server is starting")
  
  go socket.SendImageToSwiftClient()
  
  handle.Manager()
  server := &http.Server{
    Addr: ":9091",
    Handler: &alephtau.Server{},    // type Handler interface {}，handler to invoke, http.DefaultServeMux if nil
    ReadTimeout: 3 * time.Minute,
    WriteTimeout: 3 * time.Minute,
    MaxHeaderBytes: 1 << 20,    // 这里依然使用右移的方式来表示存储空间
  }
  
  fmt.Println("Alephtau Server is running")
  log.Fatal(server.ListenAndServe())
}
