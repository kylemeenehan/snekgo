package mouse

import "github.com/kylemeenehan/go-opengl-play/cell"

type Mouse struct {
	X, Y int
}

func NewMouse(x, y int) Mouse {
	m := Mouse {
		X: x,
		Y: y,
	}
	return m
}

func (m Mouse) Draw() {
	cell.DrawAt(m.X, m.Y)
}