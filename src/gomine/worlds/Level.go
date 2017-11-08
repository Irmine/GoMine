package worlds

import "gomine"

type Level struct {
	server     gomine.Server
	name       string
	dimensions map[int]Dimension
}

/**
 * Returns a new Level with the given level name.
 */
func NewLevel(levelName string, server gomine.Server) Level {
	return Level{server, levelName, make(map[int]Dimension)}
}

/**
 * Returns the server.
 */
func (level *Level) getServer() gomine.Server {
	return level.server
}

/**
 * Returns the name of this level.
 */
func (level *Level) getName() string {
	return level.name
}

/**
 * Returns a map containing the dimensions of this level.
 */
func (level *Level) getDimension() map[int]Dimension {
	return level.dimensions
}
