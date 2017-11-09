package worlds

import "gomine/interfaces"

const (
	OverworldId = 0
	NetherId    = 1
	EndId	    = 2
)

type Dimension struct {
	name 		string
	dimensionId int
	level       interfaces.ILevel
}

/**
 * Returns a new dimension with the given dimension ID.
 */
func NewDimension(name string, dimensionId int, level *Level) *Dimension {
	return &Dimension{name, dimensionId, level}
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

func (dimension *Dimension) TickDimension() {

}
