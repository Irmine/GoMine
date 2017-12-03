package full

import "gomine/worlds/blocks"

type Stone struct {
	*blocks.Block
}

const (
	StoneBlastResistance = 30
)

func NewStone(data byte) *Stone {
	var stone = &Stone{blocks.NewBlock(blocks.Stone, data, "Stone")}
	stone.SetBlastResistance(StoneBlastResistance)

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