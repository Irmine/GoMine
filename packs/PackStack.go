package packs

import "github.com/irmine/gomine/interfaces"

type PackStack struct {
	packs []interfaces.IPack
}

func NewPackStack() *PackStack {
	return &PackStack{[]interfaces.IPack{}}
}

/**
 * Returns all packs in the stack.
 */
func (stack *PackStack) GetPacks() []interfaces.IPack {
	return stack.packs
}

/**
 * Returns the first pack on the stack.
 */
func (stack *PackStack) GetFirstPack() interfaces.IPack {
	return stack.packs[0]
}

/**
 * Adds a pack on the top of the stack.
 */
func (stack *PackStack) AddPackOnTop(pack interfaces.IPack) {
	var newPacks = []interfaces.IPack{pack}
	stack.packs = append(newPacks, stack.packs...)
}

/**
 * Adds a pack on the bottom of the stack.
 */
func (stack *PackStack) AddPackOnBottom(pack interfaces.IPack) {
	stack.packs = append(stack.packs, pack)
}
