package selectors

type AllEntitiesSelector struct {
	*TargetSelector
}

func NewAllEntitiesSelector() *AllEntitiesSelector {
	return &AllEntitiesSelector{NewTargetSelector(AllEntities)}
}
