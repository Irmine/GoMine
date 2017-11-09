package interfaces

type ILevel interface {
	GetServer() IServer
	GetName() string
	GetDimensions() map[int]IDimension
}
