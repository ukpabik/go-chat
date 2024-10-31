package main


import (
  "fmt"
  "github.com/ukpabik/go-chat/pkg/server"
)


func main() {
  g := server.createServer("Example")
  fmt.Println(g.Name)
}
