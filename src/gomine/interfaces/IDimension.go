package interfaces

type IDimension interface {
	GetDimensionId() int
	GetLevel() ILevel
	GetName() string
	TickDimension()
	SetChunk(int, int, IChunk)
	GetChunk(int, int) IChunk
	GetChunkPlayers(x, z int) []IPlayer
	AddChunkPlayer(int, int, IPlayer)
}
