package server


import (
  "testing"
  "time"
  "github.com/ukpabik/go-chat/pkg/client"
  "flag"
  "strconv"
)

/*
* To run test, use this command in main directory: 
* go test -v ./pkg/server -args -maxConnections=?
* for the ?, input whatever max you want to test
*/

var maxConnections int

func TestMain(m *testing.M) {
    flag.IntVar(&maxConnections, "maxConnections", 1000, "maximum number of client connections to test")
    flag.Parse()


    m.Run()
}

func TestMaxConnections(t *testing.T){

    // Create and establish our server
    netServer := CreateServer("TestServer")
    go func() {
        if err := netServer.Start(); err != nil {
            t.Fatalf("Failed to start server: %v", err)
        } else {
            t.Log("Server started successfully on localhost:8080")
        }
    }()

    // Let the server start up before attempting to connect
    time.Sleep(100 * time.Millisecond)

    successes := 0
    failures := 0

    for i := 0; i < maxConnections; i++ {
      name := "Example " + strconv.Itoa(i)

      user, err := client.CreateClient(name)
      if err != nil {
        t.Logf("Client number %d failed to connect: %v", i, err)
        failures++
        continue
      }

      successes++
      go client.Listen(user)


      time.Sleep(10 * time.Millisecond)
      user.Socket.Close()
    }


    time.Sleep(500 * time.Millisecond)
    t.Logf("Test completed with %d successful connections and %d failed connections", successes, failures)

    if failures > 0 {
      t.Errorf("Expected all connections to succeed, but %d connections failed", failures)
    }


}
