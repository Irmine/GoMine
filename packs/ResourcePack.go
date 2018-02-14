package packs

type ResourcePack struct {
	*Pack
}

/**
 * Returns a new resource pack to the given path.
 */
func NewResourcePack(path string) *ResourcePack {
	return &ResourcePack{NewPack(path, Resource)}
}