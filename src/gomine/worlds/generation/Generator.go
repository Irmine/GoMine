package generation

import "gomine/vectors"

type Generator struct {
	name  string
	spawn vectors.TripleVector
}

type IGenerator interface {
	GetName() string
	GetSpawn() vectors.TripleVector
	GenerateChunk(x, z int)
	PopulateChunk(x, z int)
}

func NewGenerator(name string) Generator {
	return Generator{name, nil}
}

func (gen *Generator) GetName() string {
	return gen.name
}

func (gen *Generator) SetSpawn(v vectors.TripleVector) {
	gen.spawn = v
}

func (gen *Generator) GetSpawn() vectors.TripleVector {
	return gen.spawn
}

func (gen *Generator) GenerateChunk() {

}

func (gen *Generator) PopulateChunk() {

}