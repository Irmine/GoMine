package entities

import (
	"gomine/vectors"
	"gomine/interfaces"
	"gomine/entities/math"
)

type Human struct {
	*LivingEntity
}

func NewHuman(nameTag string, position *vectors.TripleVector, rotation *math.Rotation, motion *vectors.TripleVector, level interfaces.ILevel, dimension interfaces.IDimension) *Human {
	var human = &Human{NewLivingEntity(position, rotation, motion, level, dimension)}
	human.SetNameTag(nameTag)

	return human
}