package chunks

type ISubChunk interface{
	IsAllAir() bool
	GetIndex(x, y, z int)
	GetBlock(x, y, z int) byte
	SetBlock(x, y, z int, data byte)
	GetBlockLight(x, y, z int) byte
	SetBlockLight(x, y, z int, data byte)
	GetBlockMetadata(x, y, z int) byte
	SetBlockMetadata(x, y, z int, data byte)
}

type SubChunk struct {
	Blocks []byte
	BlockLight []byte
	SkyLight []byte
	Metadata []byte
}

func NewSubChunk() *SubChunk {
	return &SubChunk{[]byte{}, []byte{}, []byte{}, []byte{}}
}

func (subChunk *SubChunk) IsAllAir() bool {
	var isAir = true
	for _, v := range subChunk.Blocks {
		if v != 0x00 {
			isAir = false
		}
	}
	return isAir
}

func (subChunk *SubChunk) GetIndex(x, y, z int) int {
	return (x << 8) | (z << 4) | y
}

func (subChunk *SubChunk) GetBlock(x, y, z int) byte {
	return subChunk.Blocks[subChunk.GetIndex(x, y, z)]
}

func (subChunk *SubChunk) SetBlock(x, y, z int, data byte) {
	subChunk.Blocks[subChunk.GetIndex(x, y, z)] = data
}

func (subChunk *SubChunk) GetBlockLight(x, y, z int) byte {
	return subChunk.BlockLight[subChunk.GetIndex(x, y, z)]
}

func (subChunk *SubChunk) SetBlockLight(x, y, z int, data byte) {
	subChunk.BlockLight[subChunk.GetIndex(x, y, z)] = data
}

func (subChunk *SubChunk) GetSkyLight(x, y, z int) byte {
	return subChunk.SkyLight[subChunk.GetIndex(x, y, z)]
}

func (subChunk *SubChunk) SetSkyLight(x, y, z int, data byte) {
	subChunk.SkyLight[subChunk.GetIndex(x, y, z)] = data
}

func (subChunk *SubChunk) GetMetadata(x, y, z int) byte {
	return subChunk.Metadata[subChunk.GetIndex(x, y, z)]
}

func (subChunk *SubChunk) SetMetadata(x, y, z int, data byte) {
	subChunk.Metadata[subChunk.GetIndex(x, y, z)] = data
}

func (subChunk *SubChunk) GetBytes() []byte {
	return []byte{}
}