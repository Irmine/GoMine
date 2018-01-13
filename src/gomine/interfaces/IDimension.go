package interfaces

type IDimension interface {
	GetDimensionId() int
	GetLevel() ILevel
	GetName() string
	TickDimension()
	SetChunk(int32, int32, IChunk)
	GetChunk(int32, int32) IChunk
	RequestChunks(IPlayer, int32)
	IsGenerated() bool
	SetGenerator(IGenerator)
	GetGenerator() IGenerator
}

type IEntityHelper interface {
	SpawnPlayerTo(IPlayer, IPlayer)
	SpawnEntityTo(IEntity, IPlayer)
	DespawnEntityFrom(IEntity, IPlayer)
	SendEntityData(IEntity, IPlayer)
}