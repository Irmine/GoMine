package vectorMath

import "math"

type TripleVector struct {
	x float64
	y float64
	z float64
}

func (vector *TripleVector) getX() float64 {
	return vector.x
}

func (vector *TripleVector) setX(value float64) {
	vector.x = value
}

func (vector *TripleVector) getY() float64 {
	return vector.y
}

func (vector *TripleVector) setY(value float64) {
	vector.y = value
}

func (vector *TripleVector) getZ() float64 {
	return vector.z
}

func (vector *TripleVector) setZ(value float64) {
	vector.z = value
}

func (vector *TripleVector) addVector(vector2 TripleVector) TripleVector {
	return TripleVector{vector.x + vector2.x, vector.y + vector2.y, vector.z + vector2.z}
}

func (vector *TripleVector) add(x float64, y float64, z float64) TripleVector {
	return TripleVector{vector.x + x, vector.y + y, vector.z + z}
}

func (vector *TripleVector) subtract(x float64, y float64, z float64) TripleVector {
	return vector.add(-x, -y, -z)
}

func (vector *TripleVector) abs() TripleVector {
	return TripleVector{math.Abs(vector.x), math.Abs(vector.y), math.Abs(vector.z)}
}
