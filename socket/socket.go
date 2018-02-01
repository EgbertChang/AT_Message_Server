package socket

import (
  "net"
  "log"
  "os"
  "fmt"
)

func SendImageToSwiftClient() {
  listen, err := net.Listen("tcp", "0.0.0.0:9090")
  if err != nil {
    log.Fatal(err)
  }
  defer func(){
    listen.Close()
  }()
  
  for {
    conn, err := listen.Accept()
    if err != nil {
      log.Fatal(err)
    }
    go newSocketConnection(conn)
  }

}

func newSocketConnection(conn net.Conn) {
  
  file, err := os.OpenFile("Vainglory_Halcyon_Fold_map.tif", os.O_RDWR, 0777)
  if err != nil {
    log.Println("Open file faild!")
  }

  fileInfo, err := file.Stat()
  fmt.Println(fileInfo.Size())
  fmt.Println(fileInfo.Name())

  // TODO:
}


