package entities

import (
	"github.com/irmine/gomine/entities/math"
	"github.com/irmine/gomine/interfaces"
	"github.com/irmine/gomine/vectors"
)

type Human struct {
	*LivingEntity
}

func NewHuman(nameTag string, position *vectors.TripleVector, rotation *math.Rotation, motion *vectors.TripleVector, level interfaces.ILevel, dimension interfaces.IDimension) *Human {
	var human = &Human{NewLivingEntity(position, rotation, motion, level, dimension)}
	human.SetNameTag(nameTag)

	return human
}

func (human *Human) GetEntityId() uint32 {
	return Player
}
