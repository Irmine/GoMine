package chunks

import (
	"errors"
	"gomine/interfaces"
	"gomine/tiles"
	"gomine/utils"
)

type Chunk struct {
	height int
	x, z int32
	subChunks map[int]interfaces.ISubChunk
	LightPopulated bool
	TerrainPopulated bool
	tiles map[uint64]tiles.Tile
	entities map[uint64]interfaces.IEntity
	biomes map[int]int
	heightMap [257]int16
}

func NewChunk(x, z int32) *Chunk {
	return &Chunk{
		256,
		x,
		z,
		make(map[int]interfaces.ISubChunk),
		true,
		true,
		make(map[uint64]tiles.Tile),
		make(map[uint64]interfaces.IEntity),
		make(map[int]int),
		[257]int16{},
	}
}

/**
 * Returns the chunk X position.
 */
func (chunk *Chunk) GetX() int32 {
	return chunk.x
}

/**
 * Returns the chunk Z position.
 */
func (chunk *Chunk) GetZ() int32 {
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
 * Sets the height of this chunk. (?)
 */
func (chunk *Chunk) SetHeight(height int) {
	chunk.height = height
}

/**
 * Returns the biome of this coordinate. (?)
 */
func (chunk *Chunk) GetBiome(x, z int) int {
	return chunk.biomes[chunk.GetBiomeIndex(x, z)]
}

/**
 * Sets the biome of this coordinate. (?)
 */
func (chunk *Chunk) SetBiome(x, z, biome int) {
	chunk.biomes[chunk.GetBiomeIndex(x, z)] = biome
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
 * Returns the biome index of a coordinate in a chunk.
 */
func (chunk *Chunk) GetBiomeIndex(x, z int) int {
	return (x << 4) | z
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
	} else {
		sub := NewSubChunk()
		sub.SetBlockId(x, y & 15, z, blockId)
		chunk.SetSubChunk(y >> 4, sub)
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
	if _, ok := chunk.subChunks[y]; ok {
		return chunk.subChunks[y], nil
	}
	chunk.subChunks[y] = NewSubChunk()
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
func (chunk *Chunk) SetHeightMap(x, z int, value int16) {
	chunk.heightMap[chunk.GetHeightMapIndex(x, z)] = value
}

/**
 * Returns the height in the HeightMap on the given index.
 */
func (chunk *Chunk) GetHeightMap(x, z int) int16 {
	return chunk.heightMap[chunk.GetHeightMapIndex(x, z)]
}

/**
 * Recalculates HeightMap (highest blocks) of the chunk
 */
func (chunk *Chunk) RecalculateHeightMap() {
	for x := 0; x < 16; x++ {
		for z := 0; z < 16; z++ {

			id := int(chunk.GetHighestBlockId(x, z))

			if GetLightFilter(id) > 0 && !DiffusesLight(id) {
				break
			}

			chunk.SetHeightMap(x, z, chunk.GetHighestBlock(x, z) + 1)
		}
	}
}

/**
 * Returns highest SubChunk in this chunk
 */
func (chunk *Chunk) GetHighestSubChunk() interfaces.ISubChunk {
	var highest interfaces.ISubChunk = NewEmptySubChunk()
	for y := 15; y >= 0; y-- {
		if _, ok := chunk.subChunks[y];! ok {
			continue
		}
		if chunk.subChunks[y].IsAllAir() {
			continue
		}
		highest = chunk.subChunks[y]
		break
	}
	return highest
}

/**
 * Returns highest block id at certain x, z coordinates in this chunk
 */
func (chunk *Chunk) GetHighestBlockId(x, z int) byte {
	return chunk.GetHighestSubChunk().GetHighestBlockId(x, z)
}

/**
 * Returns highest block meta data at certain x, z coordinates in this chunk
 */
func (chunk *Chunk) GetHighestBlockData(x, z int) byte {
	 return chunk.GetHighestSubChunk().GetHighestBlockData(x, z)
}

/**
 * Returns highest light filtering block at certain x, z coordinates in this chunk
 */
func (chunk *Chunk) GetHighestBlock(x, z int) int16 {
	return int16(chunk.GetHighestSubChunk().GetHighestBlock(x, z))
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
	var subChunkCount = chunk.GetFilledSubChunks()

	stream.PutByte(subChunkCount)
	for i := 0; i < int(subChunkCount); i++ {
		stream.PutBytes(chunk.subChunks[i].ToBinary())
	}

	for i := 256; i >= 0; i-- {
		stream.PutShort(chunk.heightMap[i])
	}

	for _, biome := range chunk.biomes {
		stream.PutByte(byte(biome))
	}
	stream.PutByte(0)

	stream.PutUnsignedVarInt(0)

	return stream.GetBuffer()
}