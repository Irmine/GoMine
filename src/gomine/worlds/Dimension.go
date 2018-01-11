package worlds

import (
	"gomine/interfaces"
	"gomine/worlds/generation"
	"gomine/worlds/chunks"
)

const (
	OverworldId = 0
	NetherId    = 1
	EndId	    = 2
)

type Dimension struct {
	name 		string
	dimensionId int
	level       interfaces.ILevel
	isGenerated bool

	chunks 		map[int]interfaces.IChunk
	updatedBlocks map[int][]interfaces.IBlock

	generator interfaces.IGenerator
}

/**
 * Returns a new dimension with the given dimension ID.
 */
func NewDimension(name string, dimensionId int, level *Level, generator string, chunks map[int]interfaces.IChunk) *Dimension {
	var dimension = &Dimension{
		name:  name,
		dimensionId: dimensionId,
		level: level,
		chunks: chunks,
		updatedBlocks: make(map[int][]interfaces.IBlock),
	}

	if len(generator) == 0 {
		dimension.generator = generation.GetGeneratorByName(level.server.GetConfiguration().DefaultGenerator)
	} else {
		dimension.generator = generation.GetGeneratorByName(generator)
	}

	return dimension
}

/**
 * Returns the dimension ID of this dimension.
 */
func (dimension *Dimension) GetDimensionId() int {
	return dimension.dimensionId
}

/**
 * Returns the name of this dimension.
 */
func (dimension *Dimension) GetName() string {
	return dimension.name
}

/**
 * Returns the level this dimension is in.
 */
func (dimension *Dimension) GetLevel() interfaces.ILevel {
	return dimension.level
}

/**
 * Returns if chunk is loaded
 */
func (dimension *Dimension) IsChunkLoaded(x, z int32) bool {
	var _, ok = dimension.chunks[GetChunkIndex(x, z)]
	return ok
}

/**
 * Sets this chunk unloaded
 */
func (dimension *Dimension) SetChunkUnloaded(x, z int32) {
	if !dimension.IsChunkLoaded(x, z) {
		delete(dimension.chunks, GetChunkIndex(x, z))
	}
}

/**
 * Sets a new chunk in the dimension at the x/z coordinates.
 */
func (dimension *Dimension) SetChunk(x, z int32, chunk interfaces.IChunk) {
	dimension.chunks[GetChunkIndex(x, z)] = chunk
}

/**
 * Gets the chunk in the dimension at the x/z coordinates.
 */
func (dimension *Dimension) GetChunk(x, z int32) interfaces.IChunk {
	if v, ok := dimension.chunks[GetChunkIndex(x, z)]; ok {
		return v
	} else {
		var chunk = dimension.generator.GetNewChunk(chunks.NewChunk(x, z))
		dimension.chunks[GetChunkIndex(x, z)] = chunk
		return chunk
	}
	return nil
}

/**
 * Returns if the dimension is generated or not.
 */
func (dimension *Dimension) IsGenerated() bool {
	return dimension.isGenerated
}

/**
 * Sets the generator of this dimension.
 */
func (dimension *Dimension) SetGenerator(generator interfaces.IGenerator) {
	dimension.generator = generator
}

/**
 * Returns the generator of this level.
 */
func (dimension *Dimension) GetGenerator() interfaces.IGenerator {
	return dimension.generator
}

/**
 * Sends all chunks required around the player.
 */
func (dimension *Dimension) RequestChunks(player interfaces.IPlayer, distance int32) {
	xD, zD := int32(player.GetPosition().X) >> 4, int32(player.GetPosition().Z) >> 4

	for x := -distance + xD; x <= distance + xD; x++ {
		for z := -distance + zD; z <= distance + zD; z++ {

			var xRel = x - xD
			var zRel = z - zD
			if xRel * xRel + zRel * zRel <= distance * distance {
				index := GetChunkIndex(x, z)

				if player.HasChunkInUse(index) {
					continue
				}

				chunk := dimension.GetChunk(x, z)
				player.SendChunk(chunk, index)

				for _, entity := range chunk.GetEntities() {
					entity.SpawnTo(player)
				}
			}
		}
	}
}

/**
 * Unloads all unused chunks
 */
func (dimension Dimension) UnloadUnusedChunks() {

}

/**
 * this function updates every block that gets changed.
 */
func (dimension *Dimension) UpdateBlocks()  {
	/*var players2 []interfaces.IPlayer
	batch := net.NewMinecraftPacketBatch()

	for i, blocks := range dimension.updatedBlocks {
		x, z := GetChunkCoordinates(i)
		players2 = dimension.GetChunkPlayers(x, z)

		if len(players2) == 0 {
			delete(dimension.chunkPlayers, GetChunkIndex(x, z))
			break
		}

		for _, block := range blocks {
			pk := packets.NewUpdateBlockPacket()
			pk.BlockId = uint32(block.GetId())
			pk.BlockMetadata = uint32(block.GetData())
			pk.Flags = 0x0
			batch.AddPacket(pk)
		}
	}

	for _, p := range players2 {
		dimension.level.GetServer().GetRakLibAdapter().SendBatch(batch, p.GetSession(), server.PriorityMedium)
	}*/
}

/**
 * Unloads all unused chunks of the dimension.
 */
func (dimension *Dimension) UpdateChunks() {
	dimension.UnloadUnusedChunks()
}

func (dimension *Dimension) TickDimension() {
	dimension.UpdateBlocks()
	//dimension.UpdateChunks()
}