package blocks

import "gomine/worlds"

type Block struct {
	*worlds.Position
	id int
	data byte
	name string

	hasCollisionBox bool
}

/**
 * Returns a new Block. Position is uninitialized.
 */
func NewBlock(id int, data byte, name string) *Block {
	return &Block{worlds.NewPosition(0, 0, 0, worlds.Level{}), id, data, name, true}
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
	return block.hasCollisionBox
}