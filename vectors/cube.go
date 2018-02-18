package vectors

import "github.com/golang/geo/r3"

type Cube struct {
	MaxX, MaxY, MaxZ float64
	MinX, MinY, MinZ float64
}

// Returns a new cube with the given points.

func NewCube(minX, minY, minZ, maxX, maxY, maxZ float64) *Cube {
	return &Cube{maxX, maxY, maxZ, minX, minY, minZ}
}

// Returns the minimum X of this cube.

func (cube *Cube) GetMinX() float64 {
	return cube.MinX
}

// Returns the minimum Y of this cube.

func (cube *Cube) GetMinY() float64 {
	return cube.MinY
}

// Returns the minimum Z of this cube.

func (cube *Cube) GetMinZ() float64 {
	return cube.MinZ
}

// Sets the minimum X of this cube.

func (cube *Cube) SetMinX(value float64) {
	if value > cube.MaxX {
		return
	}
	cube.MinX = value
}

// Sets the minimum Y of this cube.

func (cube *Cube) SetMinY(value float64) {
	if value > cube.MaxY {
		return
	}
	cube.MinY = value
}

// Sets the minimum Z of this cube.

func (cube *Cube) SetMinZ(value float64) {
	if value > cube.MaxZ {
		return
	}
	cube.MinZ = value
}

// Returns the maximum X of this cube.

func (cube *Cube) GetMaxX() float64 {
	return cube.MaxX
}

// Returns the maximum Y of this cube.

func (cube *Cube) GetMaxY() float64 {
	return cube.MaxY
}

// Returns the maximum Z of this cube.

func (cube *Cube) GetMaxZ() float64 {
	return cube.MaxZ
}

// Sets the maximum X of this cube.

func (cube *Cube) SetMaxX(value float64) {
	if value < cube.MinX {
		return
	}
	cube.MaxX = value
}

// Sets the maximum Y of this cube.

func (cube *Cube) SetMaxY(value float64) {
	if value < cube.MinY {
		return
	}
	cube.MaxY = value
}

// Sets the maximum Z of this cube.

func (cube *Cube) SetMaxZ(value float64) {
	if value < cube.MinZ {
		return
	}
	cube.MaxZ = value
}

// Returns all vectors within the cube with the given density.

func (cube *Cube) GetVectorsWithin(density float64) []r3.Vector {
	var vectors []r3.Vector
	for x := cube.MinX; x <= cube.MaxX; x += density {
		for y := cube.MinY; y <= cube.MaxY; y += density {
			for z := cube.MinZ; z <= cube.MinZ; z += density {
				vectors = append(vectors, r3.Vector{x, y, z})
			}
		}
	}
	return vectors
}

// Checks whether the given vector is within this cube or not.

func (cube *Cube) IsInside(vector r3.Vector) bool {
	return vector.X <= cube.MaxX && vector.X >= cube.MinX &&
		vector.Y <= cube.MaxY && vector.Y >= cube.MinY &&
		vector.Z <= cube.MaxZ && vector.Z >= cube.MinZ
}

// Checks if this cube can be treated as nil.

func (cube *Cube) IsNil() bool {
	return ((cube.MaxX - cube.MinX) * (cube.MaxY - cube.MinY) * (cube.MaxZ - cube.MinZ)) == float64(0)
}
