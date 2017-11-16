package worlds

import (
	"gomine/interfaces"
	"gomine/worlds/chunks"
)

type Level struct {
	server     interfaces.IServer
	name       string
	dimensions map[string]interfaces.IDimension
}

/**
 * Returns a new Level with the given level name.
 */
func NewLevel(levelName string, server interfaces.IServer, chunks []chunks.Chunk) *Level {
	var level = &Level{server, levelName, make(map[string]interfaces.IDimension)}
	level.AddDimension("Overworld", OverworldId, chunks)
	return level
}

/**
 * Returns the server.
 */
func (level *Level) GetServer() interfaces.IServer {
	return level.server
}

/**
 * Returns the name of this level.
 */
func (level *Level) GetName() string {
	return level.name
}

/**
 * Returns a map containing the dimensions of this level.
 * Dimension Name : Dimension
 */
func (level *Level) GetDimensions() map[string]interfaces.IDimension {
	return level.dimensions
}

/**
 * Returns whether a dimension with the given name exists on this level.
 */
func (level *Level) DimensionExists(name string) bool {
	var _, exists = level.dimensions[name]
	return exists
}

/**
 * Adds a new dimension with the given name and dimension ID.
 * Returns false if the dimension already exists, true otherwise.
 */
func (level *Level) AddDimension(name string, dimensionId int, chunks []chunks.Chunk) bool {
	if level.DimensionExists(name) {
		return false
	}
	level.dimensions[name] = NewDimension(name, dimensionId, level, chunks)
	return true
}

/**
 * Removes a dimension from this level.
 * Returns false if the dimension doesn't exist, true if it was removes successfully.
 */
func (level *Level) RemoveDimension(name string) bool {
	if !level.DimensionExists(name) {
		return false
	}
	delete(level.dimensions, name)
	return true
}

/**
 * Ticks the whole level. (All dimensions)
 * Internal. Not to be used by plugins.
 */
func (level *Level) TickLevel() {
	for _, dimension := range level.dimensions  {
		dimension.TickDimension()
	}
}