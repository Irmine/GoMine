package defaults

import (
	"math"

	"github.com/irmine/gomine/interfaces"
)

type Whack struct {
	*Generator
}

func NewWhackGenerator() Whack {
	return Whack{NewGenerator("Whack")}
}

// Generates and populates new chunk.

func (f Whack) GetNewChunk(chunk interfaces.IChunk) interfaces.IChunk {
	f.GenerateChunk(chunk)
	f.PopulateChunk(chunk)

	return chunk
}

func (f Whack) GenerateChunk(chunk interfaces.IChunk) {
	var y int

	for i := 0; i < 16; i++ {
		x := int(math.Cos(float64(i) * 16))
		z := int(math.Sin(float64(i) * 16))
		y = 0
		chunk.SetBlockId(i+x, y+i, i+z, 7)
		y++
		chunk.SetBlockId(i+x, y+i, i+z, 3)
		y++
		chunk.SetBlockId(i+x, y+i, i+z, 3)
		y++
		chunk.SetBlockId(i+x, y+i, i+z, 2)

		chunk.SetHeight(y)

		for i := y - 1; i >= 0; i-- {
			chunk.SetSkyLight(x, y, z, 0)
		}
	}
	chunk.RecalculateHeightMap()
}
