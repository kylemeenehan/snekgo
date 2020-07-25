package cell

import "github.com/go-gl/gl/v4.6-core/gl"

var (
	width int
	height int
	numRows int
	numColumns int
)

func Init(w, h, r, c int) {
	width = w
	height = h
	numRows = r
	numColumns = c
}

var square = []float32{
	-0.5, 0.5,
	-0.5, -0.5,
	0.5, -0.5,

	-0.5, 0.5,
	0.5, 0.5,
	0.5, -0.5,
}

type ActiveCell struct {
	X int
	Y int
}

func (a *ActiveCell) GoUp() {
	newVal := a.Y + 1
	a.Y = bound(newVal, numRows)
}

func (a *ActiveCell) GoDown() {
	newVal := a.Y - 1
	a.Y = bound(newVal, numRows)
}

func (a *ActiveCell) GoLeft() {
	newVal := a.X - 1
	a.X = bound(newVal, numColumns)
}

func (a *ActiveCell) GoRight() {
	newVal := a.X + 1
	a.X = bound(newVal, numColumns)
}

func bound(current int, max int) int {
	arrayLimit := max - 1
	if current > arrayLimit {
		return 0
	} else if current < 0 {
		return arrayLimit
	}
	return current
}

type Cell struct {
	drawable uint32

	x int
	y int
}

func (c *Cell) Draw() {
	gl.BindVertexArray(c.drawable)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(square) / 2))
}

func MakeCells() [][]*Cell {
	cells := make([][]*Cell, numColumns)
	for x := range cells {
		cells[x] = make([]*Cell, numRows)
		for y := range cells[x] {
			cells[x][y] = newCell(x, y)
		}
	}
	return cells
}
func newCell(x, y int) *Cell {
	points := make([]float32, len(square), len(square))
	copy(points, square)

	for i, point := range points {
		var position float32
		var size float32

		if mod := i % 2; mod == 0 {
			size = 1.0 / float32(numColumns)
			position = float32(x) * size
		} else {
			size = 1.0 / float32(numRows)
			position = float32(y) * size
		}

		if point < 0 {
			points[i] = (position * 2) - 1
		} else {
			points[i] = ((position + size) * 2) - 1
		}
	}

	//for i := 0; i < len(points); i++ {
	//	var position float32
	//	var size float32
	//	switch i % 3 {
	//	case 0:
	//		size = 1.0 / float32(numColumns)
	//		position = float32(x) * size
	//	case 1:
	//		size = 1.0 / float32(numRows)
	//		position = float32(y) * size
	//	default:
	//		continue
	//	}
	//
	//	if points[i] < 0 {
	//		points[i] = (position * 2) - 1
	//	} else {
	//		points[i] = ((position + size) * 2) - 1
	//	}
	//}

	return &Cell{
		drawable: makeVao(points),

		x: x,
		y: y,
	}
}

// makeVao initializes and returns a vertex array from the points provided.
func makeVao(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 0, nil)

	return vao
}
