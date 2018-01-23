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

type IGenerator interface {
	GetName() string
	GetNewChunk(IChunk) IChunk
	GenerateChunk(IChunk)
	PopulateChunk(IChunk)
}

type IChunk interface {
	AddEntity(IEntity) bool
	RemoveEntity(IEntity)
	GetIndex(int, int, int) int
	GetX() int32
	GetZ() int32
	IsLightPopulated() bool
	SetLightPopulated(bool)
	IsTerrainPopulated() bool
	SetTerrainPopulated(bool)
	GetHeight() int
	SetHeight(int)
	GetBiome(int, int) int
	SetBiome(int, int, int)
	SetBlockId(int, int, int, byte)
	GetBlockId(int, int, int) byte
	SetBlockData(int, int, int, byte)
	GetBlockData(int, int, int) byte
	SetBlockLight(int, int, int, byte)
	GetBlockLight(int, int, int) byte
	SetSkyLight(int, int, int, byte)
	GetSkyLight(int, int, int) byte
	SetSubChunk(int, ISubChunk) bool
	GetSubChunk(int) (ISubChunk, error)
	GetSubChunks() map[int]ISubChunk
	GetHighestBlockId(int, int) byte
	GetHighestBlockData(int, int) byte
	GetHighestBlock(int, int) int16
	ToBinary() []byte
	RecalculateHeightMap()
	GetEntities() map[uint64]IEntity
}

type ISubChunk interface{
	IsAllAir() bool
	GetIdIndex(int, int, int) int
	GetDataIndex(int, int, int) int
	GetBlockId(int, int, int) byte
	SetBlockId(int, int, int, byte)
	GetBlockLight(int, int, int) byte
	SetBlockLight(int, int, int, byte)
	GetSkyLight(int, int, int) byte
	SetSkyLight(int, int, int, byte)
	GetBlockData(int, int, int) byte
	SetBlockData(int, int, int, byte)
	GetHighestBlockId(int, int) byte
	GetHighestBlockData(int, int) byte
	GetHighestBlock(int, int) int
	ToBinary() []byte
}

type ILevel interface {
	GetServer() IServer
	GetName() string
	GetDimensions() map[string]IDimension
	AddDimension(IDimension)
	DimensionExists(string) bool
	RemoveDimension(string) bool
	SetDefaultDimension(IDimension)
	GetDefaultDimension() IDimension
	TickLevel()
	GetGameRules() map[string]IGameRule
	GetGameRule(string) IGameRule
	AddGameRule(IGameRule)
	GetRuntimeId() int
	GetEntityHelper() IEntityHelper
}

type IGameRule interface {
	GetName() string
	GetValue() interface{}
	SetValue(interface{}) bool
}

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
