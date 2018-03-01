package protocol

// Handler is an interface satisfied by every packet handler.
type Handler interface {
	GetPriority() int
	SetPriority(int) bool
}
