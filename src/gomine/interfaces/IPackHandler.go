package interfaces

type IPackHandler interface {
	GetResourcePacks() map[string]IPack
	GetBehaviorPacks() map[string]IPack
	LoadResourcePacks()
	GetSelectedResourcePack() IPack
	IsResourcePackLoaded(uuid string) bool
	GetResourcePack(uuid string) IPack
	GetBehaviorPackSlice() []IPack
	GetResourcePackSlice() []IPack
	IsBehaviorPackLoaded(uuid string) bool
	IsPackLoaded(uuid string) bool
	GetBehaviorPack(uuid string) IPack
	GetPack(uuid string) IPack
}
