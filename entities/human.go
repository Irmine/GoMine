package entities

import (
	"github.com/irmine/gomine/entities/math"
	"github.com/irmine/gomine/interfaces"
	"github.com/golang/geo/r3"
)

type Human struct {
	*LivingEntity
}

func NewHuman(nameTag string, position r3.Vector, rotation *math.Rotation, motion r3.Vector, level interfaces.ILevel, dimension interfaces.IDimension) *Human {
	var human = &Human{NewLivingEntity(position, rotation, motion, level, dimension)}
	human.SetNameTag(nameTag)

	return human
}

func (human *Human) GetEntityId() uint32 {
	return Player
}
