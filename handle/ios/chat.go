package ios

import (
  "fmt"
  "AT_Message_Server/alephtau"
  "net/http"
  "encoding/json"
  "strconv"
  "AT_Message_Server/database/message_query"
)

// 不要ios，直接指定swift语言，暂时精简技术开发
type responseToSwift struct {
  Code int `json:"code"`
  Description string  `json:"description"`
  Result string `json:"result"`
}


func Handles() {
  handle := alephtau.ATHandle{}
  handle.RegisterHandle("/ios/new-message", newMessage)
  handle.RegisterHandle("/path2", path2)
  handle.RegisterHandle("/path3", path3)
  handle.RegisterHandle("/path4", nil)
}


func newMessage(w http.ResponseWriter, r *http.Request) {
  var from int
  var to int
  var message string
  
  
  //fmt.Println("Enter path1")
  //fmt.Println(r.URL.Path)
  //fmt.Println(r.Method)
  
  //fmt.Println(r.PostFormValue("from"))
  //fmt.Println(r.PostFormValue("to"))
  //fmt.Println(r.PostFormValue("message"))

  
  r.PostFormValue("")    //如果不用这行代码解析一个空值，下面一行的r.Form将会返回一个空数组
  values := r.Form
  //fmt.Println(values)
  from, _ = strconv.Atoi(values["from"][0])
  to, _ = strconv.Atoi(values["to"][0])
  message = values["message"][0]
  
  // 这里实现mysql数据插入操作！！！
  query.InsertANewMessageToMySQL(from, to, message)
  
  // 数据库操作完成后执行响应（流程没走完，当然不能提前给出结果，这也是事物之理）
  response := responseToSwift{Code: 756, Description: "OK", Result: ""}
  responseBytes, err := json.Marshal(response)
  if err != nil {
    w.WriteHeader(404)
    //w.Write([]byte("已经接收到了来自swift客户端的数据 ... "))
  }
  w.Write(responseBytes)
  fmt.Println(string(responseBytes))
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

