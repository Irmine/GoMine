package blocks

import (
	"github.com/irmine/gomine/vectors"
)

type Block struct {
	id   int
	data byte
	name string

	hasCollisionBox bool
	CollisionBox    *vectors.CubesBox
	BoundingBox     *vectors.CubesBox

	hardness        float32
	blastResistance int

	lightEmissionLevel byte
	diffusesLight      bool
	lightFilterLevel   byte
}

/**
 * Returns a new Block.
 */
func NewBlock(id int, data byte, name string) *Block {
	var block = &Block{id: id, data: data, name: name, hasCollisionBox: true, CollisionBox: vectors.NewCubesBox([]*vectors.Cube{vectors.NewCube(0, 0, 0, 1, 1, 1)}), BoundingBox: vectors.NewCubesBox([]*vectors.Cube{vectors.NewCube(0, 0, 0, 1, 1, 1)})}

	block.diffusesLight = true
	block.lightFilterLevel = 15

	return block
}

/**
 * Returns the block ID of this block.
 */
func (block *Block) GetId() int {
	return block.id
}

/**
 * Returns the (meta)data of this block.
 */
func (block *Block) GetData() byte {
	return block.data
}

/**
 * Sets the block's (meta)data/variant.
 */
func (block *Block) SetData(data byte) {
	block.data = data
}

/**
 * Sets the block's (meta)data/variant.
 */
func (block *Block) SetVariant(data byte) {
	block.data = data
}

/**
 * Returns the name of this block.
 */
func (block *Block) GetName() string {
	return block.name
}

/**
 * Returns whether the block has a collision box or not.
 * Blocks such as flowers do not have collision boxes.
 */
func (block *Block) HasCollisionBox() bool {
	return block.hasCollisionBox || block.CollisionBox.IsNil()
}

/**
 * Returns the collision box of this block.
 */
func (block *Block) GetCollisionBox() *vectors.CubesBox {
	return block.CollisionBox
}

/**
 * Sets the collision box of this block.
 */
func (block *Block) SetCollisionBox(box *vectors.CubesBox) {
	block.CollisionBox = box
}

/**
 * Returns the bounding box of this block.
 * All blocks (except air) have bounding boxes.
 */
func (block *Block) GetBoundingBox() *vectors.CubesBox {
	return block.BoundingBox
}

/**
 * Sets the bounding box of this block.
 */
func (block *Block) SetBoundingBox(box *vectors.CubesBox) {
	block.BoundingBox = box
}

/**
 * Sets the blast resistance of this block.
 */
func (block *Block) SetBlastResistance(value int) {
	block.blastResistance = value
}

/**
 * Returns the blast resistance of this block.
 */
func (block *Block) GetBlastResistance() int {
	return block.blastResistance
}

/**
 * Returns the hardness of this block.
 */
func (block *Block) GetHardness() float32 {
	return block.hardness
}

/**
 * Sets the hardness of this block.
 */
func (block *Block) SetHardness(value float32) {
	block.hardness = value
}

/**
 * Returns the light level that this block emits.
 */
func (block *Block) GetLightEmissionLevel() byte {
	return block.lightEmissionLevel
}

/**
 * Sets the light level this block emits.
 */
func (block *Block) SetLightEmissionLevel(level byte) {
	block.lightEmissionLevel = level
}

/**
 * Returns if this block diffuses (breaks) sky light.
 */
func (block *Block) DiffusesLight() bool {
	return block.diffusesLight
}

/**
 * Sets this block to diffuse sky light.
 */
func (block *Block) SetLightDiffusing(value bool) {
	block.diffusesLight = value
}

/**
 * Returns the amount of light levels this block will filter when light goes through.
 */
func (block *Block) GetLightFilterLevel() byte {
	if !block.diffusesLight {
		return 0
	}
	return block.lightFilterLevel
}

/**
 * Sets the light filter level of this block.
 */
func (block *Block) SetLightFilterLevel(value byte) {
	block.lightFilterLevel = value
}
