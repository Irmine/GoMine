package entities

type Attribute struct {
	name         string
	minValue     float32
	maxValue     float32
	value        float32
	defaultValue float32
}

func (attribute *Attribute) GetName() string {
	return attribute.name
}

func (attribute *Attribute) GetMinValue() float32 {
	return attribute.minValue
}

func (attribute *Attribute) GetMaxValue() float32 {
	return attribute.maxValue
}

func (attribute *Attribute) GetValue() float32 {
	return attribute.value
}

func (attribute *Attribute) GetDefaultValue() float32 {
	return attribute.defaultValue
}

func (attribute *Attribute) SetValue(value float32) {
	attribute.value = value
}

func (attribute *Attribute) SetDefaultValue(value float32) {
	attribute.defaultValue = value
}
