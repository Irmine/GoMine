package selectors

type SelfSelector struct {
	*TargetSelector
}

func NewSelfSelector() *SelfSelector {
	return &SelfSelector{NewTargetSelector(Self)}
}
