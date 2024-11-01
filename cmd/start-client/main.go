package main

import (
  "fmt"
  "github.com/ukpabik/go-chat/pkg/client"
  "bufio"
  "os"
  "log"
)

func main() {
  fmt.Println("What is your name?")
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  name := scanner.Text()
  if len(name) < 1 {
    fmt.Println("Name too short")
    return
  }

  user, err := client.CreateClient(name)
  if err != nil {
    fmt.Println("Couldn't create client")
    return
  }

  log.Println("Created client successfully")

  go client.Listen(user)

  for scanner.Scan() {
    message := scanner.Text()
    if len(message) == 0 {
      continue
    }
    err := user.SendMessage(message)
    if err != nil {
      log.Println("Unable to send message")
      continue
    }
  }
}
