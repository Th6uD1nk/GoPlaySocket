//go:build mobile

package main

import (
  "log"

  "golang.org/x/mobile/app"
  "golang.org/x/mobile/event/lifecycle"
  "golang.org/x/mobile/event/paint"
  "golang.org/x/mobile/event/size"
  "golang.org/x/mobile/gl"
)

func main() {
  worldState := NewWorldState()

  // Start UDP client
  client, err := NewUDPClient("127.0.0.1:8888", worldState)
  if err != nil {
    log.Fatalf("Cannot create UDP client: %v", err)
  }
  defer client.Conn.Close()

  client.StartReceiving()
  client.StartSending()

  game := NewGame(worldState)

  app.Main(func(a app.App) {
    var glctx gl.Context
    var sz size.Event

    for e := range a.Events() {
      switch e := e.(type) {

      case lifecycle.Event:
        if e.To == lifecycle.StageDead {
          return
        }

        if e.From < lifecycle.StageAlive && e.To >= lifecycle.StageAlive && glctx == nil {
          var ok bool
          glctx, ok = e.DrawContext.(gl.Context)
          if !ok {
            log.Fatal("unable to get GL context")
          }
        }

      case size.Event:
        sz = e

      case paint.Event:
        if glctx == nil {
          continue
        }

        game.Draw(int(sz.WidthPx), int(sz.HeightPx))

        a.Send(paint.Event{})
      }
    }
  })
}
