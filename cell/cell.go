package cell

import "github.com/go-gl/gl/v4.6-core/gl"

var (
	width      int
	height     int
	NumRows    int
	NumColumns int
	cellMatrix [][]*Cell
)

func Init(w, h, r, c int) {
	width = w
	height = h
	NumRows = r
	NumColumns = c
	cellMatrix = MakeCells(c, r)
}

var square = []float32{
	-0.5, 0.5,
	-0.5, -0.5,
	0.5, -0.5,

	-0.5, 0.5,
	0.5, 0.5,
	0.5, -0.5,
}

type Coordinates struct {
	X int
	Y int
}

func (a *Coordinates) GoUp() {
	newVal := a.Y + 1
	a.Y = Bound(newVal, NumRows)
}

func (a *Coordinates) GoDown() {
	newVal := a.Y - 1
	a.Y = Bound(newVal, NumRows)
}

func (a *Coordinates) GoLeft() {
	newVal := a.X - 1
	a.X = Bound(newVal, NumColumns)
}

func (a *Coordinates) GoRight() {
	newVal := a.X + 1
	a.X = Bound(newVal, NumColumns)
}

func (c *Coordinates) Draw() {
	cellMatrix[c.X][c.Y].Draw()
}

func Bound(current int, max int) int {
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

func MakeCells(numColumns, numRows int) [][]*Cell {
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
			size = 1.0 / float32(NumColumns)
			position = float32(x) * size
		} else {
			size = 1.0 / float32(NumRows)
			position = float32(y) * size
		}

		if point < 0 {
			points[i] = (position * 2) - 1
		} else {
			points[i] = ((position + size) * 2) - 1
		}
	}

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
