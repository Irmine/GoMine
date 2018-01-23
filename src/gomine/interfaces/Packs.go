package interfaces

type IPackStack interface {
	GetPacks() []IPack
	GetFirstPack() IPack
	AddPackOnTop(pack IPack)
	AddPackOnBottom(pack IPack)
}

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

type IPack interface {
	GetUUID() string
	GetVersion() string
	GetFileSize() int64
	GetSha256() string
	GetChunk(offset int, length int) []byte
	GetPath() string
}
