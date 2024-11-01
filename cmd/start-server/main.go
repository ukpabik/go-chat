package main
import (
  "github.com/ukpabik/go-chat/pkg/server"
)


func main() {
  ns := server.CreateServer("Example Server")
  ns.Start()
}
