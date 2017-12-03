package chunks

import (
	"errors"
	"gomine/interfaces"
	"gomine/tiles"
	"gomine/utils"
)

type Chunk struct {
	height int
	x, z int
	subChunks map[int]interfaces.ISubChunk
	LightPopulated bool
	TerrainPopulated bool
	tiles map[uint64]tiles.Tile
	entities map[uint64]interfaces.IEntity
	biomes [256]byte
	heightMap [4096]byte
}

func NewChunk(height, x, z int, subChunks map[int]interfaces.ISubChunk, lightPopulated, terrainPopulated bool, biomes [256]byte, heightMap [4096]byte) *Chunk {
	return &Chunk{
		height,
		x,
		z,
		subChunks,
		lightPopulated,
		terrainPopulated,
		map[uint64]tiles.Tile{},
		map[uint64]interfaces.IEntity{},
		biomes,
		heightMap,
	}
}

/**
 * Returns the chunk X position.
 */
func (chunk *Chunk) GetX() int {
	return chunk.x
}

/**
 * Returns the chunk Z position.
 */
func (chunk *Chunk) GetZ() int {
	return chunk.z
}

/**
 * Returns if this chunk is light populated.
 */
func (chunk *Chunk) IsLightPopulated() bool {
	return chunk.LightPopulated
}

/**
 * Sets the chunk light populated.
 */
func (chunk *Chunk) SetLightPopulated(v bool) {
	chunk.LightPopulated = v
}

/**
 * Returns if this chunk is terrain populated.
 */
func (chunk *Chunk) IsTerrainPopulated() bool {
	return chunk.LightPopulated
}

/**
 * Sets this chunk terrain populated.
 */
func (chunk *Chunk) SetTerrainPopulated(v bool) {
	chunk.TerrainPopulated = v
}

/**
 * Returns the height of this chunk. (?)
 */
func (chunk *Chunk) GetHeight() int {
	return chunk.height
}

/**
 * Adds a new entity to this chunk.
 */
func (chunk *Chunk) AddEntity(entity interfaces.IEntity) bool {
	if entity.IsClosed() {
		panic("Cannot add closed entity to chunk")
	}
	chunk.entities[entity.GetRuntimeId()] = entity
	return true
}

/**
 * Removes an entity from this chunk.
 */
func (chunk *Chunk) RemoveEntity(entity interfaces.IEntity) {
	if k, ok := chunk.entities[entity.GetRuntimeId()]; ok {
		delete(chunk.entities, k.GetRuntimeId())
	}
}

func (chunk *Chunk) AddTile(tile tiles.Tile) bool {
	if tile.IsClosed() {
		panic("Cannot add closed entity to chunk")
	}
	chunk.tiles[tile.GetId()] = tile
	return true
}

func (chunk *Chunk) RemoveTile(tile tiles.Tile) {
	if k, ok := chunk.entities[tile.GetId()]; ok {
		delete(chunk.entities, k.GetRuntimeId())
	}
}

/**
 * Returns the index of a position in a chunk.
 */
func (chunk *Chunk) GetIndex(x, y, z int) int {
	return (x << 12) | (z << 8) | y
}

/**
 * Returns the index of a position in the HeightMap of this chunk.
 */
func (chunk *Chunk) GetHeightMapIndex(x, z int) int {
	return (z << 4) | x
}

/**
 * Sets the block ID on a position in this chunk.
 */
func (chunk *Chunk) SetBlockId(x, y, z int, blockId byte)  {
	v, err := chunk.GetSubChunk(y >> 4)
	if err == nil {
		v.SetBlockId(x, y & 15, z, blockId)
	}
}

/**
 * Returns the block ID on a position in this chunk.
 */
func (chunk *Chunk) GetBlockId(x, y, z int) byte {
	v, err := chunk.GetSubChunk(y >> 4)
	if err == nil {
		return v.GetBlockId(x, y & 15, z)
	}
	return 0
}

/**
 * Sets the block data on a position in this chunk.
 */
func (chunk *Chunk) SetBlockData(x, y, z int, data byte)  {
	v, err := chunk.GetSubChunk(y >> 4)
	if err == nil {
		v.SetBlockData(x, y & 15, z, data)
	}
}

/**
 * Returns the block data on a position in this chunk.
 */
func (chunk *Chunk) GetBlockData(x, y, z int) byte {
	v, err := chunk.GetSubChunk(y >> 4)
	if err == nil {
		return v.GetBlockData(x, y & 15, z)
	}
	return 0
}

/**
 * Sets the block light on a position in this chunk.
 */
func (chunk *Chunk) SetBlockLight(x, y, z int, level byte)  {
	v, err := chunk.GetSubChunk(y >> 4)
	if err == nil {
		v.SetBlockLight(x, y & 15, z, level)
	}
}

/**
 * Returns the block light on a position in this chunk.
 */
func (chunk *Chunk) GetBlockLight(x, y, z int) byte {
	v, err := chunk.GetSubChunk(y >> 4)
	if err == nil {
		return v.GetBlockLight(x, y & 15, z)
	}
	return 0
}

/**
 * Sets the sky light on a position in this chunk.
 */
func (chunk *Chunk) SetSkyLight(x, y, z int, level byte)  {
	v, err := chunk.GetSubChunk(y >> 4)
	if err == nil {
		v.SetSkyLight(x, y & 15, z, level)
	}
}

/**
 * Returns the sky light on a position in this chunk.
 */
func (chunk *Chunk) GetSkyLight(x, y, z int) byte {
	v, err := chunk.GetSubChunk(y >> 4)
	if err == nil {
		return v.GetSkyLight(x, y & 15, z)
	}
	return 0
}

/**
 * Sets a SubChunk on a position in this chunk.
 */
func (chunk *Chunk) SetSubChunk(y int, subChunk interfaces.ISubChunk) bool {
	if y > chunk.height || y < 0 {
		return false
	}
	chunk.subChunks[y] = subChunk
	return true
}

/**
 * Returns a SubChunk on a given height index in this chunk.
 */
func (chunk *Chunk) GetSubChunk(y int) (interfaces.ISubChunk, error) {
	if y > chunk.height || y < 0 {
		return NewEmptySubChunk(), errors.New("SubChunk does not exist")
	}
	return chunk.subChunks[y], nil
}

/**
 * Returns all SubChunks in this chunk.
 */
func (chunk *Chunk) GetSubChunks() map[int]interfaces.ISubChunk {
	return chunk.subChunks
}

/**
 * Sets the HeightMap of this chunk.
 */
func (chunk *Chunk) SetHeightMap(x, z int, value byte) {
	chunk.heightMap[chunk.GetHeightMapIndex(x, z)] = value
}

/**
 * Returns the height in the HeightMap on the given index.
 */
func (chunk *Chunk) GetHeightMap(x, z int) byte {
	return chunk.heightMap[chunk.GetHeightMapIndex(x, z)]
}

/**
 * Returns the count of non-empty SubChunks in this chunk.
 */
func (chunk *Chunk) GetFilledSubChunks() byte {
	chunk.PruneEmptySubChunks()
	return byte(len(chunk.subChunks))
}

/**
 * Prunes all empty SubChunks in this chunk.
 */
func (chunk *Chunk) PruneEmptySubChunks() {
	for y, subChunk := range chunk.subChunks {
		if y > chunk.height || y < 0 {
			delete(chunk.subChunks, y)
			continue
		}
		if subChunk.IsAllAir() {
			chunk.subChunks[y] = NewEmptySubChunk()
		}
	}
}

/**
 * Converts the chunk to binary preparing it to send to the client.
 */
func (chunk *Chunk) ToBinary() []byte {
	var stream = utils.NewStream()
	stream.ResetStream()
	var subChunkCount = chunk.GetFilledSubChunks()

	stream.PutByte(subChunkCount)
	for i := 0; i < int(subChunkCount); i++ {
		stream.PutBytes(chunk.subChunks[i].ToBinary())
	}

	for i := 4095; i >= 0; i-- {
		stream.PutByte(1)
	}

	for _, biome := range chunk.biomes {
		stream.PutByte(biome)
	}
	stream.PutByte(0)

	stream.PutUnsignedVarInt(0)

	return stream.GetBuffer()
}