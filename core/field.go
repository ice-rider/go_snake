package core

type FieldObject uint8

const (
	Empty FieldObject = iota
	Wall
	SnakePart
	Apple
)

var FieldObjectString = [4]string{" ", "#", "@", "*"}

type GameField struct {
	Width, Height int
	data          [][]FieldObject
}

func NewGameField(width, height int) *GameField {
	field := make([][]FieldObject, width)
	for i := range field {
		field[i] = make([]FieldObject, height)
	}
	return &GameField{
		Width:  width,
		Height: height,
		data:   field,
	}
}

func (f *GameField) Draw(pos Point, drawType FieldObject) {
	if pos.X < 0 || pos.X >= f.Width || pos.Y < 0 || pos.Y >= f.Height {
		panic("draw pos out of bounds")
	}
	f.data[pos.X][pos.Y] = drawType
}

func (f *GameField) Get(pos Point) FieldObject {
	if pos.X < 0 || pos.X >= f.Width || pos.Y < 0 || pos.Y >= f.Height {
		return Wall // Возвращаем стену для координат за пределами поля
	}
	return f.data[pos.X][pos.Y]
}

func (f *GameField) Clear() {
	for x := 0; x < f.Width; x++ {
		for y := 0; y < f.Height; y++ {
			f.data[x][y] = Empty
		}
	}
}

func (f *GameField) String() string {
	var out string
	for x := 0; x < f.Width; x++ {
		for y := 0; y < f.Height; y++ {
			out += FieldObjectString[f.data[x][y]]
		}
		out += "\n"
	}
	return out
}