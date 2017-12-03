package transparent

import "gomine/worlds/blocks"

type Air struct {
	*blocks.Block
}

func NewAir(data byte) *Air {
	var air = &Air{blocks.NewBlock(blocks.Air, data, "Air")}
	air.BoundingBox.Clear()
	air.CollisionBox.Clear()

	return air
}