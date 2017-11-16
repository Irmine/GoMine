package worlds

import (
	"gomine/interfaces"
	"gomine/worlds/chunks"
)

const (
	OverworldId = 0
	NetherId    = 1
	EndId	    = 2
)

type Dimension struct {
	name 		string
	dimensionId int
	level       interfaces.ILevel
	chunks 		[]chunks.Chunk
}

/**
 * Returns a new dimension with the given dimension ID.
 */
func NewDimension(name string, dimensionId int, level *Level, chunks []chunks.Chunk) *Dimension {
	return &Dimension{name, dimensionId, level, chunks}
}

/**
 * Returns the dimension ID of this dimension.
 */
func (dimension *Dimension) GetDimensionId() int {
	return dimension.dimensionId
}

/**
 * Returns the name of this dimension.
 */
func (dimension *Dimension) GetName() string {
	return dimension.name
}

/**
 * Returns the level this dimension is in.
 */
func (dimension *Dimension) GetLevel() interfaces.ILevel {
	return dimension.level
}

func (dimension *Dimension) GetChunk(x int, z int) (chunks.Chunk, error) {
	var chunk chunks.Chunk
	return chunk, nil
}

func (dimension *Dimension) TickDimension() {

}
