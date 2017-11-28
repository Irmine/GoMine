package interfaces

type ILevel interface {
	GetServer() IServer
	GetName() string
	GetDimensions() map[string]IDimension
	AddDimension(string, int, []IChunk) bool
	DimensionExists(string) bool
	RemoveDimension(string) bool
	TickLevel()
	ToggleGameRule(gameRule string)
	GetGameRules() map[string]bool
	SetGameRule(gameRule string, value bool)
	GetGameRule(gameRule string) bool
}
