package entities

import (
	"gomine/vectors"
	"gomine/interfaces"
	"gomine/entities/math"
)

type LivingEntity struct {
	*Entity
}

func NewLivingEntity(position vectors.TripleVector, rotation math.Rotation, motion vectors.TripleVector, level interfaces.ILevel, dimension interfaces.IDimension) *LivingEntity {
	return &LivingEntity{NewEntity(position, rotation, motion, level, dimension)}
}