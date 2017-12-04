package defaults

import (
	"gomine/worlds/generation"
	"gomine/worlds/chunks"
)

type Flat struct {
	*generation.Generator
}

func NewFlatGenerator() Flat {
	return Flat{generation.NewGenerator("Flat")}
}

func (f Flat) GeneratorChunk(x, z int) {
	f.Chunk = chunks.NewChunk(x, z)
	f.Level.GetDefaultDimension().SetChunk(x, z, f.Chunk)
	f.PopulateChunk()
}

func (f Flat) PopulateChunk() {
	var y int
	for x := -16; x < 16; x++ {
		for z := -16; z < 16; z++ {
			y = 0
			y++
			f.Chunk.SetBlockId(x, y, z, 7)
			y++
			f.Chunk.SetBlockId(x, y, z, 3)
			y++
			f.Chunk.SetBlockId(x, y, z, 3)
			y++
			f.Chunk.SetBlockId(x, y, z, 2)
			f.Chunk.SetHeight(y)
			for i := y - 1; i >= 0; i-- {
				f.Chunk.SetSkyLight(x, y, z, 0)
			}
			f.Chunk.SetBiome(x, z, 1)
		}
	}
}