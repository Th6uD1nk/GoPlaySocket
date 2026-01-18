package main

import (
  "fmt"
  "net"
  "sync"
  "time"
)

type Client struct {
  addr      *net.UDPAddr
  lastSeen  time.Time
}

type Server struct {
  conn    *net.UDPConn
  clients map[string]*Client
  mu      sync.RWMutex
}

func NewServer(port int) (*Server, error) {
  addr := net.UDPAddr{
    Port: port,
    IP:   net.ParseIP("0.0.0.0"),
  }
    
  conn, err := net.ListenUDP("udp", &addr)
  if err != nil {
    return nil, err
  }
    
  return &Server{
    conn:    conn,
    clients: make(map[string]*Client),
  }, nil
}

func (server *Server) addOrUpdateClient(addr *net.UDPAddr) {
  server.mu.Lock()
  defer server.mu.Unlock()
  
  key := addr.String()
  if client, exists := server.clients[key]; exists {
    client.lastSeen = time.Now()
  } else {
    server.clients[key] = &Client{
      addr:     addr,
      lastSeen: time.Now(),
    }
    fmt.Printf("+ New client: %s\n", addr.String())
  }
}

func (server *Server) listClients() {
  server.mu.RLock()
  defer server.mu.RUnlock()
  
  fmt.Println("\n=== Clients list ===")
  if len(server.clients) == 0 {
    fmt.Println("No connected client")
  } else {
    for key, client := range server.clients {
      fmt.Printf("- %s (last activity: %s)\n", 
        key, 
        time.Since(client.lastSeen).Round(time.Second))
    }
  }
  fmt.Println("========================\n")
}

func (server *Server) cleanInactiveClients(timeout time.Duration) {
  server.mu.Lock()
  defer server.mu.Unlock()
    
  now := time.Now()
  for key, client := range server.clients {
    if now.Sub(client.lastSeen) > timeout {
      fmt.Printf("x Client timeout: %s\n", key)
      delete(server.clients, key)
    }
  }
}

func (server *Server) Start() {
  defer server.conn.Close()
  
  fmt.Printf("UDP server stated on port %d\n", server.conn.LocalAddr().(*net.UDPAddr).Port)
  
  // clean inactive clients
  go func() {
    ticker := time.NewTicker(10 * time.Second)
    defer ticker.Stop()
    for range ticker.C {
      server.cleanInactiveClients(30 * time.Second)
    }
  }()
  
  // display client list periodically
  go func() {
    ticker := time.NewTicker(15 * time.Second)
    defer ticker.Stop()
    for range ticker.C {
      server.listClients()
    }
  }()
  
  buffer := make([]byte, 1024)
  
  for {
    nByte, addr, err := server.conn.ReadFromUDP(buffer)
    if err != nil {
      fmt.Printf("Read error %v\n", err)
      continue
    }
    
    server.addOrUpdateClient(addr)
    
    message := string(buffer[:nByte])
    fmt.Printf("+ received from %s: %s\n", addr.String(), message)
    
    response := fmt.Sprintf("ACK: %s", message)
    _, err = server.conn.WriteToUDP([]byte(response), addr)
    if err != nil {
      fmt.Printf("Sent error: %v\n", err)
    }
  }
}

func main() {
  server, err := NewServer(8888)
  if err != nil {
    fmt.Printf("Error on create: %v\n", err)
    return
  }
  server.Start()
}
