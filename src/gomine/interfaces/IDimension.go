package interfaces

type IDimension interface {
	GetDimensionId() int
	GetLevel() ILevel
	GetName() string
	TickDimension()
}
