package entityPkg

type attribute struct {
	name string
	value float32
	defaultValue float32
}

func (attribute *attribute) getName() string {
	return attribute.name
}

func (attribute *attribute) getValue() float32 {
	return attribute.value
}

func (attribute *attribute) getDefaultValue() float32 {
	return attribute.defaultValue
}

func (attribute *attribute) setValue(value float32) {
	attribute.value = value
}

func (attribute *attribute) setDefaultValue(value float32) {
	attribute.defaultValue = value
}
