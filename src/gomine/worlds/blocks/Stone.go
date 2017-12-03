package blocks

type Stone struct {
	*Block
}

const (
	StoneBlastResistance = 30
	StoneHardness = 1.5
)

func NewStone(data byte) *Stone {
	var stone = &Stone{NewBlock(STONE, data, "Stone")}
	stone.SetBlastResistance(StoneBlastResistance)
	stone.SetHardness(StoneHardness)

	return stone
}

/**
 * Returns the name of stone adapting to its meta.
 */
func (stone *Stone) GetName() string {
	switch stone.GetData() {
	case 1:
		return "Granite"
	case 2:
		return "Polished Granite"
	case 3:
		return "Diorite"
	case 4:
		return "Polished Diorite"
	case 5:
		return "Andesite"
	case 6:
		return "Polished Andesite"
	}
	return "Stone"
}