package packs

// Stack is a struct that allows ordering the stack of packs.
type Stack struct {
	packs []Pack
}

// NewStack returns a new pack stack.
func NewStack() *Stack {
	return &Stack{[]Pack{}}
}

// GetPacks returns all packs on the stack.
func (stack *Stack) GetPacks() []Pack {
	return stack.packs
}

// GetPackAtOffset returns the pack at the given offset on the stack.
func (stack *Stack) GetPackAtOffset(offset int) Pack {
	return stack.packs[offset]
}

// Pop removes the pack on top of the stack.
func (stack *Stack) Pop() {
	stack.packs = stack.packs[1:]
}

// Push adds the given pack on top of the stack.
func (stack *Stack) Push(pack Pack) {
	var newPacks = []Pack{pack}
	stack.packs = append(newPacks, stack.packs...)
}

// Swap swaps the packs at the given offsets with each other.
func (stack *Stack) Swap(offset1, offset2 int) {
	var pack1 = stack.packs[offset1]
	var pack2 = stack.packs[offset2]
	stack.packs[offset1] = pack2
	stack.packs[offset2] = pack1
}

// Len returns the length of the stack.
func (stack *Stack) Len() int {
	return len(stack.packs)
}

// Peek returns the top pack of the stack.
func (stack *Stack) Peek() Pack {
	return stack.packs[0]
}