//go:build mobile
package rgl

import (
  "unsafe"
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

func Str(s string) string {
  return s
}

func GetUniformLocation(program uint32, name string) int32 {
  uniform := glctx.GetUniformLocation(gl.Program{Value: program}, name)
  return int32(uniform.Value)
}

func GetAttribLocation(program uint32, name string) int32 {
  attrib := glctx.GetAttribLocation(gl.Program{Value: program}, name)
  return int32(attrib.Value)
}

func CreateProgram() uint32 {
  program := glctx.CreateProgram()
  return uint32(program.Value)
}

func AttachShader(program uint32, shader uint32) {
  glctx.AttachShader(
    gl.Program{Value: program}, 
    gl.Shader{Value: shader},
  )
}

func LinkProgram(program uint32) {
  glctx.LinkProgram(gl.Program{Value: program})
}

func CreateShader(shaderType uint32) uint32 {
  shader := glctx.CreateShader(gl.Enum(shaderType))
  return uint32(shader.Value)
}

func Strs(sources ...string) (**uint8, func()) {
  lastShaderSource = sources[0]
  return nil, func() {}
}

var lastShaderSource string

func ShaderSource(shader uint32, count int32, source **uint8, length *int32) {
  glctx.ShaderSource(gl.Shader{Value: shader}, lastShaderSource)
}

func CompileShader(shader uint32) {
  glctx.CompileShader(gl.Shader{Value: shader})
}

func DeleteShader(shader uint32) {
  glctx.DeleteShader(gl.Shader{Value: shader})
}

func GetProgramiv(program uint32, pname uint32, params *int32) {
  *params = int32(glctx.GetProgrami(gl.Program{Value: program}, gl.Enum(pname)))
}

func GetProgramInfoLog(program uint32, maxLength int32, length *int32, infoLog *byte) string {
  return glctx.GetProgramInfoLog(gl.Program{Value: program})
}

func GetShaderiv(shader uint32, pname uint32, params *int32) {
  *params = int32(glctx.GetShaderi(gl.Shader{Value: shader}, gl.Enum(pname)))
}

func GetShaderInfoLog(shader uint32, maxLength int32, length *int32, infoLog *byte) string {
  return glctx.GetShaderInfoLog(gl.Shader{Value: shader})
}

func UseProgram(program uint32) {
  glctx.UseProgram(gl.Program{Value: program})
}

func DeleteProgram(program uint32) {
  glctx.DeleteProgram(gl.Program{Value: program})
}

func UniformMatrix4fv(location int32, count int32, transpose bool, value *float32) {
  data := unsafe.Slice(value, 16*count)
  glctx.UniformMatrix4fv(gl.Uniform{Value: location}, data)
}

func Uniform4f(location int32, v0, v1, v2, v3 float32) {
  glctx.Uniform4f(gl.Uniform{Value: location}, v0, v1, v2, v3)
}

func GenBuffers(n int32, buffers *uint32) {
  bufs := unsafe.Slice(buffers, n)
  for i := range bufs {
    buf := glctx.CreateBuffer()
    bufs[i] = uint32(buf.Value)
  }
}

func BindBuffer(target uint32, buffer uint32) {
  glctx.BindBuffer(gl.Enum(target), gl.Buffer{Value: buffer})
}

func BufferData(target uint32, size int, data unsafe.Pointer, usage uint32) {

  var byteData []byte
  if data != nil {
    byteData = unsafe.Slice((*byte)(data), size)
  } else {
    byteData = make([]byte, size)
  }
  glctx.BufferData(gl.Enum(target), byteData, gl.Enum(usage))
}

func Ptr(data interface{}) unsafe.Pointer {

  switch v := data.(type) {
  case []float32:
    if len(v) > 0 {
      return unsafe.Pointer(&v[0])
    }
  case []uint16:
    if len(v) > 0 {
      return unsafe.Pointer(&v[0])
    }
  case []byte:
    if len(v) > 0 {
      return unsafe.Pointer(&v[0])
    }
  }
  return nil
}

func EnableVertexAttribArray(index uint32) {
  glctx.EnableVertexAttribArray(gl.Attrib{Value: uint(index)})
}

func VertexAttribPointer(index uint32, size int32, xtype uint32, normalized bool, stride int32, pointer unsafe.Pointer) {
  offset := int(uintptr(pointer))
  glctx.VertexAttribPointer(gl.Attrib{Value: uint(index)}, int(size), gl.Enum(xtype), normalized, int(stride), offset)
}

func DrawArrays(mode uint32, first int32, count int32) {
  glctx.DrawArrays(gl.Enum(mode), int(first), int(count))
}

func DisableVertexAttribArray(index uint32) {
  glctx.DisableVertexAttribArray(gl.Attrib{Value: uint(index)})
}

func DeleteBuffers(n int32, buffers *uint32) {
  bufs := unsafe.Slice(buffers, n)
  for _, buf := range bufs {
    glctx.DeleteBuffer(gl.Buffer{Value: buf})
  }
}

func Viewport(x, y, width, height int32) {
  glctx.Viewport(int(x), int(y), int(width), int(height))
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
