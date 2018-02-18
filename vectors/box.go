package vectors

type CubesBox struct {
	cubes []*Cube
}

// Returns a new cubes box with the given cubes.

func NewCubesBox(cubes []*Cube) *CubesBox {
	return &CubesBox{cubes}
}

// Returns all cubes in this CollisionBox.

func (box *CubesBox) GetCubes() []*Cube {
	return box.cubes
}

// Checks if the given vector is inside of this collision box.

func (box *CubesBox) IsInside(vector TripleVector) bool {
	for _, cube := range box.cubes {
		if cube.IsInside(vector) {
			return true
		}
	}
	return false
}

// Checks if this cubes box can be treated as nil.

func (box *CubesBox) IsNil() bool {
	for _, cube := range box.cubes {
		if !cube.IsNil() {
			return false
		}
	}
	return true
}

// Clears all cubes from the cubes box.

func (box *CubesBox) Clear() {
	box.cubes = []*Cube{}
}
