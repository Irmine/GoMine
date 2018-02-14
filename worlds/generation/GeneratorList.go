package generation

import (
	"github.com/irmine/gomine/interfaces"
	"github.com/irmine/gomine/worlds/generation/defaults"
)

var list = map[string]interfaces.IGenerator{}

func init() {
	RegisterGenerator(defaults.NewFlatGenerator())
	RegisterGenerator(defaults.NewWhackGenerator())
}

func RegisterGenerator(generator interfaces.IGenerator) {
	list[generator.GetName()] = generator
}

func GeneratorExists(generator interfaces.IGenerator) bool {
	var _, ok = list[generator.GetName()]
	return ok
}

func GeneratorNameExists(generator string) bool {
	var _, ok = list[generator]
	return ok
}

func DeRegisterGenerator(generator interfaces.IGenerator) {
	if GeneratorExists(generator) {
		delete(list, generator.GetName())
	}
}

func DeRegisterGeneratorByName(generator string) {
	if GeneratorNameExists(generator) {
		delete(list, generator)
	}
}

func GetGeneratorByName(generator string) interfaces.IGenerator {
	return list[generator]
}
