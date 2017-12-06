package entities

type Human struct {
	*LivingEntity
}

func NewHuman(nameTag string) *Human {
	return &Human{}
}