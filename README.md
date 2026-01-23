# A socket-based game server and client written in Go

This is an experimental OpenGL/OpenGL ES game client and server written in Go.  
The project uses UDP for real-time gameplay and TCP for zone management and critical state updates.

The client is available on both PC and mobile platforms using Go Mobile.

Please refer to the individual README files in the `client` and `server` directories for more detailed information.


## roadmap

### already existing

- simple udp socket client
  - basic 3d rendering using opengl
  - opengl wrappers for pc and mobile using gomobile
  - basic android manifest
  - basic external configuration
  - pc and mobile main entry points

- simple udp socket server
  - spawn, update, cleanup
  - basic world state registry

### overall next steps

- camera (WIP)
- world generation
- player control and world state updates
- gameplay
