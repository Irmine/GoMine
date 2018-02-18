package entities

import (
	"github.com/irmine/gomine/entities/math"
	"github.com/irmine/gomine/interfaces"
	"github.com/irmine/gomine/vectors"
)

type LivingEntity struct {
	*Entity
}

func NewLivingEntity(position *vectors.TripleVector, rotation *math.Rotation, motion *vectors.TripleVector, level interfaces.ILevel, dimension interfaces.IDimension) *LivingEntity {
	return &LivingEntity{NewEntity(position, rotation, motion, level, dimension)}
}
