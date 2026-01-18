package main

import (
  "fmt"
  "net"
  "time"
)

func main() {
  serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8888")
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }
  
  conn, err := net.DialUDP("udp", nil, serverAddr)
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }
  defer conn.Close()
  
  // receive ACK continuously
  go func() {
    buffer := make([]byte, 1024)
    for {
      n, err := conn.Read(buffer)
      if err != nil {
        fmt.Printf("Receive error: %v\n", err)
        continue
      }
      fmt.Printf("+ ACK received: %s\n", string(buffer[:n]))
    }
  }()
  
  // send messages continuously
  go func() {
    counter := 1
    ticker := time.NewTicker(2 * time.Second)
    defer ticker.Stop()
    
    for range ticker.C {
      message := fmt.Sprintf("Message %d", counter)
      _, err := conn.Write([]byte(message))
      if err != nil {
        fmt.Printf("Send error: %v\n", err)
        continue
      }
      fmt.Printf("- Sent: %s\n", message)
      counter++
    }
  }()
  
  select {}
}
