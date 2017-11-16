package interfaces

import "gomine/worlds/chunks"

type ILevel interface {
	GetServer() IServer
	GetName() string
	GetDimensions() map[string]IDimension
	AddDimension(string, int, []chunks.Chunk) bool
	DimensionExists(string) bool
	RemoveDimension(string) bool
	TickLevel()
}
