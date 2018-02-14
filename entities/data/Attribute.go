package data

const (
	AttributeHealth = "minecraft:health"
	AttributeMovementSpeed = "minecraft:movement"
	AttributeAttackDamage = "minecraft:attack_damage"

	AttributeAbsorption = "minecraft:absorption"

	AttributeHunger = "minecraft:hunger"
	AttributeSaturation = "minecraft:saturation"
	AttributeExhaustion = "minecraft:exhaustion"

	AttributeKnockBackResistance = "minecraft:knockback_resistance"

	AttributeFollowRange = "minecraft:follow_range"

	AttributeExperienceLevel = "minecraft:player.level"
	AttributeExperience = "minecraft:player.experience"

	AttributeJumpStrength = "minecraft:horse.jump_strength"
)

type Attribute struct {
	name         string
	minValue     float32
	maxValue     float32
	value        float32
	defaultValue float32
}

/**
 * Returns a new Attribute with the given name.
 */
func NewAttribute(name string, value, maxValue float32) *Attribute {
	return &Attribute{name, 0, maxValue, value, value}
}

/**
 * Returns the name of the attribute.
 */
func (attribute *Attribute) GetName() string {
	return attribute.name
}

/**
 * Returns the minimum value of this attribute.
 */
func (attribute *Attribute) GetMinValue() float32 {
	return attribute.minValue
}

/**
 * Returns the maximum value of this attribute.
 */
func (attribute *Attribute) GetMaxValue() float32 {
	return attribute.maxValue
}

/**
 * Returns the current value of this attribute.
 */
func (attribute *Attribute) GetValue() float32 {
	return attribute.value
}

/**
 * Sets the current value of this attribute.
 */
func (attribute *Attribute) SetValue(value float32) {
	attribute.value = value
}

/**
 * Returns the default value of this attribute.
 */
func (attribute *Attribute) GetDefaultValue() float32 {
	return attribute.defaultValue
}

/**
 * Sets the default value of this attribute.
 */
func (attribute *Attribute) SetDefaultValue(value float32) {
	attribute.defaultValue = value
}