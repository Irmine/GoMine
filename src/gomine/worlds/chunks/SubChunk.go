package chunks

type SubChunk struct {
	BlockIds []byte
	BlockData []byte
	BlockLight []byte
	SkyLight []byte
}

func NewSubChunk() *SubChunk {
	return &SubChunk{make([]byte, 4096), make([]byte, 2048), make([]byte, 2048), make([]byte, 2048)}
}

/**
 * Checks if this SubChunk is completely empty.
 */
func (subChunk *SubChunk) IsAllAir() bool {
	return string(subChunk.BlockIds) == string(make([]byte, 4096))
}

/**
 * Returns the index of the given xyz values for IDs in the SubChunk.
 */
func (subChunk *SubChunk) GetIdIndex(x, y, z int) int {
	return (x << 8) | (z << 4) | y
}

/**
 * Returns the index of the given xyz values for data in the SubChunk.
 */
func (subChunk *SubChunk) GetDataIndex(x, y, z int) int {
	return (x << 7) + (z << 3) + (y >> 1)
}

/**
 * Returns the block ID in the SubChunk at the given position.
 */
func (subChunk *SubChunk) GetBlockId(x, y, z int) byte {
	return subChunk.BlockIds[subChunk.GetIdIndex(x, y, z)]
}

/**
 * Sets the block ID in the SubChunk at the given position.
 */
func (subChunk *SubChunk) SetBlockId(x, y, z int, id byte) {
	subChunk.BlockIds[subChunk.GetIdIndex(x, y, z)] = id
}

/**
 * Returns the block light in the SubChunk at the given position.
 */
func (subChunk *SubChunk) GetBlockLight(x, y, z int) byte {
	var data = subChunk.BlockLight[subChunk.GetDataIndex(x, y, z)]
	if (y & 0x01) == 0 {
		return data & 0x0f
	}
	return data >> 4
}

/**
 * Sets the block light in the SubChunk at the given position.
 */
func (subChunk *SubChunk) SetBlockLight(x, y, z int, data byte) {
	var i = subChunk.GetDataIndex(x, y, z)
	var d = subChunk.BlockLight[i]
	if (y & 0x01) == 0 {
		subChunk.BlockLight[i] = (d & 0xf0) | (data & 0x0f)
		return
	}
	subChunk.BlockLight[i] = ((data & 0x0f) << 4) | (d & 0x0f)
}

/**
 * Returns the sky light in the SubChunk at the given position.
 */
func (subChunk *SubChunk) GetSkyLight(x, y, z int) byte {
	var data = subChunk.SkyLight[subChunk.GetDataIndex(x, y, z)]
	if (y & 0x01) == 0 {
		return data & 0x0f
	}
	return data >> 4
}

/**
 * Sets the sky light in the SubChunk at the given position.
 */
func (subChunk *SubChunk) SetSkyLight(x, y, z int, data byte) {
	var i = subChunk.GetDataIndex(x, y, z)
	var d = subChunk.SkyLight[i]
	if (y & 0x01) == 0 {
		subChunk.SkyLight[i] = (d & 0xf0) | (data & 0x0f)
		return
	}
	subChunk.SkyLight[i] = ((data & 0x0f) << 4) | (d & 0x0f)
}

/**
 * Returns the block data of a block in the SubChunk on the given position.
 */
func (subChunk *SubChunk) GetBlockData(x, y, z int) byte {
	var data = subChunk.BlockData[subChunk.GetDataIndex(x, y, z)]
	if (y & 0x01) == 0 {
		return data & 0x0f
	}
	return data >> 4
}

/**
 * Sets the block data of a block in the SubChunk on the given position.
 */
func (subChunk *SubChunk) SetBlockData(x, y, z int, data byte) {
	var i = subChunk.GetDataIndex(x, y, z)
	var d = subChunk.BlockData[i]
	if (y & 0x01) == 0 {
		subChunk.BlockData[i] = (d & 0xf0) | (data & 0x0f)
		return
	}
	subChunk.BlockData[i] = ((data & 0x0f) << 4) | (d & 0x0f)
}

/**
 * Returns highest block id at certain x, z coordinates in this subchunk
 */
func (subChunk *SubChunk) GetHighestBlockId(x, z int) byte {
	var id byte

	for y := 15; y >= 0; y-- {
		id = subChunk.GetBlockId(x, y, z)
		if id != 0 {
			return id
		}
	}

	return 0
}

/**
 * Returns block meta data at certain x, z coordinates in this subchunk
 */
func (subChunk *SubChunk) GetHighestBlockData(x, z int) byte {
	for y := 15; y >= 0; y-- {
		return subChunk.GetBlockData(x, y, z)
	}

	return 0
}

/**
 * Returns highest light filtering at certain x, z coordinates in this subchunk
 */
func (subChunk *SubChunk) GetHighestBlock(x, z int) int {
	for y := 15; y >= 0; y-- {
		if subChunk.GetBlockId(x, y, z) != 0 {
			return y
		}
	}

	return 0
}

/**
 * Converts the sub chunk into binary.
 */
func (subChunk *SubChunk) ToBinary() []byte {
	var bytes = []byte{00}
	bytes = append(bytes, subChunk.BlockIds...)
	bytes = append(bytes, subChunk.BlockData...)
	return bytes
}