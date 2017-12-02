package interfaces

type ILevel interface {
	GetServer() IServer
	GetName() string
	GetDimensions() map[string]IDimension
	AddDimension(string, int, map[int]IChunk) bool
	DimensionExists(string) bool
	RemoveDimension(string) bool
	TickLevel()
	ToggleGameRule(string)
	GetGameRules() map[string]bool
	SetGameRule(string, bool)
	GetGameRule(string) bool
}
