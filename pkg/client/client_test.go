package client

import (
    "testing"
    "time"
    "github.com/ukpabik/go-chat/pkg/server"
)

func TestListen(t *testing.T) {

    // Create and establish our server
    netServer := server.CreateServer("TestServer")
    go func() {
        if err := netServer.Start(); err != nil {
            t.Fatalf("Failed to start server: %v", err)
        } else {
            t.Log("Server started successfully on localhost:8080")
        }
    }()

    // Let the server start up before attempting to connect
    time.Sleep(100 * time.Millisecond)

    // Initialize our testing client
    user, err := CreateClient("Bowser") 
    if err != nil {
      t.Fatal("FAILED TEST")
      return
    }
    
    t.Logf("Client created successfully: %s", user.Name)
    

    // Running Listen in a goroutine to listen while sending messages
    go Listen(user)

    testMessage := "Hello World!\n"

    // Write test message to the server
    _, err = user.Socket.Write([]byte(testMessage))
    if err != nil {
        t.Fatalf("failed to write to server: %v", err)
    }
    t.Logf("Sending message to the server: %q", testMessage)

    // Create a buffer to store the message
    buff := make([]byte, 1024)
    n, err := user.Socket.Read(buff)
    if err != nil {
       t.Fatalf("Failed to read message %v", err)
    }

    // Check if our expected response is the same
    expectedResponse := "Server received message"
    receivedResponse := string(buff[:n])
    if receivedResponse != expectedResponse {
        t.Errorf("Messages do not match:\nExpected: %q\nGot: %q", expectedResponse, receivedResponse)
    }
    t.Logf("Expected response: %q", expectedResponse)
    t.Logf("Received response: %q", receivedResponse)

    // Close the client connection
    user.Socket.Close()
}
