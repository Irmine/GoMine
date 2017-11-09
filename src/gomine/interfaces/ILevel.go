package interfaces

type ILevel interface {
	GetServer() IServer
	GetName() string
	GetDimensions() map[string]IDimension
	AddDimension(string, int) bool
	DimensionExists(string) bool
	RemoveDimension(string) bool
	TickLevel()
}
