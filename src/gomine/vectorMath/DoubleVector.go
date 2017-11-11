package vectorMath

import "math"

type DoubleVector struct {
	x float64
	y float64
}

/**
 * Returns the X value of this DoubleVector.
 */
func (vector *DoubleVector) GetX() float64 {
	return vector.x
}

/**
 * Returns the Y value of this DoubleVector.
 */
func (vector *DoubleVector) GetY() float64 {
	return vector.y
}

/**
 * Sets the X value of this DoubleVector.
 */
func (vector *DoubleVector) SetX(value float64) {
	vector.x = value
}

/**
 * Sets the Y value of this DoubleVector.
 */
func (vector *DoubleVector) SetY(value float64) {
	vector.y = value
}

/**
 * Adds the given vector to the current vector and returns a new DoubleVector.
 */
func (vector *DoubleVector) AddVector(vector2 DoubleVector) DoubleVector {
	return DoubleVector{vector.x + vector2.x, vector.y + vector2.y}
}

/**
 * Adds the given xyz values to the current vector and returns a new DoubleVector.
 */
func (vector *DoubleVector) Add(x float64, y float64) DoubleVector {
	return DoubleVector{vector.x + x, vector.y + y}
}

/**
 * Subtracts the given vector from the current vector and returns a new DoubleVector.
 */
func (vector *DoubleVector) SubtractVector(vector2 DoubleVector) DoubleVector {
	return vector.Add(-vector2.x, -vector2.y)
}

/**
 * Subtracts the given xyz values from the current vector and returns a new DoubleVector.
 */
func (vector *DoubleVector) Subtract(x float64, y float64) DoubleVector {
	return vector.Add(-x, -y)
}

/**
 * Returns a new DoubleVector with the values made absolute.
 */
func (vector *DoubleVector) Abs() DoubleVector {
	return DoubleVector{math.Abs(vector.x), math.Abs(vector.y)}
}
