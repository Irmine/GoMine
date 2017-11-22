package chunks

import (
	"gomine/entities"
	"errors"
	"gomine/interfaces"
	"gomine/tiles"
)

type Chunk struct {
	height int
	x, z int
	subChunks []ISubChunk
	LightPopulated bool
	TerrainPopulated bool
	tiles map[uint64]tiles.Tile
	entities map[uint64]interfaces.IEntity
	biomes [256]byte
	heightMap [4096]byte
}

func NewChunk(height, x, z int, subChunks []ISubChunk, lightPopulated, terrainPopulated bool, biomes [256]byte, heightMap [4096]byte) *Chunk {
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

func (chunk *Chunk) GetX() int {
	return chunk.x
}

func (chunk *Chunk) SetX(x int) {
	chunk.x = x
}

func (chunk *Chunk) GetZ() int {
	return chunk.z
}

func (chunk *Chunk) SetZ(z int) {
	chunk.x = z
}

func (chunk *Chunk) GetLightPopulated() bool {
	return chunk.LightPopulated
}

func (chunk *Chunk) SetLightPopulated(v bool) {
	chunk.LightPopulated = v
}

func (chunk *Chunk) GetTerrainPopulated() bool {
	return chunk.LightPopulated
}

func (chunk *Chunk) SetTerrainPopulated(v bool) {
	chunk.TerrainPopulated = v
}

func (chunk *Chunk) GetHeight() int {
	return chunk.height
}

func (chunk *Chunk) AddEntity(entity interfaces.IEntity) bool {
	if entity.IsClosed() {
		panic("Cannot add closed entity to chunk")
	}
	chunk.entities[entity.GetId()] = entity
	return true
}

func (chunk *Chunk) RemoveEntity(entity entities.Entity) {
	if k, ok := chunk.entities[entity.GetId()]; ok {
		delete(chunk.entities, k.GetId())
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
		delete(chunk.entities, k.GetId())
	}
}

func (chunk *Chunk) GetIndex(x, y, z int) int {
	return (x << 12) | (z << 8) | y
}

func (chunk *Chunk) GetHeightMapIndex(x, z int) int {
	return (z << 4) | x
}

func (chunk *Chunk) SetBlock(x, y, z int, blockId byte)  {
	v, err := chunk.GetSubChunk(y >> 4)
	if err == nil {
		v.SetBlock(x, y & 15, z, blockId)
	}
}

func (chunk *Chunk) GetBlock(x, y, z int) byte {
	v, err := chunk.GetSubChunk(y >> 4)
	if err == nil {
		return v.GetBlock(x, y & 15, z)
	}
	return 0
}

func (chunk *Chunk) SetMetadata(x, y, z int, meta byte)  {
	v, err := chunk.GetSubChunk(y >> 4)
	if err == nil {
		v.SetBlockMetadata(x, y & 15, z, meta)
	}
}

func (chunk *Chunk) GetMetadata(x, y, z int) byte {
	v, err := chunk.GetSubChunk(y >> 4)
	if err == nil {
		return v.GetBlockMetadata(x, y & 15, z)
	}
	return 0
}

func (chunk *Chunk) SetBlockLight(x, y, z int, level byte)  {
	v, err := chunk.GetSubChunk(y >> 4)
	if err == nil {
		v.SetBlockLight(x, y & 15, z, level)
	}
}

func (chunk *Chunk) GetBlockLight(x, y, z int) byte {
	v, err := chunk.GetSubChunk(y >> 4)
	if err == nil {
		return v.GetBlockLight(x, y & 15, z)
	}
	return 0
}

func (chunk *Chunk) SetSkyLight(x, y, z int, level byte)  {
	v, err := chunk.GetSubChunk(y >> 4)
	if err == nil {
		v.SetSkyLight(x, y & 15, z, level)
	}
}

func (chunk *Chunk) GetSkyLight(x, y, z int) byte {
	v, err := chunk.GetSubChunk(y >> 4)
	if err == nil {
		return v.GetSkyLight(x, y & 15, z)
	}
	return 0
}

func (chunk *Chunk) SetSubChunk(y int, subChunk ISubChunk) bool {
	if y > chunk.height || y < 0 {
		return false
	}
	chunk.subChunks[y] = subChunk
	return true
}

func (chunk *Chunk) GetSubChunk(y int) (ISubChunk, error) {
	if y > chunk.height || y < 0 {
		return NewEmptySubChunk(), errors.New("SubChunk does not exist")
	}
	return chunk.subChunks[y], nil
}

func (chunk *Chunk) GetSubChunks() []ISubChunk {
	return chunk.subChunks
}

func (chunk *Chunk) SetHeightMap(x, z int, value byte) {
	chunk.heightMap[chunk.GetHeightMapIndex(x, z)] = value
}

func (chunk *Chunk) GetHeightMap(x, z int) byte {
	return chunk.heightMap[chunk.GetHeightMapIndex(x, z)]
}

func (chunk *Chunk) PruneEmptySubChunks() {
	for y, subChunk := range chunk.subChunks {
		if y > chunk.height || y < 0 {
			chunk.subChunks = append(chunk.subChunks[:y], chunk.subChunks[y+1:]...)
			continue
		}
		if subChunk.IsAllAir() {
			chunk.subChunks[y] = NewEmptySubChunk()
		}
	}
}