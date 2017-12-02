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
	collisionBox *vectors.CubesBox
	boundingBox *vectors.CubesBox
}

/**
 * Returns a new Block. Position is uninitialized.
 */
func NewBlock(id int, data byte, name string) *Block {
	return &Block{worlds.NewPosition(0, 0, 0, worlds.Level{}), id, data, name, true, vectors.NewCubesBox([]*vectors.Cube{vectors.NewCube(0, 0, 0, 1, 1, 1)}), vectors.NewCubesBox([]*vectors.Cube{vectors.NewCube(0, 0, 0, 1, 1, 1)})}
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
	return block.hasCollisionBox || block.collisionBox.IsNil()
}

/**
 * Returns the collision box of this block.
 */
func (block *Block) GetCollisionBox() *vectors.CubesBox {
	return block.collisionBox
}

/**
 * Sets the collision box of this block.
 */
func (block *Block) SetCollisionBox(box *vectors.CubesBox) {
	block.collisionBox = box
}

/**
 * Returns the bounding box of this block.
 */
func (block *Block) GetBoundingBox() *vectors.CubesBox {
	return block.boundingBox
}

/**
 * Sets the bounding box of this block.
 */
func (block *Block) SetBoundingBox(box *vectors.CubesBox) {
	block.boundingBox = box
}