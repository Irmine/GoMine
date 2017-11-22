package generation

import "gomine/vectorMath"

type Generator struct {
	name string
	spawn vectorMath.TripleVector
}

type IGenerator interface {
	GetName() string
	GetSpawn() vectorMath.TripleVector
	GenerateChunk(x, z int)
	PopulateChunk(x, z int)
}

func NewGenerator(name string) Generator {
	return Generator{name, nil}
}

func (gen *Generator) GetName() string {
	return gen.name
}

func (gen *Generator) SetSpawn(v vectorMath.TripleVector) {
	gen.spawn = v
}

func (gen *Generator) GetSpawn() vectorMath.TripleVector {
	return gen.spawn
}

func (gen *Generator) GenerateChunk() {

}

func (gen *Generator) PopulateChunk() {

}