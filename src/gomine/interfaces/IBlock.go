package interfaces

import "gomine/vectorMath"

type IBlock interface{
	GetName() string
	IsSolid() bool
	GetLightFilter() int
	GetHardness() float64
	GetBlastResistance() float64
	Place(player IPlayer, vector vectorMath.TripleVector)
}
