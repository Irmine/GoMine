package entities

type Human struct {
	*Entity
}

func NewHuman(nameTag string) *Human {
	return &Human{}
}
