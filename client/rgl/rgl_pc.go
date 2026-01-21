//go:build !mobile
package rgl

import gl "github.com/go-gl/gl/v2.1/gl"

const (
  FALSE                  = gl.FALSE
  TRUE                   = gl.TRUE
  INFO_LOG_LENGTH        = gl.INFO_LOG_LENGTH
  COMPILE_STATUS         = gl.COMPILE_STATUS
  LINK_STATUS            = gl.LINK_STATUS
  VERTEX_SHADER          = gl.VERTEX_SHADER
  FRAGMENT_SHADER        = gl.FRAGMENT_SHADER
  ARRAY_BUFFER           = gl.ARRAY_BUFFER
  STATIC_DRAW            = gl.STATIC_DRAW
  TRIANGLES              = gl.TRIANGLES
  LINES                  = gl.LINES
  FLOAT                  = gl.FLOAT
  COLOR_BUFFER_BIT       = gl.COLOR_BUFFER_BIT
  DEPTH_BUFFER_BIT       = gl.DEPTH_BUFFER_BIT
  DEPTH_TEST             = gl.DEPTH_TEST
  BLEND                  = gl.BLEND
  SRC_ALPHA              = gl.SRC_ALPHA
  ONE_MINUS_SRC_ALPHA    = gl.ONE_MINUS_SRC_ALPHA
)

var (
  Init                     = gl.Init
  GetUniformLocation       = gl.GetUniformLocation
  GetAttribLocation        = gl.GetAttribLocation
  CreateProgram            = gl.CreateProgram
  AttachShader             = gl.AttachShader
  LinkProgram              = gl.LinkProgram
  GetProgramiv             = gl.GetProgramiv
  GetProgramInfoLog        = gl.GetProgramInfoLog
  CreateShader             = gl.CreateShader
  Strs                     = gl.Strs
  Str                      = gl.Str
  ShaderSource             = gl.ShaderSource
  CompileShader            = gl.CompileShader
  GetShaderiv              = gl.GetShaderiv
  GetShaderInfoLog         = gl.GetShaderInfoLog
  UseProgram               = gl.UseProgram
  UniformMatrix4fv         = gl.UniformMatrix4fv
  Uniform4f                = gl.Uniform4f
  GenBuffers               = gl.GenBuffers
  BindBuffer               = gl.BindBuffer
  BufferData               = gl.BufferData
  EnableVertexAttribArray  = gl.EnableVertexAttribArray
  VertexAttribPointer      = gl.VertexAttribPointer
  DrawArrays               = gl.DrawArrays
  DisableVertexAttribArray = gl.DisableVertexAttribArray
  DeleteBuffers            = gl.DeleteBuffers
  Viewport                 = gl.Viewport
  ClearColor               = gl.ClearColor
  Clear                    = gl.Clear
  Enable                   = gl.Enable
  Disable                  = gl.Disable
  BlendFunc                = gl.BlendFunc
  Ptr                      = gl.Ptr
)
