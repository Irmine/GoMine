package worlds

import (
	"gomine/interfaces"
)

type Level struct {
	server interfaces.IServer
	name string
	id int
	dimensions map[string]interfaces.IDimension
	defaultDimension interfaces.IDimension

	gameRules map[string]interfaces.IGameRule
}

/**
 * Returns a new Level with the given level name.
 */
func NewLevel(levelName string, levelId int, server interfaces.IServer, chunks map[int]interfaces.IChunk) *Level {
	var level = &Level{server: server, name: levelName, id: levelId, dimensions: make(map[string]interfaces.IDimension), gameRules: make(map[string]interfaces.IGameRule)}

	var defaultDimension = NewDimension("Overworld", OverworldId, level, chunks)
	level.SetDefaultDimension(defaultDimension)

	level.initializeGameRules()
	return level
}

/**
 * Returns a GameRule with the given name.
 */
func (level *Level) GetGameRule(gameRule string) interfaces.IGameRule {
	return level.gameRules[gameRule]
}

/**
 * Returns a name => GameRule map of all GameRules.
 */
func (level *Level) GetGameRules() map[string]interfaces.IGameRule {
	return level.gameRules
}

/**
 * Adds a GameRule to this level.
 */
func (level *Level) AddGameRule(rule interfaces.IGameRule) {
	level.gameRules[rule.GetName()] = rule
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
func (level *Level) AddDimension(dimension interfaces.IDimension) {
	level.dimensions[dimension.GetName()] = dimension
}

/**
 * Sets the default dimension of this level.
 */
func (level *Level) SetDefaultDimension(dimension interfaces.IDimension) {
	level.AddDimension(dimension)

	level.defaultDimension = dimension
}

/**
 * Returns the default dimension of this level.
 */
func (level *Level) GetDefaultDimension() interfaces.IDimension {
	return level.defaultDimension
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
	level.AddGameRule(NewGameRule(GameRuleCommandBlockOutput, true))
	level.AddGameRule(NewGameRule(GameRuleDoDaylightCycle, true))
	level.AddGameRule(NewGameRule(GameRuleDoEntityDrops, true))
	level.AddGameRule(NewGameRule(GameRuleDoFireTick, true))
	level.AddGameRule(NewGameRule(GameRuleDoMobLoot, true))
	level.AddGameRule(NewGameRule(GameRuleDoMobSpawning, true))
	level.AddGameRule(NewGameRule(GameRuleDoTileDrops, true))
	level.AddGameRule(NewGameRule(GameRuleDoWeatherCycle, true))
	level.AddGameRule(NewGameRule(GameRuleDrowningDamage, true))
	level.AddGameRule(NewGameRule(GameRuleFallDamage, true))
	level.AddGameRule(NewGameRule(GameRuleFireDamage, true))
	level.AddGameRule(NewGameRule(GameRuleKeepInventory, false))
	level.AddGameRule(NewGameRule(GameRuleMobGriefing, true))
	level.AddGameRule(NewGameRule(GameRuleNaturalRegeneration, true))
	level.AddGameRule(NewGameRule(GameRulePvp, true))
	level.AddGameRule(NewGameRule(GameRuleSendCommandFeedback, true))
	level.AddGameRule(NewGameRule(GameRuleShowCoordinates, true))
	level.AddGameRule(NewGameRule(GameRuleRandomTickSpeed, uint32(3)))
}