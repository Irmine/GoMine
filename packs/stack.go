package packs

// Stack is a struct that allows ordering the stack of packs.
type Stack []Pack

// NewStack returns a new pack stack.
func NewStack() *Stack {
	return &Stack{}
}

// GetPackAtOffset returns the pack at the given offset on the stack.
func (stack *Stack) GetPackAtOffset(offset int) Pack {
	return (*stack)[offset]
}

// Pop removes the pack on top of the stack.
func (stack *Stack) Pop() {
	*stack = (*stack)[1:]
}

// Push adds the given pack on top of the stack.
func (stack *Stack) Push(pack Pack) {
	*stack = append([]Pack{pack}, *stack...)
}

// Swap swaps the packs at the given offsets with each other.
func (stack *Stack) Swap(offset1, offset2 int) {
	pack1 := (*stack)[offset1]
	pack2 := (*stack)[offset2]
	(*stack)[offset1] = pack2
	(*stack)[offset2] = pack1
}

// Len returns the length of the stack.
func (stack *Stack) Len() int {
	return len(*stack)
}

// Peek returns the top pack of the stack.
func (stack *Stack) Peek() Pack {
	return (*stack)[0]
}
