package defaults

import (
	"github.com/irmine/gomine/interfaces"
)

type Generator struct {
	name string
}

func NewGenerator(name string) *Generator {
	return &Generator{name: name}
}

func (gen *Generator) GetName() string {
	return gen.name
}

func (gen *Generator) GenerateChunk(interfaces.IChunk) {
}

func (gen *Generator) PopulateChunk(interfaces.IChunk) {
}
