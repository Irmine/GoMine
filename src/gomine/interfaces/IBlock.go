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
	SetCollisionBox(*vectors.CubesBox)
	GetBoundingBox() *vectors.CubesBox
	SetBoundingBox(*vectors.CubesBox)
	GetLightEmissionLevel() byte
	SetLightEmissionLevel(byte)
	GetBlastResistance() int
	SetBlastResistance(int)
	SetHardness(float32)
	GetHardness() float32
	DiffusesLight() bool
	SetLightDiffusing(bool)
	GetLightFilterLevel() byte
	SetLightFilterLevel(byte)
	// IsSolid() bool
	// Place(player IPlayer, vector vectors.TripleVector)
}
