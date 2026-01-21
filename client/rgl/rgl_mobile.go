//go:build mobile
package rgl

import (
  gl "golang.org/x/mobile/gl"
)

var glctx gl.Context

func Init(ctx gl.Context) error {
  glctx = ctx
  return nil
}

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

func GetUniformLocation(program gl.Program, name string) gl.Uniform {
  return glctx.GetUniformLocation(program, name)
}

func GetAttribLocation(program gl.Program, name string) gl.Attrib {
  return glctx.GetAttribLocation(program, name)
}

func CreateProgram() gl.Program {
  return glctx.CreateProgram()
}

func AttachShader(program gl.Program, shader gl.Shader) {
  glctx.AttachShader(program, shader)
}

func LinkProgram(program gl.Program) {
  glctx.LinkProgram(program)
}

func GetProgramiv(program gl.Program, pname gl.Enum, params *int32) {
  *params = int32(glctx.GetProgrami(program, pname))
}

func GetProgramInfoLog(program gl.Program) string {
  return glctx.GetProgramInfoLog(program)
}

func CreateShader(typ gl.Enum) gl.Shader {
  return glctx.CreateShader(typ)
}

func ShaderSource(shader gl.Shader, source string) {
  glctx.ShaderSource(shader, source)
}

func CompileShader(shader gl.Shader) {
  glctx.CompileShader(shader)
}

func GetShaderiv(shader gl.Shader, pname gl.Enum, params *int32) {
  *params = int32(glctx.GetShaderi(shader, pname))
}

func GetShaderInfoLog(shader gl.Shader) string {
  return glctx.GetShaderInfoLog(shader)
}

func UseProgram(program gl.Program) {
  glctx.UseProgram(program)
}

func UniformMatrix4fv(location gl.Uniform, data []float32) {
  glctx.UniformMatrix4fv(location, data)
}

func Uniform4f(location gl.Uniform, v0, v1, v2, v3 float32) {
  glctx.Uniform4f(location, v0, v1, v2, v3)
}

func GenBuffers(n int) []gl.Buffer {
  bufs := make([]gl.Buffer, n)
  for i := range bufs {
    bufs[i] = glctx.CreateBuffer()
  }
  return bufs
}

func BindBuffer(target gl.Enum, buffer gl.Buffer) {
  glctx.BindBuffer(target, buffer)
}

func BufferData(target gl.Enum, data []byte, usage gl.Enum) {
  glctx.BufferData(target, data, usage)
}

func EnableVertexAttribArray(index gl.Attrib) {
  glctx.EnableVertexAttribArray(index)
}

func VertexAttribPointer(index gl.Attrib, size int, xtype gl.Enum, normalized bool, stride int, offset int) {
  glctx.VertexAttribPointer(index, size, xtype, normalized, stride, offset)
}

func DrawArrays(mode gl.Enum, first, count int) {
  glctx.DrawArrays(mode, first, count)
}

func DisableVertexAttribArray(index gl.Attrib) {
  glctx.DisableVertexAttribArray(index)
}

func DeleteBuffers(buffers []gl.Buffer) {
  for _, buf := range buffers {
    glctx.DeleteBuffer(buf)
  }
}

func Viewport(x, y, width, height int) {
  glctx.Viewport(x, y, width, height)
}

func ClearColor(r, g, b, a float32) {
  glctx.ClearColor(r, g, b, a)
}

func Clear(mask gl.Enum) {
  glctx.Clear(mask)
}

func Enable(cap gl.Enum) {
  glctx.Enable(cap)
}

func Disable(cap gl.Enum) {
  glctx.Disable(cap)
}

func BlendFunc(sfactor, dfactor gl.Enum) {
  glctx.BlendFunc(sfactor, dfactor)
}
