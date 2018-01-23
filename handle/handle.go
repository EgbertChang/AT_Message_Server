package handle

import (
  "AT_Message_Server/alephtau"
  "net/http"
  "fmt"
)

func RegisterHandles() {
  handle := alephtau.ATHandle{}
  handle.RegisterHandle("/path1", path1)
  handle.RegisterHandle("/path2", path2)
  handle.RegisterHandle("/path3", path3)
  handle.RegisterHandle("/path4", nil)
}


func path1(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Enter path1")
  fmt.Println(r.URL.Path)
  
  w.Write([]byte("<h2>This is a path1 ...</h2>"))
}

func path2(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Enter path2")
  fmt.Println(r.URL.Path)
  
  w.Write([]byte("<body><h2>This is a path2 ...</h2></body>"))
}

func path3(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Enter path3")
  fmt.Println(r.URL.Path)
  
  w.Write([]byte("<h2>This is a path3 ...</h2>"))
}
