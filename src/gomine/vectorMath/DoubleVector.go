package vectorMath

import "math"

type DoubleVector struct {
	x float64
	y float64
}

func (vector *DoubleVector) getX() float64 {
	return vector.x
}

func (vector *DoubleVector) getY() float64 {
	return vector.y
}

func (vector *DoubleVector) setX(value float64) {
	vector.x = value
}

func (vector *DoubleVector) setY(value float64) {
	vector.y = value
}

func (vector *DoubleVector) addVector(vector2 DoubleVector) DoubleVector {
	return DoubleVector{vector.x + vector2.x, vector.y + vector2.y}
}

func (vector *DoubleVector) add(x float64, y float64) DoubleVector {
	return DoubleVector{vector.x + x, vector.y + y}
}

func (vector *DoubleVector) subtractVector(vector2 DoubleVector) DoubleVector {
	return vector.add(-vector2.x, -vector2.y)
}

func (vector *DoubleVector) subtract(x float64, y float64) DoubleVector {
	return vector.add(-x, -y)
}

func (vector *DoubleVector) abs() DoubleVector {
	return DoubleVector{math.Abs(vector.x), math.Abs(vector.y)}
}
