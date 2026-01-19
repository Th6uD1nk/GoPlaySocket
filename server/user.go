package main

import (
  "net"
  "time"
)

type UserType string

const (
  UserTypePlayer UserType = "player"
  UserTypeBot    UserType = "bot"
  UserTypeAdmin  UserType = "admin"
)

type User struct {
  id               string
  userType         UserType
  location         Vector3
  previousLocation Vector3
  orientation      float32
  isActive         bool
  lastUpdate       time.Time
  conn             *net.UDPConn
}

func NewUser(id string, userType UserType, conn *net.UDPConn) *User {
  return &User{
    id:          id,
    userType:    userType,
    location:    Vector3{x: 0, y: 0, z: 0},
    orientation: 0,
    isActive:    true,
    lastUpdate:  time.Now(),
    conn:        conn,
  }
}

func (user *User) calculateLastDirection() Vector3 {
  return Vector3{
    x: user.location.x - user.previousLocation.x,
    y: user.location.y - user.previousLocation.y,
    z: user.location.z - user.previousLocation.z,
  }
}

func (user *User) updatePosition(newLocation Vector3) {
  user.previousLocation = user.location
  user.location = newLocation
  user.lastUpdate = time.Now()
}
