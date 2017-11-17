package entities

type Attribute struct {
	name         string
	value        float32
	defaultValue float32
}

func (attribute *Attribute) getName() string {
	return attribute.name
}

func (attribute *Attribute) getValue() float32 {
	return attribute.value
}

func (attribute *Attribute) getDefaultValue() float32 {
	return attribute.defaultValue
}

func (attribute *Attribute) setValue(value float32) {
	attribute.value = value
}

func (attribute *Attribute) setDefaultValue(value float32) {
	attribute.defaultValue = value
}
