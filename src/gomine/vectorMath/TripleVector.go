package vectorMath

import "math"

type TripleVector struct {
	x float64
	y float64
	z float64
}

/**
 * Returns the X value of this TripleVector.
 */
func (vector *TripleVector) GetX() float64 {
	return vector.x
}

/**
 * Sets the X value of this TripleVector.
 */
func (vector *TripleVector) SetX(value float64) {
	vector.x = value
}

/**
 * Returns the Y value of this TripleVector.
 */
func (vector *TripleVector) GetY() float64 {
	return vector.y
}

/**
 * Sets the Y value of this TripleVector.
 */
func (vector *TripleVector) SetY(value float64) {
	vector.y = value
}

/**
 * Returns the Z value of this TripleVector.
 */
func (vector *TripleVector) GetZ() float64 {
	return vector.z
}

/**
 * Sets the Z value of this TripleVector.
 */
func (vector *TripleVector) SetZ(value float64) {
	vector.z = value
}

/**
 * Adds the given vector to the current vector and creates a new TripleVector.
 */
func (vector *TripleVector) AddVector(vector2 TripleVector) TripleVector {
	return TripleVector{vector.x + vector2.x, vector.y + vector2.y, vector.z + vector2.z}
}

/**
 * Adds the given xyz values to the current vector and creates a new TripleVector.
 */
func (vector *TripleVector) Add(x float64, y float64, z float64) TripleVector {
	return TripleVector{vector.x + x, vector.y + y, vector.z + z}
}

/**
 * Subtracts the given vector from the current vector and creates a new TripleVector.
 */
func (vector *TripleVector) SubtractVector(vector2 TripleVector) TripleVector {
	return vector.Add(-vector2.x, -vector2.y, -vector2.z)
}

/**
 * Subtracts the given xyz values from the current vector and creates a new TripleVector.
 */
func (vector *TripleVector) Subtract(x float64, y float64, z float64) TripleVector {
	return vector.Add(-x, -y, -z)
}

/**
 * Returns a new TripleVector with the current values made absolute.
 */
func (vector *TripleVector) Abs() TripleVector {
	return TripleVector{math.Abs(vector.x), math.Abs(vector.y), math.Abs(vector.z)}
}
