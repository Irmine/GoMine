package packs

// ResourcePack is a pack that modifies the visual side of game play.
type ResourcePack struct {
	*Base
}

// NewResourcePack returns a new resource pack with the given path.
func NewResourcePack(path string) *ResourcePack {
	return &ResourcePack{newBase(path, Resource)}
}
