package chunks

type SubChunk struct {
	BlockIds []byte
	BlockData []byte
	BlockLight []byte
	SkyLight []byte
}

func NewSubChunk() *SubChunk {
	return &SubChunk{make([]byte, 4096), make([]byte, 4096), make([]byte, 4096), make([]byte, 4096)}
}

/**
 * Checks if this SubChunk is completely empty.
 */
func (subChunk *SubChunk) IsAllAir() bool {
	var isAir = true
	for _, v := range subChunk.BlockIds {
		if v != 0x00 {
			isAir = false
		}
	}
	return isAir
}

/**
 * Returns the index of the given xyz values in the SubChunk.
 */
func (subChunk *SubChunk) GetIndex(x, y, z int) int {
	return (x << 8) | (z << 4) | y
}

/**
 * Returns the block ID in the SubChunk at the given position.
 */
func (subChunk *SubChunk) GetBlockId(x, y, z int) byte {
	return subChunk.BlockIds[subChunk.GetIndex(x, y, z)]
}

/**
 * Sets the block ID in the SubChunk at the given position.
 */
func (subChunk *SubChunk) SetBlockId(x, y, z int, id byte) {
	subChunk.BlockIds[subChunk.GetIndex(x, y, z)] = id
}

/**
 * Returns the block light in the SubChunk at the given position.
 */
func (subChunk *SubChunk) GetBlockLight(x, y, z int) byte {
	return subChunk.BlockLight[subChunk.GetIndex(x, y, z)]
}

/**
 * Sets the block light in the SubChunk at the given position.
 */
func (subChunk *SubChunk) SetBlockLight(x, y, z int, data byte) {
	subChunk.BlockLight[subChunk.GetIndex(x, y, z)] = data
}

/**
 * Returns the sky light in the SubChunk at the given position.
 */
func (subChunk *SubChunk) GetSkyLight(x, y, z int) byte {
	return subChunk.SkyLight[subChunk.GetIndex(x, y, z)]
}

/**
 * Sets the sky light in the SubChunk at the given position.
 */
func (subChunk *SubChunk) SetSkyLight(x, y, z int, data byte) {
	subChunk.SkyLight[subChunk.GetIndex(x, y, z)] = data
}

/**
 * Returns the block data of a block in the SubChunk on the given position.
 */
func (subChunk *SubChunk) GetBlockData(x, y, z int) byte {
	return subChunk.BlockData[subChunk.GetIndex(x, y, z)]
}

/**
 * Sets the block data of a block in the SubChunk on the given position.
 */
func (subChunk *SubChunk) SetBlockData(x, y, z int, data byte) {
	subChunk.BlockData[subChunk.GetIndex(x, y, z)] = data
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