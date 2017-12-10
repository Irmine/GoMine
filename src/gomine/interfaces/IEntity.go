package interfaces

import (
	"gomine/vectors"
	"gomine/players/math"
)

type IEntity interface {
	GetNameTag() string
	SetNameTag(string)
	IsClosed() bool
	Close()
	GetHealth() float32
	SetHealth(float32)
	Kill()
	Tick()
	GetRuntimeId() uint64
	GetPosition() vectors.TripleVector
	SetPosition(vector vectors.TripleVector)
	GetDimension() IDimension
	SetDimension(dimension IDimension)
	GetLevel() ILevel
	SetLevel(level ILevel)
	GetRotation() math.Rotation
	SetRotation(rotation math.Rotation)
	GetMotion() vectors.TripleVector
	SetMotion(vector vectors.TripleVector)
}