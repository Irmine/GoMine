package interfaces

type IChunk interface {
	AddEntity(IEntity) bool
	RemoveEntity(entity IEntity)
	GetIndex(x, y, z int) int
	SetBlock(x, y, z int, blockId byte)
	GetBlock(x, y, z int)
	SetMetadata(x, y, z int, meta byte)
	GetMetadata(x, y, z int) byte
	SetBlockLight(x, y, z int, level byte)
	GetBlockLight(x, y, z int) byte
	SetSkyLight(x, y, z int, level byte)
	GetSkyLight(x, y, z int) byte
	SetSubChunk(y int, subChunk ISubChunk) bool
	GetSubChunk(y int) ISubChunk
	GetSubChunks() []ISubChunk
}
