package worlds

import (
	"gomine/vectorMath"
)

type Position struct {
	*vectorMath.TripleVector
	Level Level
}

func NewPosition(x, y, z float32, level Level) *Position {
	return &Position{vectorMath.NewTripleVector(x, y, z), level}
}

func (pos *Position) SetLevel(level Level) {
	pos.Level = level
}

func (pos *Position) GetLevel() Level {
	return pos.Level
}

func (pos *Position) AsPosition() *Position {
	return NewPosition(pos.GetX(), pos.GetY(), pos.GetZ(), pos.Level)
}