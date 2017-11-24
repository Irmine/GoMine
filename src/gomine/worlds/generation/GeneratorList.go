package generation

var list map[string]IGenerator

func RegisterGenerator(generator IGenerator) {
	list[generator.GetName()] = generator
}

func GeneratorExists(generator IGenerator) bool {
	var _, ok = list[generator.GetName()]
	return ok
}

func GeneratorNameExists(generator string) bool {
	var _, ok = list[generator]
	return ok
}

func DeRegisterGenerator(generator IGenerator)  {
	if GeneratorExists(generator) {
		delete(list, generator.GetName())
	}
}

func DeRegisterGeneratorByName(generator string)  {
	if GeneratorNameExists(generator) {
		delete(list, generator)
	}
}

func GetGeneratorByName(generator string) IGenerator {
	return list[generator]
}