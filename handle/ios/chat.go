package ios

import (
  "fmt"
  "AT_Message_Server/alephtau"
  "net/http"
)

func Handles() {
  handle := alephtau.ATHandle{}
  handle.RegisterHandle("/ios/new-message", newMessage)
  handle.RegisterHandle("/path2", path2)
  handle.RegisterHandle("/path3", path3)
  handle.RegisterHandle("/path4", nil)
}


func newMessage(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Enter path1")
  fmt.Println(r.URL.Path)
  
  w.Write([]byte("<h2>This is a path1 ...</h2>"))
}


func path2(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Enter path2")
  fmt.Println(r.URL.Path)
  
  // 加上<body>浏览器可以解析成h1样式
  w.Write([]byte("<body><h2>This is a path2 ...</h2></body>"))
}


func path3(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Enter path3")
  fmt.Println(r.URL.Path)
  
  w.Write([]byte("<h2>This is a path3 ...</h2>"))
}

