package data

import (
	"math"
)

// AttributeMap is a struct containing an unlimited amount of attributes.
type AttributeMap struct {
	attributes map[string]*Attribute
}

var defaultAttributes = map[string]*Attribute{
	AttributeHealth:              NewAttribute(AttributeHealth, 20, 1024),
	AttributeMovementSpeed:       NewAttribute(AttributeMovementSpeed, 0.1, 1024),
	AttributeAttackDamage:        NewAttribute(AttributeAttackDamage, 2, 2048),
	AttributeAbsorption:          NewAttribute(AttributeAbsorption, 0, 1024),
	AttributeHunger:              NewAttribute(AttributeHunger, 20, 20),
	AttributeSaturation:          NewAttribute(AttributeSaturation, 20, 20),
	AttributeExhaustion:          NewAttribute(AttributeExhaustion, 0, 5),
	AttributeKnockBackResistance: NewAttribute(AttributeKnockBackResistance, 0, 1),
	AttributeFollowRange:         NewAttribute(AttributeFollowRange, 32, 2048),
	AttributeExperience:          NewAttribute(AttributeExperience, 0, 1),
	AttributeExperienceLevel:     NewAttribute(AttributeExperienceLevel, 0, math.MaxInt32),
	AttributeJumpStrength:        NewAttribute(AttributeJumpStrength, 0.7, 2),
}

// NewAttributeMap returns a new attribute map with default attributes.
func NewAttributeMap() *AttributeMap {
	return &AttributeMap{defaultAttributes}
}

// Exists checks if an attribute with the given name exists.
func (attMap *AttributeMap) Exists(name string) bool {
	var _, ok = attMap.attributes[name]
	return ok
}

// GetAttributes returns all attributes in a name => attribute map.
func (attMap *AttributeMap) GetAttributes() map[string]*Attribute {
	return attMap.attributes
}

// SetAttribute sets an attribute in this attribute map.
func (attMap *AttributeMap) SetAttribute(attribute *Attribute) {
	attMap.attributes[attribute.GetName()] = attribute
}

// GetAttribute returns an attribute with the given name.
func (attMap *AttributeMap) GetAttribute(name string) *Attribute {
	return attMap.attributes[name]
}
