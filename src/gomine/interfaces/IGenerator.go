package interfaces

import "gomine/vectors"

type IGenerator interface {
	GetName() string
	SetSpawn(vectors.TripleVector)
	GetSpawn() vectors.TripleVector
	GenerateChunk(x, z int)
	PopulateChunk()
	GetLevel() ILevel
	SetLevel(ILevel)
}