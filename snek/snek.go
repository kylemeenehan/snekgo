package snek

import "github.com/kylemeenehan/go-opengl-play/cell"

var UP = 0
var DOWN = 1
var LEFT = 2
var RIGHT = 3

type snekSegment struct {
	Coordinates *cell.Coordinates
	Next        *snekSegment
	hasNext     bool
}

type Snek struct {
	Head *snekSegment
	Tail *snekSegment
}

func (s *Snek) Draw() {
	c := s.Tail
	c.Coordinates.Draw()
	for {
		c = c.Next
		c.Coordinates.Draw()
		if !c.hasNext {
			break
		}
	}
}

func (s *Snek) Move(d int) {
	x, y := s.Head.Coordinates.X, s.Head.Coordinates.Y

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
	c := cell.Coordinates {
		X: x,
		Y: y,
	}
	seg := newSnekSegment(c)
	s.Head.Next = seg
	s.Head.hasNext = true
	s.Head = seg
	s.Tail = s.Tail.Next
}

func NewSnek(x , y int) Snek {
	a := cell.Coordinates {
		X: x,
		Y: y,
	}
	seg1 := newSnekSegment(a)

	b := cell.Coordinates {
		X: x + 1,
		Y: y,
	}
	seg2 := newSnekSegment(b)

	c := cell.Coordinates {
		X: x + 2,
		Y: y,
	}
	seg3 := newSnekSegment(c)

	seg1.Next = seg2
	seg1.hasNext = true
	seg2.Next = seg3
	seg2.hasNext = true

	s := Snek {
		Head: seg3,
		Tail: seg1,
	}
	return s
}

func newSnekSegment(c cell.Coordinates) *snekSegment {
	segment := snekSegment {
		Coordinates: &c,
		Next: nil,
		hasNext: false,
	}
	return &segment
}