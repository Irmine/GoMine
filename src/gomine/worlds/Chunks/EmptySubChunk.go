package Chunks

type EmptySubChunk struct {
	Blocks []byte
	BlockLight []byte
	SkyLight []byte
	Metadata []byte
}

func NewEmptySubChunk() *EmptySubChunk {
	return &EmptySubChunk{[]byte{}, []byte{}, []byte{}, []byte{}}
}

func (subChunk *EmptySubChunk) IsAllAir() bool {
	return true
}

func (subChunk *EmptySubChunk) GetIndex(x, y, z int) int {
	return 0
}

func (subChunk *EmptySubChunk) GetBlock(x, y, z int) byte {
	return byte(0)
}

func (subChunk *EmptySubChunk) SetBlock(x, y, z int, data byte) {

}

func (subChunk *EmptySubChunk) GetBlockLight(x, y, z int) byte {
	return byte(0)
}

func (subChunk *EmptySubChunk) SetBlockLight(x, y, z int, data byte) {

}

func (subChunk *EmptySubChunk) GetSkyLight(x, y, z int) byte {
	return byte(0)
}

func (subChunk *EmptySubChunk) SetSkyLight(x, y, z int, data byte) {

}

func (subChunk *EmptySubChunk) GetMetadata(x, y, z int) byte {
	return byte(0)
}

func (subChunk *EmptySubChunk) SetMetadata(x, y, z int, data byte) {

}

func (subChunk *EmptySubChunk) GetBytes() []byte {
	return []byte{}
}
