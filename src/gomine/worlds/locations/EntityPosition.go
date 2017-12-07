package locations

import (
	"gomine/players/math"
	"gomine/interfaces"
	"gomine/vectors"
)

type EntityPosition struct {
	*vectors.TripleVector
	Level interfaces.ILevel
	Dimension interfaces.IDimension
	Rotation math.Rotation
}

func NewEntityPosition(x, y, z, pitch, yaw, headYaw float32, level interfaces.ILevel, dimension interfaces.IDimension) EntityPosition {
	return EntityPosition{vectors.NewTripleVector(x, y, z), level, dimension, math.NewRotation(pitch, yaw, headYaw)}
}

/**
 * Sets the level of this position.
 */
func (pos *EntityPosition) SetLevel(level interfaces.ILevel) {
	pos.Level = level
}

/**
 * Returns the level of this position.
 */
func (pos *EntityPosition) GetLevel() interfaces.ILevel {
	return pos.Level
}

/**
 * Returns the dimension of this position.
 */
func (pos *EntityPosition) GetDimension() interfaces.IDimension {
	return pos.Dimension
}

/**
 * Sets the dimension of this position.
 */
func (pos *EntityPosition) SetDimension(dimension interfaces.IDimension) {
	pos.Dimension = dimension
}

/**
 * Returns the current instance as a position.
 */
func (pos *EntityPosition) AsPosition() *Position {
	return NewPosition(pos.GetX(), pos.GetY(), pos.GetZ(), pos.Level, pos.Dimension)
}

/**
 * Sets the rotation of this position
 */
func (pos *EntityPosition) SetRotation(rot math.Rotation) {
	pos.Rotation = rot
}

/**
 * returns the rotation of this position
 */
func (pos *EntityPosition) GetRotation() math.Rotation {
	return pos.Rotation
}