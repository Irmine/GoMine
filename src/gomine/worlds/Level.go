package worlds

import (
	"gomine/interfaces"
)

type Level struct {
	server     interfaces.IServer
	name       string
	dimensions map[int]interfaces.IDimension
}

/**
 * Returns a new Level with the given level name.
 */
func NewLevel(levelName string, server interfaces.IServer) *Level {
	return &Level{server, levelName, make(map[int]interfaces.IDimension)}
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
 */
func (level *Level) GetDimensions() map[int]interfaces.IDimension {
	return level.dimensions
}
