package server




// Server object
type Server struct {
  Name string
  Connections int
}


// Creates a server
func createServer(name string) Server* {
  return Server {
    Name: name,
    Connections: 0,
  }
}
