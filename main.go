package main

import (
	"fmt"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/kylemeenehan/go-opengl-play/cell"
	"github.com/kylemeenehan/go-opengl-play/snek"
	"log"
	"runtime"
	"strings"
	"time"
)



const (
	width  = 500
	height = 500
	rows = 10
	columns = 10

	vertexShaderSource = `
		#version 410
		in vec2 vp;
		void main() {
			gl_Position = vec4(vp, 0, 1.0);
		}
	` + "\x00"

	fragmentShaderSource = `
		#version 410
		out vec4 frag_colour;
		void main() {
			frag_colour = vec4(1, 1, 1, 1.0);
		}
	` + "\x00"
)

//var headCell = cell.Coordinates{ X: 0, Y: 0}
var gameSnek snek.Snek

func main() {
	runtime.LockOSThread()
	window := initGlfw()
	defer glfw.Terminate()
	program := initOpenGL()
	cell.Init(width, height, rows, columns)
	gameSnek = snek.NewSnek(0, 0)
	for !window.ShouldClose() {
		draw(window, program)
		time.Sleep(time.Second / 30)
	}
}
func draw(window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)
	gameSnek.Draw()
	//headCell.Draw()
	glfw.PollEvents()
	window.SwapBuffers()
}

// initGlfw initializes glfw and returns a Window to use.
func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "Snek", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	window.SetKeyCallback(handleKeys)

	return window
}

// initOpenGL initializes OpenGL and returns an intiialized program.
func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	posAttrib := uint32(gl.GetAttribLocation(prog, gl.Str("position\x00")))
	gl.VertexAttribPointer(posAttrib, 2, gl.FLOAT, false, 0, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(posAttrib)
	return prog
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		logs := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(logs))

		return 0, fmt.Errorf("failed to compile %v: %v", source, logs)
	}

	return shader, nil
}



func handleKeys(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Press {
		return
	}
	switch key {
	case glfw.KeyUp:
		log.Println("up")
		gameSnek.Move(snek.UP)
	case glfw.KeyDown:
		log.Println("down")
		gameSnek.Move(snek.DOWN)
	case glfw.KeyLeft:
		log.Println("left")
		gameSnek.Move(snek.LEFT)
	case glfw.KeyRight:
		log.Println("right")
		gameSnek.Move(snek.RIGHT)
	}
}