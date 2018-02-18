package data

const (
	AttributeHealth        = "minecraft:health"
	AttributeMovementSpeed = "minecraft:movement"
	AttributeAttackDamage  = "minecraft:attack_damage"

	AttributeAbsorption = "minecraft:absorption"

	AttributeHunger     = "minecraft:hunger"
	AttributeSaturation = "minecraft:saturation"
	AttributeExhaustion = "minecraft:exhaustion"

	AttributeKnockBackResistance = "minecraft:knockback_resistance"

	AttributeFollowRange = "minecraft:follow_range"

	AttributeExperienceLevel = "minecraft:player.level"
	AttributeExperience      = "minecraft:player.experience"

	AttributeJumpStrength = "minecraft:horse.jump_strength"
)

// Attribute is a struct containing data of an entity property.
type Attribute struct {
	name         string
	minValue     float32
	maxValue     float32
	value        float32
	defaultValue float32
}

// NewAttribute returns a new Attribute with the given name.
func NewAttribute(name string, value, maxValue float32) *Attribute {
	return &Attribute{name, 0, maxValue, value, value}
}

// GetName returns the name of the attribute.
func (attribute *Attribute) GetName() string {
	return attribute.name
}

// GetMinValue returns the minimum value of this attribute.
func (attribute *Attribute) GetMinValue() float32 {
	return attribute.minValue
}

// GetMaxValue returns the maximum value of this attribute.
func (attribute *Attribute) GetMaxValue() float32 {
	return attribute.maxValue
}

// GetValue returns the current value of this attribute.
func (attribute *Attribute) GetValue() float32 {
	return attribute.value
}

// SetValue sets the current value of this attribute.
func (attribute *Attribute) SetValue(value float32) {
	attribute.value = value
}

// GetDefaultValue returns the default value of this attribute.
func (attribute *Attribute) GetDefaultValue() float32 {
	return attribute.defaultValue
}

// SetDefaultValue sets the default value of this attribute.
func (attribute *Attribute) SetDefaultValue(value float32) {
	attribute.defaultValue = value
}
