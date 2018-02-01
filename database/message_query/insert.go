package query

import (
  "database/sql"
  "log"
)


func InsertANewMessageToMySQL(from int, to int, message string) {
  println(from)
  println(to)
  println(message)
  
  
  db, err := sql.Open("mysql", "root:egbert@tcp(121.43.183.17:3306)/AT_Message_Dev")
  if err != nil {
    log.Println(err)
  }
  defer func() {
    // 函数结束后db就被关闭了，底层应该是关闭这个socket。
    // 一个goroutine只用一次，所以这里没有什么线程池的概念，用一次就建立一次连接。
    // 所以后面需要优化，便于复用conn
    db.Close()
  }()
  
  stmt,_ := db.Prepare("INSERT messages (sender, receiver, message) values (?,?,?)")
  _, err = stmt.Exec(from, to, message)
  if err != nil {
    log.Println(err)
  }
  
}
