package main

import (
    "fmt"
    "net"
    "time"
)

func main() {
  serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8888")
  if err != nil {
    fmt.Printf("eRROR: %v\n", err)
    return
  }
    
  conn, err := net.DialUDP("udp", nil, serverAddr)
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }
  defer conn.Close()
    
  // Send periodic message
  for i := 0; i < 5; i++ {
    message := fmt.Sprintf("Message %d", i+1)
    _, err = conn.Write([]byte(message))
    if err != nil {
      fmt.Printf("Error on sent: %v\n", err)
      continue
    }
    
    buffer := make([]byte, 1024)
    conn.SetReadDeadline(time.Now().Add(2 * time.Second))
    nByte, err := conn.Read(buffer)
    if err != nil {
      fmt.Printf("Error on receive: %v\n", err)
      continue
    }
    
    fmt.Printf("Response: %s\n", string(buffer[:nByte]))
    time.Sleep(3 * time.Second)
  }
}
