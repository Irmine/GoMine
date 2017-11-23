package interfaces

type ISubChunk interface{
	IsAllAir() bool
	GetIndex(int, int, int) int
	GetBlockId(int, int, int) int
	SetBlockId(int, int, int, int)
	GetBlockLight(int, int, int) byte
	SetBlockLight(int, int, int, byte)
	GetSkyLight(int, int, int) byte
	SetSkyLight(int, int, int, byte)
	GetBlockData(int, int, int) byte
	SetBlockData(int, int, int, byte)
}