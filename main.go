package main

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/kylemeenehan/go-opengl-play/cell"
	"github.com/kylemeenehan/go-opengl-play/graphics"
	"github.com/kylemeenehan/go-opengl-play/snek"
	"log"
	"runtime"
	"time"
)

const (
	width  = 500
	height = 500
	rows = 10
	columns = 10
)

var window *glfw.Window
var gameSnek snek.Snek

func main() {
	runtime.LockOSThread()
	window = graphics.InitGlfw(width, height)
	window.SetKeyCallback(handleKeys)
	defer glfw.Terminate()
	program := graphics.InitOpenGL()
	cell.Init(width, height, rows, columns)
	gameSnek = snek.NewSnek(0, 0, 5)
	for !window.ShouldClose() {
		gameSnek.Move(gameSnek.Direction)
		draw(window, program)
		time.Sleep(time.Second / 5)
	}
}

func draw(window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)
	gameSnek.Draw()
	glfw.PollEvents()
	window.SwapBuffers()
}

func handleKeys(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Press {
		return
	}
	switch key {
	case glfw.KeyUp:
		log.Println("up")
		gameSnek.Direction = snek.UP
	case glfw.KeyDown:
		gameSnek.Direction = snek.DOWN
	case glfw.KeyLeft:
		gameSnek.Direction = snek.LEFT
	case glfw.KeyRight:
		gameSnek.Direction = snek.RIGHT
	case glfw.KeyEscape:
		window.SetShouldClose(true)
	}
}