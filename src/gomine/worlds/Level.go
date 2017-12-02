package worlds

import (
	"gomine/interfaces"
)

type Level struct {
	server interfaces.IServer
	name string
	id int
	dimensions map[string]interfaces.IDimension
	gameRules map[string]bool
}

/**
 * Returns a new Level with the given level name.
 */
func NewLevel(levelName string, levelId int, server interfaces.IServer, chunks map[int]interfaces.IChunk) *Level {
	var level = &Level{server, levelName, levelId, make(map[string]interfaces.IDimension), make(map[string]bool)}
	level.AddDimension("Overworld", OverworldId, chunks)
	level.initializeGameRules()
	return level
}

/**
 * Returns the value of the given game rule.
 */
func (level *Level) GetGameRule(gameRule string) bool {
	return level.gameRules[gameRule]
}

/**
 * Sets a game rule to the given value.
 */
func (level *Level) SetGameRule(gameRule string, value bool) {
	level.gameRules[gameRule] = value
}

/**
 * Toggles the given game rule.
 */
func (level *Level) ToggleGameRule(gameRule string) {
	if level.gameRules[gameRule] == true {
		level.gameRules[gameRule] = false
	} else {
		level.gameRules[gameRule] = true
	}
}

/**
 * Returns a name => value map for all game rules.
 */
func (level *Level) GetGameRules() map[string]bool {
	return level.gameRules
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
 * Returns the ID of this level.
 */
func (level *Level) GetId() int {
	return level.id
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
func (level *Level) AddDimension(name string, dimensionId int, chunks map[int]interfaces.IChunk) bool {
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
 * Gets the chunk index for a certain position in a chunk
 */
func GetChunkIndex(x, z int) int {
	return (x & 429496729500) | (z & 4294967295)
}

/**
 * Gets the chunk block index for a saving changed blocks
 */
func GetBlockIndex(x, y, z int) int {
	return (x & 429496729500) << 36 | (y & 255) << 28 | (z & 4294967295)
}

/**
 * Gets the block coordinates from a chunk index
 */
func GetChunkCoordinates(index int) (int, int) {
	return index >> 32, (index & 4294967295) << 36 >> 36
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

/**
 * Initializes all game rules of the level.
 */
func (level *Level) initializeGameRules() {
	level.SetGameRule(GameRuleCommandBlockOutput, true)
	level.SetGameRule(GameRuleDoDaylightCycle, true)
	level.SetGameRule(GameRuleDoEntityDrops, true)
	level.SetGameRule(GameRuleDoFireTick, true)
	level.SetGameRule(GameRuleDoMobLoot, true)
	level.SetGameRule(GameRuleDoMobSpawning, true)
	level.SetGameRule(GameRuleDoTileDrops, true)
	level.SetGameRule(GameRuleDoWeatherCycle, true)
	level.SetGameRule(GameRuleDrowningDamage, true)
	level.SetGameRule(GameRuleFallDamage, true)
	level.SetGameRule(GameRuleFireDamage, true)
	level.SetGameRule(GameRuleKeepInventory, false)
	level.SetGameRule(GameRuleMobGriefing, true)
	level.SetGameRule(GameRuleNaturalRegeneration, true)
	level.SetGameRule(GameRulePvp, true)
	level.SetGameRule(GameRuleSendCommandFeedback, true)
}