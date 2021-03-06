package main

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/kylemeenehan/snekgo/cell"
	"github.com/kylemeenehan/snekgo/graphics"
	"github.com/kylemeenehan/snekgo/mouse"
	"github.com/kylemeenehan/snekgo/snek"
	"math/rand"
	"runtime"
	"time"
)

const (
	width   = 500
	height  = 500
	rows    = 15
	columns = 15
)

var (
	window            *glfw.Window
	gameSnek          snek.Snek
	ActiveMouse       mouse.Mouse
	previousDirection int
)

func main() {
	runtime.LockOSThread()
	window = graphics.InitGlfw(width, height)
	window.SetKeyCallback(handleKeys)
	defer glfw.Terminate()
	program := graphics.InitOpenGL()
	cell.Init(rows, columns)
	closeWindow := make(chan bool, 1)
	gameSnek = snek.NewSnek(0, 0, 5, closeWindow)
	makeMouse(0)
	frame := time.Tick(time.Second / 10)
	for !window.ShouldClose() {
		select {
		case <-closeWindow:
			window.SetShouldClose(true)
		default:
			mouseEaten := gameSnek.Move(gameSnek.Direction, ActiveMouse)
			if mouseEaten {
				makeMouse(0)
			}
			draw(window, program)
			<-frame
		}
	}
}

func draw(window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)
	gameSnek.Draw()
	ActiveMouse.Draw()
	previousDirection = gameSnek.Direction
	glfw.PollEvents()
	window.SwapBuffers()
}

func handleKeys(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Press {
		return
	}
	switch key {
	case glfw.KeyUp:
		if previousDirection != snek.DOWN {
			gameSnek.Direction = snek.UP
		}
	case glfw.KeyDown:
		if previousDirection != snek.UP {
			gameSnek.Direction = snek.DOWN
		}
	case glfw.KeyLeft:
		if previousDirection != snek.RIGHT {
			gameSnek.Direction = snek.LEFT
		}
	case glfw.KeyRight:
		if previousDirection != snek.LEFT {
			gameSnek.Direction = snek.RIGHT
		}
	case glfw.KeyEscape:
		window.SetShouldClose(true)
	}
}

func makeMouse(numTries int) {
	maxTries := 100
	x := rand.Intn(columns)
	y := rand.Intn(rows)
	//hasX, hasY := gameSnek.HasSegment(x, y)
	if !(gameSnek.HasSegment(x, y)) {
		ActiveMouse = mouse.NewMouse(x, y)
	} else {
		numTries++
		if numTries >= maxTries {
			panic("too many tries")
		}
		makeMouse(numTries)
		// TODO: optimise to scan
		//for numTries < maxTries {
		//	switch {
		//	case hasX && hasY:
		//		x++
		//		y++
		//	case hasX:
		//		x++
		//	}
		//}
	}
}
