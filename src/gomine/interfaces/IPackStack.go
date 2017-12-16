package interfaces

type IPackStack interface {
	GetPacks() []IPack
	GetFirstPack() IPack
	AddPackOnTop(pack IPack)
	AddPackOnBottom(pack IPack)
}
