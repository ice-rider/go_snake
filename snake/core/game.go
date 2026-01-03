package core

type Game struct {
	Field     *GameField
	Snake     *Snake
	ApplePos  *Point
	MaxScore  int
}

func NewGame(width, height int) *Game {
	if width < 7 || height < 7 {
		panic("Minimal field is 7x7")
	}

	field := NewGameField(width, height)
	centerPoint := Point{
		X: width / 2,
		Y: height / 2,
	}
	snake := NewSnake(centerPoint, Right)

	return &Game{
		Field:    field,
		Snake:    snake,
		ApplePos: nil,
		MaxScore: width * height,
	}
}

func (g *Game) generateApple() {
	for g.Snake.Size < g.MaxScore {
		randomPoint := generatePoint(g.Field.Width, g.Field.Height)
		if g.Field.Get(randomPoint) == Empty {
			g.ApplePos = &randomPoint
			break
		}
	}
}

func (g *Game) Start() {
	g.drawField()
	g.generateApple()
}

func (g *Game) drawField() {
	g.Field.Clear()
	g.Snake.Draw(g.Field)

	// Рисуем границы поля
	for x := 0; x < g.Field.Width; x++ {
		g.Field.Draw(Point{X: x, Y: 0}, Wall)
		g.Field.Draw(Point{X: x, Y: g.Field.Height - 1}, Wall)
	}
	for y := 0; y < g.Field.Height; y++ {
		g.Field.Draw(Point{X: 0, Y: y}, Wall)
		g.Field.Draw(Point{X: g.Field.Width - 1, Y: y}, Wall)
	}

	if g.ApplePos != nil {
		g.Field.Draw(*g.ApplePos, Apple)
	}
}

func (g *Game) RotateSnake(newDirection Direction) {
	g.Snake.Direction = newDirection
}

func (g *Game) Next() {
	nextPos := g.Snake.NextHeadPos()

	// Проверяем, не выходит ли змейка за границы
	if nextPos.X < 0 || nextPos.X >= g.Field.Width || nextPos.Y < 0 || nextPos.Y >= g.Field.Height {
		g.GameOver()
		return
	}

	switch g.Field.Get(*nextPos) {
	case Empty:
		g.Snake.Move(0)
	case Wall:
		g.GameOver()
	case SnakePart:
		g.GameOver()
	case Apple:
		g.Snake.Move(1)
		g.generateApple()
	default:
		return
	}
	g.drawField()
}

func (g *Game) GameOver() {
	panic("game over")
}