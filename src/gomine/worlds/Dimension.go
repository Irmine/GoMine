package worlds

const (
	Overworld = 0
	Nether    = 1
	End       = 2
)

type Dimension struct {
	dimensionId int
	level       Level
}

/**
 * Returns a new dimension with the given dimension ID.
 */
func NewDimension(level Level, dimensionId int) Dimension {
	return Dimension{dimensionId, level}
}

/**
 * Returns the dimension ID of this dimension.
 */
func (dimension *Dimension) getDimensionId() int {
	return dimension.dimensionId
}

/**
 * Returns the level this dimension is in.
 */
func (dimension *Dimension) getLevel() Level {
	return dimension.level
}
