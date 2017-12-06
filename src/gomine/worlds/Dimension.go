package worlds

import (
	"gomine/interfaces"
	"gomine/net"
	"gomine/net/packets"
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
	chunkPlayers map[int][]interfaces.IPlayer
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
		chunkPlayers: make(map[int][]interfaces.IPlayer),
		updatedBlocks: make(map[int][]interfaces.IBlock),
	}

	if len(generator) == 0 {
		dimension.generator = generation.GetGeneratorByName(level.server.GetConfiguration().DefaultGenerator)
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
 * Gets all the players located in a chunk.
 */
func (dimension *Dimension) GetChunkPlayers(x, z int32) []interfaces.IPlayer {
	if v, ok := dimension.chunkPlayers[GetChunkIndex(x, z)]; ok {
		return v
	}
	return nil
}

/**
 * Adds a player to a chunk.
 */
func (dimension *Dimension) AddChunkPlayer(x, z int32, player interfaces.IPlayer) {
	dimension.chunkPlayers[GetChunkIndex(x, z)] = append(dimension.chunkPlayers[GetChunkIndex(x, z)], player)
}

/**
 * this function updates every block that gets changed.
 */
func (dimension *Dimension) UpdateBlocks()  {
	var players []interfaces.IPlayer
	batch := net.NewMinecraftPacketBatch()

	for i, blocks := range dimension.updatedBlocks {
		x, z := GetChunkCoordinates(i)
		players = dimension.GetChunkPlayers(x, z)

		if len(players) == 0 {
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

	for _, p := range players {
		dimension.level.GetServer().GetRakLibAdapter().SendBatch(batch, p.GetSession())
	}
}

func (dimension *Dimension) RequestChunks(player interfaces.IPlayer)  {
	distance := player.GetViewDistance()
	for x := -distance; x <= distance; x++ {
		for z := -distance; z <= distance; z++ {
			player.SendChunk(dimension.GetChunk(x, z))
		}
	}
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

/*func (dimension *Dimension) SendChunks() {
	for _, p := range dimension.chunkPlayers {
		for _, p2 := range p {

			p2.SendChunk(dimension.GetChunk(int(p2.GetPosition().X) >> 4, int(p2.GetPosition().X) >> 4))
		}
	}
}*/

func (dimension *Dimension) TickDimension() {
	dimension.UpdateBlocks()
}