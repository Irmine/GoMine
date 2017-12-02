package interfaces

type ILevel interface {
	GetServer() IServer
	GetName() string
	GetDimensions() map[string]IDimension
	AddDimension(IDimension)
	DimensionExists(string) bool
	RemoveDimension(string) bool
	SetDefaultDimension(IDimension)
	GetDefaultDimension() IDimension
	TickLevel()
	ToggleGameRule(string)
	GetGameRules() map[string]bool
	SetGameRule(string, bool)
	GetGameRule(string) bool
}
