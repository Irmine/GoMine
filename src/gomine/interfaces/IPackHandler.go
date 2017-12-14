package interfaces

type IPackHandler interface {
	GetResourcePacks() map[string]IPack
	GetBehaviorPacks() map[string]IPack
	LoadResourcePacks()
	LoadBehaviorPacks()
	GetSelectedResourcePack() IPack
	GetSelectedBehaviorPack() IPack
	IsResourcePackLoaded(uuid string) bool
	IsBehaviorPackLoaded(uuid string) bool
	GetResourcePack(uuid string) IPack
	GetBehaviorPack(uuid string) IPack
	GetBehaviorPackSlice() []IPack
	GetResourcePackSlice() []IPack
	IsPackLoaded(uuid string) bool
	GetPack(uuid string) IPack
}
