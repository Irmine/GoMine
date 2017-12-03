package worlds

import "reflect"

const (
	GameRuleCommandBlockOutput = "commandblockoutput"
	GameRuleDoDaylightCycle = "dodaylightcycle"
	GameRuleDoEntityDrops = "doentitydrops"
	GameRuleDoFireTick = "dofiretick"
	GameRuleDoMobLoot = "domobloot"
	GameRuleDoMobSpawning = "domobspawning"
	GameRuleDoTileDrops = "dotiledrops"
	GameRuleDoWeatherCycle = "doweathercycle"
	GameRuleDrowningDamage = "drowningdamage"
	GameRuleFallDamage = "falldamage"
	GameRuleFireDamage = "firedamage"
	GameRuleKeepInventory = "keepinventory"
	GameRuleMobGriefing = "mobgriefing"
	GameRuleNaturalRegeneration = "naturalregeneration"
	GameRulePvp = "pvp"
	GameRuleSendCommandFeedback = "sendcommandfeedback"
	GameRuleShowCoordinates = "showcoordinates"
	GameRuleRandomTickSpeed = "randomtickspeed"
)

type GameRule struct {
	name string
	value interface{}
}

/**
 * Returns a new GameRule with the given name and value.
 */
func NewGameRule(name string, value interface{}) *GameRule {
	return &GameRule{name, value}
}

/**
 * Returns the name of this GameRule.
 */
func (rule *GameRule) GetName() string {
	return rule.name
}

/**
 * Returns the value this GameRule holds.
 * Either an uint32, bool or float32.
 */
func (rule *GameRule) GetValue() interface{} {
	return rule.value
}

/**
 * Sets the value of this GameRule.
 * Returns false if the value is not valid for this GameRule.
 */
func (rule *GameRule) SetValue(value interface{}) bool {
	if reflect.TypeOf(value).Kind() != reflect.TypeOf(rule.value).Kind() {
		return false
	}
	rule.value = value
	return true
}