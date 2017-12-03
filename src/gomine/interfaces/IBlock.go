package interfaces

import "gomine/vectors"

type IBlock interface{
	GetId() int
	GetData() byte
	SetVariant(byte)
	SetData(byte)
	GetName() string
	HasCollisionBox() bool
	GetCollisionBox() *vectors.CubesBox
	SetCollisionBox(box *vectors.CubesBox)
	GetBoundingBox() *vectors.CubesBox
	SetBoundingBox(box *vectors.CubesBox)
	// IsSolid() bool
	// GetLightFilter() int
	// GetHardness() float64
	// GetBlastResistance() float64
	// Place(player IPlayer, vector vectors.TripleVector)
}
