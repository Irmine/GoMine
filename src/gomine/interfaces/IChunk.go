package interfaces

type IChunk interface {
	AddEntity(IEntity) bool
	RemoveEntity(IEntity)
	GetIndex(int, int, int) int
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
	ToBinary() []byte
}
