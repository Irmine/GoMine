package selectors

type AllPlayersSelector struct {
	*TargetSelector
}

func NewAllPlayersSelector() *AllPlayersSelector {
	return &AllPlayersSelector{NewTargetSelector(AllPlayers)}
}
