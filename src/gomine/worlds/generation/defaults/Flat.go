package defaults

import "gomine/worlds/generation"

type Flat struct {
	generation.Generator
}

func NewFlatGenerator() Flat {
	return Flat{generation.NewGenerator("Flat")}
}

func (f Flat) GeneratorChunk() {

}

func (f Flat) PopulateChunk() {

}