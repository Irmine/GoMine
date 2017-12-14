package packs

type BehaviorPack struct {
	*Pack
}

/**
 * Returns a new behaviour pack to the given path.
 */
func NewBehaviorPack(path string) *BehaviorPack {
	return &BehaviorPack{NewPack(path, Behavior)}
}