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

// CreateClient establishes a connection to the server and sends the client's name
func CreateClient(name string) (*Client, error) {
  conn, err := net.Dial("tcp", "localhost:8080")
  if err != nil {
    log.Fatal("Unable to connect to server")
    return nil, fmt.Errorf("Unable to connect to server: %v", err)
  }
  client := &Client {
    Socket: conn,
    Name: name,
    Reader: conn,
    Writer: conn,
  }

  _, err = conn.Write([]byte(name + "\n"))
  if err != nil {
    return nil, fmt.Errorf("Unable to send name to server: %v", err)
  }

  return client, nil
}

// Listen reads incoming messages from the server and displays them
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
    fmt.Print(string(buffer[:message]))
  }

}

// SendMessage sends a message to the server from the client
func (client *Client) SendMessage(message string) error {
  _, err := client.Writer.Write([]byte (message + "\n"))
  if err != nil {
    return fmt.Errorf("error sending message: %v", err)
  }
  return nil
}

