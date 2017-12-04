package defaults

import (
	"gomine/interfaces"
)

type Flat struct {
	*Generator
}

func NewFlatGenerator() Flat {
	return Flat{NewGenerator("Flat")}
}

/**
 * Generates and populates new chunk.
 */
func (f Flat) GetNewChunk(chunk interfaces.IChunk) interfaces.IChunk {
	f.GenerateChunk(chunk)
	f.PopulateChunk(chunk)

	return chunk
}

func (f Flat) GenerateChunk(chunk interfaces.IChunk) {
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

			chunk.SetHeight(y)

			for i := y - 1; i >= 0; i-- {
				chunk.SetSkyLight(x, y, z, 0)
			}
		}
	}
	//chunk.RecalculateHeightMap()
}