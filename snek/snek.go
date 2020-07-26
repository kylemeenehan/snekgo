package snek

import (
	"github.com/kylemeenehan/go-opengl-play/cell"
	"github.com/kylemeenehan/go-opengl-play/mouse"
	"log"
)

var UP = 0
var DOWN = 1
var LEFT = 2
var RIGHT = 3

type snekSegment struct {
	X, Y int
	Next        *snekSegment
	hasNext     bool
}

type Snek struct {
	Head *snekSegment
	Tail *snekSegment
	Direction int
	Close chan<- bool
}

func (s *Snek) Draw() {
	c := s.Tail
	cell.DrawAt(c.X, c.Y)
	for {
		c = c.Next
		cell.DrawAt(c.X, c.Y)
		if !c.hasNext {
			break
		}
	}
}

func (s *Snek) GetLength() int {
	// TODO add toIterable method
	length := 0
	c := s.Tail
	for {
		length++
		c = c.Next
		if !c.hasNext {
			break
		}
	}
	return length
}

func (s *Snek) HasXY(x, y int) (bool, bool) {
	hasX, hasY := false, false
	c := s.Tail
	for {
		if !hasX {
			hasX = c.X == x
		}
		if !hasY {
			hasY = c.Y == y
		}
		if !c.hasNext || (hasX && hasY) {
			break
		}
		c = c.Next
	}
	return hasX, hasY
}

func (s *Snek) HasSegment(x, y int) bool {
	c := s.Tail
	for {
		if c.X == x && c.Y == y {
			return true
		}
		if !c.hasNext {
			break
		}
		c = c.Next
	}
	return false
}
// Returns whether a mouse has been eaten
func (s *Snek) Move(d int, m mouse.Mouse) bool {
	x, y := s.Head.X, s.Head.Y

	switch d {
	case UP:
		y++
	case DOWN:
		y--
	case LEFT:
		x--
	case RIGHT:
		x++
	}
	x = cell.Bound(x, cell.NumColumns)
	y = cell.Bound(y, cell.NumRows)
	if s.HasSegment(x, y) {
		s.Close <- true
	}
	seg := newSnekSegment(x, y)
	s.Head.Next = seg
	s.Head.hasNext = true
	s.Head = seg
	if !(m.X == x && m.Y == y) {
		s.Tail = s.Tail.Next
		return false
	}
	return true
}

func NewSnek(x , y, len int, close chan<- bool) Snek {
	if len < 1 {
		log.Println("Snek must have at least 1 segment.")
		panic("Snek must have at least 1 segment.")
	}

	seg := newSnekSegment(x, y)
	newSnek := Snek{
		Tail: seg,
		Head: seg,
		Direction: RIGHT,
		Close: close,
	}

	for i := 1; i < len;  i++ {
		x++
		seg = newSnekSegment(x, y)
		newSnek.Head.hasNext = true
		newSnek.Head.Next = seg
		newSnek.Head = seg

	}
	return newSnek
}

func newSnekSegment(x, y int) *snekSegment {
	segment := snekSegment {
		X: x,
		Y: y,
		Next: nil,
		hasNext: false,
	}
	return &segment
}
