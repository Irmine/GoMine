package interfaces

import "gomine/vectors"

type IBlock interface{
	GetId() int
	GetData() byte
	SetData(byte)
	GetName() string
	HasCollisionBox() bool
	// IsSolid() bool
	// GetLightFilter() int
	// GetHardness() float64
	// GetBlastResistance() float64
	// Place(player IPlayer, vector vectors.TripleVector)
}
