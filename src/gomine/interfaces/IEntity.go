package interfaces

type IEntity interface {
	GetNameTag() string
	SetNameTag(string)
	IsClosed() bool
	Close()
	GetHealth() int
	SetHealth(int)
	Kill()
	Tick()
	GetId() uint64
}
