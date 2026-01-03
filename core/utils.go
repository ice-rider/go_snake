package core

import (
	"math/rand"
	"time"
)

// Point represents a coordinate on the game field
type Point struct {
	X, Y int
}

func (p *Point) Add(target Point) *Point {
	return &Point{X: p.X + target.X, Y: p.Y + target.Y}
}

// Direction represents movement direction
type Direction uint8

const (
	Up Direction = iota
	Right
	Down
	Left
)

// Segment represents a part of the snake
type Segment struct {
	Next *Segment
	Pos  Point
}

var MoveMatrix = map[Direction]Point{
	Left:  {X: 0, Y: -1},
	Down:  {X: 1, Y: 0},
	Right: {X: 0, Y: 1},
	Up:    {X: -1, Y: 0},
}

// Drawable interface for objects that can be drawn on the field
type Drawable interface {
	Draw(pos Point, drawType FieldObject)
}

// Random number generator
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// generatePoint creates a random point within the field boundaries (excluding walls)
func generatePoint(width, height int) Point {
	return Point{
		X: 1 + rng.Intn(width-2),  // Exclude left and right walls
		Y: 1 + rng.Intn(height-2), // Exclude top and bottom walls
	}
}