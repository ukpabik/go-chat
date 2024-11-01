package server



import (
  "net"
  "log"
  "fmt"
  "io"
  "strings"


  "github.com/fatih/color"
)

type Client struct {
  Conn net.Conn
  Name string
}


type NetServer struct {
  Connections []Client
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

// Starts our server and listens for incoming client connections
func (s *NetServer) Start() error {
  listener, err := net.Listen("tcp", "localhost:8080")
  if err != nil{
    log.Printf("unable to start server: %v", err)
    return err
  }
  defer listener.Close()

  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Printf("unable to accept connection: %v", err)
      continue
    }

    // Creating buffer for receiving the name
    nameBuf := make([]byte, 1024)
    n, err := conn.Read(nameBuf)
    if err != nil {
      log.Printf("Can't read client name: %v", err)
      conn.Close()
      continue
    }

    clientName := strings.TrimSpace(string(nameBuf[:n]))
    client := Client {
      Conn: conn,
      Name: clientName,
    }

    s.Connections = append(s.Connections, client)
    fmt.Printf("New connection added: %s --> %s\n", client.Name, conn.RemoteAddr())
    fmt.Printf("There are now %d users connected!\n", len(s.Connections))


    // TODO: Handle connection
    go s.receiveMessage(client)
  }
}


// Loops and reads all messages from client
func (s *NetServer) receiveMessage(client Client){
  defer func() {
    client.Conn.Close()
    s.removeConnection(client)
    log.Println("Connection closed by client:", client.Conn.RemoteAddr())
  }()

  for {
    // Make read buffer and read all bytes sent from the connection
    buff := make([]byte, 1024)
    n, err := client.Conn.Read(buff)

    if err != nil {
      log.Println("Error reading bytes: ", err.Error())
      return
    }

    if n > 0 {
      message := strings.TrimSpace(string(buff[:n]))

      // Run the sendGlobal in a goroutine so it doesnt block any other actions
      go s.sendGlobalMessage(client, message)
    }
  }
}

func (s *NetServer) removeConnection(client Client) {
  for i, activeConnection := range s.Connections {
    if activeConnection == client {
      s.Connections = append(s.Connections[:i], s.Connections[i+1:]...)
      break
    }
  }
}

func (s *NetServer) sendGlobalMessage(sender Client, message string) {
  for _, conn := range s.Connections {
    if conn == sender {
      continue
    }
    convertedMessage := []byte(colorName(sender.Name) + ": " + message + "\n")
    _, err := conn.Conn.Write(convertedMessage)
    if err != nil {
      log.Println("Error writing to client:", err)
      return
    }
  }
  fmt.Printf("Global message from %s: %s\n", colorName(sender.Name), message)
}

func colorName(name string) string{
  blue := color.New(color.FgBlue).SprintFunc()
  return fmt.Sprintf(blue("[%s]"), name)
}
