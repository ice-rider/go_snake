package core

const (
	SnakeDrawType    = SnakePart
	StartSnakeSize   = 3
	MinimalSnakeSize = 1
)

type Snake struct {
	Size      int
	Direction Direction
	Head      *Segment
	Tail      *Segment
}

func NewSnake(head Point, direction Direction) *Snake {
	headSegment := &Segment{Next: nil, Pos: head}
	tailSegment := &Segment{Next: headSegment, Pos: head}

	snake := &Snake{
		Size:      MinimalSnakeSize,
		Direction: direction,
		Head:      headSegment,
		Tail:      tailSegment,
	}

	for snake.Size < StartSnakeSize {
		snake.Move(1)
	}
	return snake
}

func (s *Snake) Move(grow int) {
	newHead := &Segment{
		Next: nil,
		Pos:  *s.NextHeadPos(),
	}
	s.Head.Next = newHead
	s.Head = newHead

	if grow == 0 {
		s.Tail = s.Tail.Next
		s.Size--
	} else {
		s.Size++
	}
}

func (s *Snake) Draw(canvas Drawable) {
	current := s.Tail
	for current != nil {
		if current.Next != nil {
			canvas.Draw(current.Pos, SnakeDrawType)
		}
		current = current.Next
	}
	canvas.Draw(s.Head.Pos, SnakeDrawType)
}

func (s *Snake) NextHeadPos() *Point {
	return s.Head.Pos.Add(MoveMatrix[s.Direction])
}