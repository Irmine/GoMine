package generation

var list map[string]IGenerator

func RegisterGenerator(generator IGenerator)  {
	list[generator.GetName()] = generator
}

func GeneratorExists(generator IGenerator) bool {
	var _, ok = list[generator.GetName()]
	return ok
}

func GeneratorExistsN(generator string) bool {
	var _, ok = list[generator]
	return ok
}

func DeRegisterGenerator(generator IGenerator)  {
	if GeneratorExists(generator) {
		delete(list, generator.GetName())
	}
}

func DeRegisterGeneratorN(generator string)  {
	if GeneratorExistsN(generator) {
		delete(list, generator)
	}
}

func GetGeneratorN(generator string) IGenerator {
	return list[generator]
}