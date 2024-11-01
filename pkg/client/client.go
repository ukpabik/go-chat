package client


import (
  "net"
  "fmt"
  "log"
  "io"
  "bufio"
)

type Client struct {
  Socket net.Conn
  Name string
  Reader io.Reader
  Writer io.Writer
}

func CreateClient(name string) (*Client, error) {
  conn, err := net.Dial("tcp", "localhost:8080")
  if err != nil {
    log.Fatal("Unable to connect to server")
    return nil, fmt.Errorf("Unable to connect to server: %v", err)
  }
  return &Client {
    Socket: conn,
    Name: name,
    Reader: conn,
    Writer: conn,
  }, nil
}


func Listen(client *Client) {
  defer client.Socket.Close()

  reader := bufio.NewReader(client.Reader)
  buffer := make([]byte, 1024)

  for {
    message, err := reader.Read(buffer)
    if err != nil {
      if err == io.EOF {
        log.Println("server DIED")
        break
      }
      log.Println("ERROR READING FROM SERVER")
      break
    }
    fmt.Println(string(buffer[:message]))
  }
}

func (client *Client) SendMessage(message string) error {
  _, err := client.Writer.Write([]byte (message + "\n"))
  if err != nil {
    return fmt.Errorf("error sending message: %v", err)
  }
  return nil
}
