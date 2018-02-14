package selectors

const (
	NearestPlayer = "@p"
	RandomPlayer = "@r"
	AllPlayers = "@a"
	AllEntities = "@e"
	Self = "@s"
)

type TargetSelector struct {
	variable string
	arguments map[string]string
}

func NewTargetSelector(variable string) *TargetSelector {
	return &TargetSelector{variable, make(map[string]string)}
}