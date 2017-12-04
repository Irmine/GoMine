package locations

import (
	"gomine/vectors"
	"gomine/interfaces"
)

type Position struct {
	*vectors.TripleVector
	Level interfaces.ILevel
	Dimension interfaces.IDimension
}

func NewPosition(x, y, z float32, level interfaces.ILevel, dimension interfaces.IDimension) *Position {
	return &Position{vectors.NewTripleVector(x, y, z), level, dimension}
}

/**
 * Sets the level of this position.
 */
func (pos *Position) SetLevel(level interfaces.ILevel) {
	pos.Level = level
}

/**
 * Returns the level of this position.
 */
func (pos *Position) GetLevel() interfaces.ILevel {
	return pos.Level
}

/**
 * Returns the dimension of this position.
 */
func (pos *Position) GetDimension() interfaces.IDimension {
	return pos.Dimension
}

/**
 * Sets the dimension of this position.
 */
func (pos *Position) SetDimension(dimension interfaces.IDimension) {
	pos.Dimension = dimension
}

/**
 * Returns the current instance as a position.
 */
func (pos *Position) AsPosition() *Position {
	return NewPosition(pos.GetX(), pos.GetY(), pos.GetZ(), pos.Level, pos.Dimension)
}