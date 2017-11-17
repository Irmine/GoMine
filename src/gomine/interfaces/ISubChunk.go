package interfaces

type ISubChunk interface{
	IsAllAir() bool
	GetIndex(x, y, z int)
	GetBlock(x, y, z int) byte
	SetBlock(x, y, z int, data byte)
	GetBlockLight(x, y, z int) byte
	SetBlockLight(x, y, z int, data byte)
	GetSkyLight(x, y, z int) byte
	SetSkyLight(x, y, z int, data byte)
	GetBlockMetadata(x, y, z int) byte
	SetBlockMetadata(x, y, z int, data byte)
}
