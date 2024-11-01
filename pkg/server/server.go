package server



import (
  "net"
  "log"
  "fmt"
  "io"
)


type NetServer struct {
  Connections []net.Conn
  Name string
  Reader io.Reader
  Writer io.Writer
}

// Creates a server with name and empty list
func CreateServer(name string) *NetServer {
  return &NetServer {
    Name: name,
  }
}

// Starts our server
func (s *NetServer) Start() error {
  net, err := net.Listen("tcp", ":8080")
  if err != nil{
    log.Printf("unable to start server: %v", err)
    return err
  }
  defer net.Close()

  for {
    conn, err := net.Accept()
    if err != nil {
      log.Printf("unable to accept connection: %v", err)
      continue
    }

    s.Connections = append(s.Connections, conn)
    fmt.Println("New connection added:", conn.RemoteAddr())


    // TODO: Handle connection
    go s.receiveMessage(conn)
  }
}


// Loops and reads all messages from client
func (s *NetServer) receiveMessage(conn net.Conn){
  defer conn.Close()

  for {
    // Make read buffer and read all bytes sent from the connection
    buff := make([]byte, 1024)
    n, err := conn.Read(buff)

    if err != nil {
      log.Println("Error reading bytes: ", err.Error())
      return
    }

    fmt.Println("Received data: ", string(buff[:n]))

    // Send response to user
    response := []byte("Server received message")

    // Send response to client
    _, err = conn.Write(response)
    if err != nil {
      log.Println("Error writing to client: ", err)
      return
    }
  }
}
