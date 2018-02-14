package blocks

type Air struct {
	*Block
}

func NewAir(data byte) *Air {
	var air = &Air{NewBlock(AIR, data, "Air")}
	air.BoundingBox.Clear()
	air.CollisionBox.Clear()

	air.SetLightDiffusing(false)
	air.SetBlastResistance(0)
	air.SetHardness(0)

	return air
}
