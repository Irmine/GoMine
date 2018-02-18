package entities

import (
	"github.com/irmine/gomine/entities/math"
	"github.com/irmine/gomine/interfaces"
	"github.com/golang/geo/r3"
)

type LivingEntity struct {
	*Entity
}

func NewLivingEntity(position r3.Vector, rotation *math.Rotation, motion r3.Vector, level interfaces.ILevel, dimension interfaces.IDimension) *LivingEntity {
	return &LivingEntity{NewEntity(position, rotation, motion, level, dimension)}
}
