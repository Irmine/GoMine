package selectors

type NearestPlayerSelector struct {
	*TargetSelector
}

func NewNearestPlayerSelector() *NearestPlayerSelector {
	return &NearestPlayerSelector{NewTargetSelector(NearestPlayer)}
}