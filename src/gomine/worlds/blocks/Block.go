package blocks

import (
	"gomine/worlds"
	"gomine/vectors"
)

type Block struct {
	*worlds.Position
	id int
	data byte
	name string

	hasCollisionBox bool
	CollisionBox *vectors.CubesBox
	BoundingBox *vectors.CubesBox

	blastResistance int
	lightLevel byte
}

/**
 * Returns a new Block.
 */
func NewBlock(id int, data byte, name string) *Block {
	return &Block{Position: worlds.NewPosition(0, 0, 0, worlds.Level{}), id: id, data: data, name: name, hasCollisionBox: true, CollisionBox: vectors.NewCubesBox([]*vectors.Cube{vectors.NewCube(0, 0, 0, 1, 1, 1)}), BoundingBox: vectors.NewCubesBox([]*vectors.Cube{vectors.NewCube(0, 0, 0, 1, 1, 1)})}
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
 * Sets the block's (meta)data.
 */
func (block *Block) SetData(data byte) {
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
 * Returns the blast resistance of this block.
 */
func (block *Block) GetBlastResistance() int {
	return block.blastResistance
}

/**
 * Sets the blast resistance of this block.
 */
func (block *Block) SetBlastResistance(value int) {
	block.blastResistance = value
}

/**
 * Returns the light level of this block.
 */
func (block *Block) GetLightLevel() byte {
	return block.lightLevel
}

/**
 * Sets the light level of this block.
 */
func (block *Block) SetLightLevel(level byte) {
	block.lightLevel = level
}