package interfaces

type IDimension interface {
	GetDimensionId() int
	GetLevel() ILevel
	GetName() string
	TickDimension()
	SetChunk(int32, int32, IChunk)
	GetChunk(int32, int32) IChunk
	GetChunkPlayers(int32, int32) []IPlayer
	AddChunkPlayer(int32, int32, IPlayer)
	RequestChunks(IPlayer, int32)
	IsGenerated() bool
	SetGenerator(IGenerator)
	GetGenerator() IGenerator
}