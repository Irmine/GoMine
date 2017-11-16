package chunks

import "gomine/entities"

type Chunk struct {
	height int
	x, z int
	SubChunks []ISubChunk
	LightPopulated bool
	TerrainPopulated bool
	//tiles []tiles.Tile
	Entities map[int]entities.Entity
	Biomes [256]byte
	HeightMap [4096]byte
}

func NewChunk(height, x, z int, subChunks []ISubChunk, lightPopulated, terrainPopulated bool, biomes [256]byte, heightMap [4096]byte) *Chunk {
	return &Chunk{
		height,
		x,
		z,
		subChunks,
		lightPopulated,
		terrainPopulated,
		map[int]entities.Entity{},
		biomes,
		heightMap,
	}
}


func (chunk *Chunk) AddEntity(entity entities.Entity) bool {
	if entity.Closed {
		panic("Cannot add closed entity to chunk")
	}
	chunk.Entities[entity.EId] = entity
	return true
}

func (chunk *Chunk) Remove(entity entities.Entity) {
	if k, ok := chunk.Entities[entity.EId]; ok {
		delete(chunk.Entities, k.EId)
	}
}

func (chunk *Chunk) GetIndex(x, y, z int) int {
	return (x << 12) | (z << 8) | y
}

func (chunk *Chunk) SetBlock(x, y, z int, blockId byte)  {
	v := chunk.GetSubChunk(y >> 4)
	if v != nil {
		v.SetBlock(x, y & 15, z, blockId)
	}
}

func (chunk *Chunk) GetBlock(x, y, z int) byte {
	v := chunk.GetSubChunk(y >> 4)
	if v != nil {
		return v.GetBlock(x, y & 15, z)
	}
	return 0
}

func (chunk *Chunk) SetMetadata(x, y, z int, meta byte)  {
	v := chunk.GetSubChunk(y >> 4)
	if v != nil {
		v.SetBlockMetadata(x, y & 15, z, meta)
	}
}

func (chunk *Chunk) GetMetadata(x, y, z int) byte {
	v := chunk.GetSubChunk(y >> 4)
	if v != nil {
		return v.GetBlockMetadata(x, y & 15, z)
	}
	return 0
}

func (chunk *Chunk) SetBlockLight(x, y, z int, level byte)  {
	v := chunk.GetSubChunk(y >> 4)
	if v != nil {
		v.SetBlockLight(x, y & 15, z, level)
	}
}

func (chunk *Chunk) GetBlockLight(x, y, z int) byte {
	v := chunk.GetSubChunk(y >> 4)
	if v != nil {
		return v.GetBlockLight(x, y & 15, z)
	}
	return 0
}

func (chunk *Chunk) SetSkyLight(x, y, z int, level byte)  {
	v := chunk.GetSubChunk(y >> 4)
	if v != nil {
		v.SetSkyLight(x, y & 15, z, level)
	}
}

func (chunk *Chunk) GetSkyLight(x, y, z int) byte {
	v := chunk.GetSubChunk(y >> 4)
	if v != nil {
		return v.GetSkyLight(x, y & 15, z)
	}
	return 0
}

func (chunk *Chunk) SetSubChunk(y int, subChunk ISubChunk) bool {
	if y > chunk.height || y < 0 {
		return false
	}
	chunk.SubChunks[y] = subChunk
	return true
}

func (chunk *Chunk) GetSubChunk(y int) ISubChunk {
	if y > chunk.height || y < 0 {
		return nil
	}
	return chunk.SubChunks[y]
}

func (chunk *Chunk) GetSubChunks() []ISubChunk {
	return chunk.SubChunks
}