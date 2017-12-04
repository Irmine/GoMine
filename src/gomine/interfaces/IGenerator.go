package interfaces

type IGenerator interface {
	GetName() string
	GetNewChunk(IChunk) IChunk
	GenerateChunk(IChunk)
	PopulateChunk(IChunk)
}