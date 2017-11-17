package chunks

import (
	"gomine/entities"
	"errors"
	"gomine/interfaces"
)

type Chunk struct {
	height int
	x, z int
	SubChunks []SubChunk
	LightPopulated bool
	TerrainPopulated bool
	//tiles []tiles.Tile
	Entities map[int]interfaces.IEntity
	Biomes [256]byte
	HeightMap [4096]byte
}

func NewChunk(height, x, z int, subChunks []SubChunk, lightPopulated, terrainPopulated bool, biomes [256]byte, heightMap [4096]byte) *Chunk {
	return &Chunk{
		height,
		x,
		z,
		subChunks,
		lightPopulated,
		terrainPopulated,
		map[int]interfaces.IEntity{},
		biomes,
		heightMap,
	}
}


func (chunk *Chunk) AddEntity(entity interfaces.IEntity) bool {
	if entity.IsClosed() {
		panic("Cannot add closed entity to chunk")
	}
	chunk.Entities[entity.GetId()] = entity
	return true
}

func (chunk *Chunk) RemoveEntity(entity entities.Entity) {
	if k, ok := chunk.Entities[entity.EId]; ok {
		delete(chunk.Entities, k.GetId())
	}
}

func (chunk *Chunk) GetIndex(x, y, z int) int {
	return (x << 12) | (z << 8) | y
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
		v.SetMetadata(x, y & 15, z, meta)
	}
}

func (chunk *Chunk) GetMetadata(x, y, z int) byte {
	v, err := chunk.GetSubChunk(y >> 4)
	if err == nil {
		return v.GetMetadata(x, y & 15, z)
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

func (chunk *Chunk) SetSubChunk(y int, subChunk SubChunk) bool {
	if y > chunk.height || y < 0 {
		return false
	}
	chunk.SubChunks[y] = subChunk
	return true
}

func (chunk *Chunk) GetSubChunk(y int) (SubChunk, error) {
	if y > chunk.height || y < 0 {
		return SubChunk{}, errors.New("SubChunk does not exist")
	}
	return chunk.SubChunks[y], nil
}

func (chunk *Chunk) GetSubChunks() []SubChunk {
	return chunk.SubChunks
}