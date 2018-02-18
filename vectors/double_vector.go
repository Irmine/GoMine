package vectors

import "math"

type DoubleVector struct {
	x float32
	y float32
}

// Returns the X value of this DoubleVector.

func (vector *DoubleVector) GetX() float32 {
	return vector.x
}

// Returns the Y value of this DoubleVector.

func (vector *DoubleVector) GetY() float32 {
	return vector.y
}

// Sets the X value of this DoubleVector.

func (vector *DoubleVector) SetX(value float32) {
	vector.x = value
}

// Sets the Y value of this DoubleVector.

func (vector *DoubleVector) SetY(value float32) {
	vector.y = value
}

// Adds the given vector to the current vector and returns a new DoubleVector.

func (vector *DoubleVector) AddVector(vector32 DoubleVector) DoubleVector {
	return DoubleVector{vector.x + vector32.x, vector.y + vector32.y}
}

// Adds the given xyz values to the current vector and returns a new DoubleVector.

func (vector *DoubleVector) Add(x float32, y float32) DoubleVector {
	return DoubleVector{vector.x + x, vector.y + y}
}

// Subtracts the given vector from the current vector and returns a new DoubleVector.

func (vector *DoubleVector) SubtractVector(vector32 DoubleVector) DoubleVector {
	return vector.Add(-vector32.x, -vector32.y)
}

// Subtracts the given xyz values from the current vector and returns a new DoubleVector.

func (vector *DoubleVector) Subtract(x float32, y float32) DoubleVector {
	return vector.Add(-x, -y)
}

// Returns a new DoubleVector with the values made absolute.

func (vector *DoubleVector) Abs() DoubleVector {
	return DoubleVector{float32(math.Abs(float64(vector.x))), float32(math.Abs(float64(vector.y)))}
}
