package interfaces

type ISubChunk interface{
	IsAllAir() bool
	GetIndex(int, int, int) int
	GetBlockId(int, int, int) byte
	SetBlockId(int, int, int, byte)
	GetBlockLight(int, int, int) byte
	SetBlockLight(int, int, int, byte)
	GetSkyLight(int, int, int) byte
	SetSkyLight(int, int, int, byte)
	GetBlockData(int, int, int) byte
	SetBlockData(int, int, int, byte)
	ToBinary() []byte
}