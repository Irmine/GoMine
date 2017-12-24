package vectors

import (
	"math"
)

type TripleVector struct {
	X float32
	Y float32
	Z float32
}

const (
	Down = iota
	Up
	North
	South
	West
	East
)

func NewTripleVector(x, y, z float32) *TripleVector {
	return &TripleVector{x, y, z}
}

/**
 * Converts any struct that has an embedded TripleVector to a new TripleVector.
 */
func (vector *TripleVector) AsTripleVector() *TripleVector {
	return NewTripleVector(vector.X, vector.Y, vector.Z)
}

/**
 * Sets the coordinates of this vector
 */
func (vector *TripleVector) SetVector(vector2 *TripleVector)  {
	vector.X = vector2.X
	vector.Y = vector2.Y
	vector.Z = vector2.Z
}

/**
 * Returns the X value of this TripleVector.
 */
func (vector *TripleVector) GetX() float32 {
	return vector.X
}

/**
 * Sets the X value of this TripleVector.
 */
func (vector *TripleVector) SetX(value float32) {
	vector.X = value
}

/**
 * Returns the Y value of this TripleVector.
 */
func (vector *TripleVector) GetY() float32 {
	return vector.Y
}

/**
 * Sets the Y value of this TripleVector.
 */
func (vector *TripleVector) SetY(value float32) {
	vector.Y = value
}

/**
 * Returns the Z value of this TripleVector.
 */
func (vector *TripleVector) GetZ() float32 {
	return vector.Z
}

/**
 * Sets the Z value of this TripleVector.
 */
func (vector *TripleVector) SetZ(value float32) {
	vector.Z = value
}

/**
 * Sets the X, Y and Z value of this TripleVector.
 */
func (vector *TripleVector) SetComponents(x, y, z float32) {
	vector.X = x
	vector.Y = y
	vector.Z = z
}

/**
 * Returns a slice containing the X, Y and Z values of this TripleVector.
 */
func (vector *TripleVector) GetComponents() []float32 {
	return []float32{vector.X, vector.Y, vector.Z}
}

/**
 * Adds the given vector to the current vector and creates a new TripleVector.
 */
func (vector *TripleVector) AddVector(vector2 TripleVector) TripleVector {
	return TripleVector{vector.X + vector2.X, vector.Y + vector2.Y, vector.Z + vector2.Z}
}

/**
 * Adds the given XYZ values to the current vector and creates a new TripleVector.
 */
func (vector *TripleVector) Add(x float32, y float32, z float32) TripleVector {
	return TripleVector{vector.X + x, vector.Y + y, vector.Z + z}
}

/**
 * Subtracts the given vector from the current vector and creates a new TripleVector.
 */
func (vector *TripleVector) SubtractVector(vector2 TripleVector) TripleVector {
	return vector.Add(-vector2.X, -vector2.Y, -vector2.Z)
}

/**
 * Subtracts the given XYZ values from the current vector and creates a new TripleVector.
 */
func (vector *TripleVector) Subtract(x float32, y float32, z float32) TripleVector {
	return vector.Add(-x, -y, -z)
}

/**
 * Returns a new TripleVector with the current values made absolute.
 */
func (vector *TripleVector) Abs() TripleVector {
	return TripleVector{float32(math.Abs(float64(vector.X))), float32(math.Abs(float64(vector.Y))), float32(math.Abs(float64(vector.Z)))}
}

/**
 * Multiplies all vectors by the given amount and returns a new TripleVector.
 */
func (vector *TripleVector) Multiply(value float32) TripleVector {
	return TripleVector{vector.X * value, vector.Y * value, vector.Z * value}
}

/**
 * Divides all vectors by the given amount and returns a new TripleVector.
 */
func (vector *TripleVector) Divide(value float32) TripleVector {
	return TripleVector{vector.X / value, vector.Y / value, vector.Z / value}
}

/**
 * Returns the square distance between this TripleVector and the given TripleVector.
 */
func (vector *TripleVector) SquareDistance(vector2 TripleVector) float32 {
	return ((vector.X - vector2.X) * (vector.X - vector2.X)) + ((vector.Y - vector2.Y) * (vector.Y - vector2.Y)) + ((vector.Z - vector2.Z) * (vector.Z - vector2.Z))
}

/**
 * Returns the distance between this TripleVector and the given TripleVector.
 */
func (vector *TripleVector) Distance(vector2 TripleVector) float32 {
	return float32(math.Sqrt(float64(vector.SquareDistance(vector2))))
}

/**
 * Walks between two TripleVectors and returns TripleVectors with the given interval.
 */
func (vector *TripleVector) Walk(vector2 TripleVector, interval float32) []TripleVector {
	var distance = vector.Distance(vector2)

	var xRelative = (vector2.X - vector.X) / distance * interval
	var yRelative = (vector2.Y - vector.Y) / distance * interval
	var zRelative = (vector2.Z - vector.Z) / distance * interval

	var vectors []TripleVector

	var distanceRelative = distance / interval
	for i := float32(1); i < distanceRelative; i++ {
		vectors = append(vectors, vector.Add(xRelative * i, yRelative * i, zRelative * i))
	}

	return vectors
}

/**
 * Returns a TripleVector slice with all adjacent vectors.
 */
func (vector *TripleVector) GetAdjacentVectors() []TripleVector {
	return []TripleVector{
		{vector.X + 1, vector.Y, vector.Z},
		{vector.X - 1, vector.Y, vector.Z},
		{vector.X, vector.Y + 1, vector.Z},
		{vector.X, vector.Y - 1, vector.Z},
		{vector.X, vector.Y, vector.Z + 1},
		{vector.X, vector.Y, vector.Z - 1},
	}
}

/**
 * Steps the given amount to the given direction and changes the x/y/z.
 */
func (vector *TripleVector) Step(direction int, steps float32) {
	switch direction {
	case Down:
		vector.Y -= steps
	case Up:
		vector.Y += steps
	case North:
		vector.Z -= steps
	case South:
		vector.Z += steps
	case West:
		vector.X -= steps
	case East:
		vector.X += steps
	}
}

/**
 * Gets the adjacent vector to the given direction and returns a new TripleVector.
 */
func (vector *TripleVector) GetAdjacent(direction int, steps float32) TripleVector {
	switch direction {
	case Down:
		return TripleVector{vector.X, vector.Y - steps, vector.Z}
	case Up:
		return TripleVector{vector.X, vector.Y + steps, vector.Z}
	case North:
		return TripleVector{vector.X, vector.Y, vector.Z - steps}
	case South:
		return TripleVector{vector.X, vector.Y, vector.Z + steps}
	case West:
		return TripleVector{vector.X - steps, vector.Y, vector.Z}
	case East:
		return TripleVector{vector.X + steps, vector.Y, vector.Z}
	}
	return TripleVector{}
}