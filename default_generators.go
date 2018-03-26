package gomine

import "github.com/irmine/worlds/chunks"

type Flat struct{}

func (flat Flat) GetName() string {
	return "Flat"
}

func (flat Flat) GenerateNewChunk(x, z int32) *chunks.Chunk {
	var chunk = chunks.New(x, z)
	var y int
	for x := 0; x < 16; x++ {
		for z := 0; z < 16; z++ {
			y = 0
			chunk.SetBlockId(x, y, z, 7)
			y++
			chunk.SetBlockId(x, y, z, 3)
			y++
			chunk.SetBlockId(x, y, z, 3)
			y++
			chunk.SetBlockId(x, y, z, 2)

			for i := y - 1; i >= 0; i-- {
				chunk.SetSkyLight(x, y, z, 0)
			}
		}
	}
	chunk.RecalculateHeightMap()
	return chunk
}
