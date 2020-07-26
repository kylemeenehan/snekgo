package snek

import (
	"github.com/kylemeenehan/go-opengl-play/cell"
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

func (s *Snek) Move(d int) {
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
	seg := newSnekSegment(x, y)
	s.Head.Next = seg
	s.Head.hasNext = true
	s.Head = seg
	s.Tail = s.Tail.Next
}

func NewSnek(x , y, len int) Snek {
	if len < 1 {
		log.Println("Snek must have at least 1 segment.")
		panic("Snek failed")
	}

	seg := newSnekSegment(x, y)
	newSnek := Snek{
		Tail: seg,
		Head: seg,
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