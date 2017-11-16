package chunks

type Chunk struct {
	height int
	x, z int
	SubChunks []ISubChunk
	LightPopulated bool
	TerrainPopulated bool
	//tiles []tiles.Tile
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
		biomes,
		heightMap,
	}
}

//write funcs
//ching chong...
//join server
//ta-da!