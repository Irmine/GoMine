package interfaces

type IPackHandler interface {
	GetResourcePacks() map[string]IPack
	GetBehaviorPacks() map[string]IPack
	LoadResourcePacks()
	LoadBehaviorPacks()
	IsResourcePackLoaded(uuid string) bool
	IsBehaviorPackLoaded(uuid string) bool
	GetResourcePack(uuid string) IPack
	GetBehaviorPack(uuid string) IPack
	IsPackLoaded(uuid string) bool
	GetPack(uuid string) IPack
	GetResourceStack() IPackStack
	GetBehaviorStack() IPackStack
}
