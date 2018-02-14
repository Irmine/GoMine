package selectors

type RandomPlayerSelector struct {
	*TargetSelector
}

func NewRandomPlayerSelector() *RandomPlayerSelector {
	return &RandomPlayerSelector{NewTargetSelector(RandomPlayer)}
}
