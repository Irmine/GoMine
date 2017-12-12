package worlds

import (
	"gomine/interfaces"
	"gomine/net"
	"gomine/net/packets"
	"gomine/worlds/generation"
	"gomine/worlds/chunks"
	"gomine/players"
	"goraklib/server"
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
	loadedChunks map[int]interfaces.IChunk

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
		loadedChunks: make(map[int]interfaces.IChunk),
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
	var _, ok = dimension.loadedChunks[GetChunkIndex(x, z)]
	return ok
}

/**
 * Sets this chunk loaded
 */
func (dimension *Dimension) SetChunkLoaded(x, z int32, chunk interfaces.IChunk) {
	if !dimension.IsChunkLoaded(x, z) {
		dimension.loadedChunks[GetChunkIndex(x, z)] = chunk
	}
}

/**
 * Sets this chunk unloaded
 */
func (dimension *Dimension) SetChunkUnloaded(x, z int32) {
	if !dimension.IsChunkLoaded(x, z) {
		delete(dimension.loadedChunks, GetChunkIndex(x, z))
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
 * Sends chunks around a player
 */
func (dimension *Dimension) RequestChunks(player interfaces.IPlayer) {
	/*distance := player.GetViewDistance()
	xD, zD := int32(player.GetPosition().X) >> 4, int32(player.GetPosition().Z) >> 4

	for x := xD - distance; x <= (xD + distance); x++ {
		for z := zD - distance; z <= (zD + distance); z++ {
			if !dimension.IsChunkLoaded(x, z) {
				chunk := dimension.GetChunk(x, z)
				player.SendChunk(chunk)
				dimension.AddChunkPlayer(x, z, player)
				dimension.SetChunkLoaded(x, z, chunk)

				if dimension.IsChunkLoaded(x, z) {
					dimension.level.GetServer().GetLogger().Debug("Chunk at x: " + strconv.Itoa(int(x)) + ", z: " + strconv.Itoa(int(z)) + "loaded!") // debug, don't remove
				}
			}
		}
	}*/ // This loads chunks incorrectly and currently dimension loaded and player loaded chunks are treated the same, which causes problems for multiple players.

	distance := player.GetViewDistance()
	for x := -distance; x <= distance; x++ {
		for z := -distance; z <= distance; z++ {
			player.SendChunk(dimension.GetChunk(x, z))
		}
	}
}

/**
 * Unloads all unused chunks
 */
func (dimension Dimension) UnloadUnusedChunks() {
	for index := range dimension.loadedChunks {
		x, z := GetChunkCoordinates(index)
		if len(dimension.GetChunkPlayers(x, z)) == 0 {
			dimension.SetChunkUnloaded(x, z)
		}
	}
}

/**
 * this function updates every block that gets changed.
 */
func (dimension *Dimension) UpdateBlocks()  {
	var players2 []interfaces.IPlayer
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
	}
}

/**
 * This functions updates all chunks for every player in it
 */
func (dimension *Dimension) UpdateChunks() {
	for _, p := range dimension.chunkPlayers {
		for _, p2 := range p {
			p2, ok := p2.(*players.Player)
			if ok {
				dimension.RequestChunks(p2)
			}
		}
	}
	dimension.UnloadUnusedChunks()
}

func (dimension *Dimension) TickDimension() {
	dimension.UpdateBlocks()
	dimension.UpdateChunks()
}