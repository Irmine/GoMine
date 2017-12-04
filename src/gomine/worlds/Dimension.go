package worlds

import (
	"gomine/interfaces"
	"gomine/net"
	"gomine/net/packets"
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
	chunks 		map[int]interfaces.IChunk
	chunkPlayers map[int][]interfaces.IPlayer
	updatedBlocks map[int][]interfaces.IBlock
}

/**
 * Returns a new dimension with the given dimension ID.
 */
func NewDimension(name string, dimensionId int, level *Level, chunks map[int]interfaces.IChunk) *Dimension {
	return &Dimension{name, dimensionId, level, chunks, make(map[int][]interfaces.IPlayer), make(map[int][]interfaces.IBlock)}
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
func (dimension *Dimension) SetChunk(x, z int, chunk interfaces.IChunk) {
	dimension.chunks[GetChunkIndex(x, z)] = chunk
}

/**
 * Gets the chunk in the dimension at the x/z coordinates.
 */
func (dimension *Dimension) GetChunk(x, z int) interfaces.IChunk {
	if v, ok := dimension.chunks[GetChunkIndex(x, z)]; ok {
		return v
	}
	return nil
}

/**
 * Gets all the players located in a chunk.
 */
func (dimension *Dimension) GetChunkPlayers(x, z int) []interfaces.IPlayer {
	if v, ok := dimension.chunkPlayers[GetChunkIndex(x, z)]; ok {
		return v
	}
	return nil
}

/**
 * Adds a player to a chunk.
 */
func (dimension *Dimension) AddChunkPlayer(x, z int, player interfaces.IPlayer) {
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
	for x := -distance; x < distance; x++ {
		for z := -distance; z < distance; z++ {
			player.SendChunk(dimension.GetChunk(int(x), int(z)))
		}
	}
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
