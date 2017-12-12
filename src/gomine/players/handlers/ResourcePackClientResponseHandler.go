package handlers

import (
	"gomine/net/info"
	"gomine/interfaces"
	"goraklib/server"
	"gomine/net/packets"
	"gomine/vectors"
	"gomine/entities/math"
	math2 "math"
)

const (
	ChunkSize = 1000000
)

type ResourcePackClientResponseHandler struct {
	*PacketHandler
}

func NewResourcePackClientResponseHandler() ResourcePackClientResponseHandler {
	return ResourcePackClientResponseHandler{NewPacketHandler(info.ResourcePackClientResponsePacket)}
}

/**
 * Handles the resource pack client response.
 */
func (handler ResourcePackClientResponseHandler) Handle(packet interfaces.IPacket, player interfaces.IPlayer, session *server.Session, server interfaces.IServer) bool {
	if response, ok := packet.(*packets.ResourcePackClientResponsePacket); ok {
		switch response.Status {
		case packets.StatusRefused:
			// TODO: Kick the player. We can't kick yet.
			return false

		case packets.StatusSendPacks:
			for _, packUUID := range response.PackUUIDs {
				if !server.GetPackHandler().IsPackLoaded(packUUID) {
					// TODO: Kick the player. We can't kick yet.
					return false
				}
				pack := server.GetPackHandler().GetPack(packUUID)

				dataInfo := packets.NewResourcePackDataInfoPacket()
				dataInfo.PackUUID = packUUID
				dataInfo.MaxChunkSize = ChunkSize
				dataInfo.ChunkCount = int32(math2.Ceil(float64(pack.GetFileSize()) / float64(ChunkSize)))
				dataInfo.CompressedPackSize = pack.GetFileSize()
				dataInfo.Sha256 = pack.GetSha256()

				player.SendPacket(dataInfo)
			}

		case packets.StatusHaveAllPacks:
			var stack = packets.NewResourcePackStackPacket()
			stack.ResourcePacks = server.GetPackHandler().GetResourcePackSlice()
			stack.BehaviorPacks = server.GetPackHandler().GetBehaviorPackSlice()
			player.SendPacket(stack)

		case packets.StatusCompleted:
			player.PlaceInWorld(vectors.TripleVector{0, 20, 0}, math.Rotation{0, 0, 0}, server.GetDefaultLevel(), server.GetDefaultLevel().GetDefaultDimension())

			var startGame = packets.NewStartGamePacket()
			startGame.PlayerGameMode = 1
			startGame.PlayerPosition = vectors.TripleVector{0, 20, 0}
			startGame.LevelSeed = 12345
			startGame.Generator = 12345
			startGame.LevelGameMode = 1
			startGame.LevelSpawnPosition = vectors.TripleVector{0, 20, 0}
			startGame.MultiPlayerGame = true
			startGame.BroadcastToXbox = true
			startGame.BroadcastToLan = true
			startGame.CommandsEnabled = true
			startGame.GameRules = server.GetDefaultLevel().GetGameRules()
			startGame.BonusChest = false
			startGame.StartMap = false
			startGame.TrustPlayers = true
			startGame.DefaultPermissionLevel = 0
			startGame.XboxBroadcastMode = 0
			startGame.LevelName = server.GetDefaultLevel().GetName()
			startGame.CurrentTick = int64(server.GetCurrentTick())
			startGame.EnchantmentSeed = 312904

			player.SendPacket(startGame)

			var playerList = packets.NewPlayerListPacket()
			playerList.Players = append(playerList.Players, player)
			playerList.ListType = packets.ListTypeAdd
			player.SendPacket(playerList)

			var craftingData = packets.NewCraftingDataPacket()
			player.SendPacket(craftingData)
		}
	}

	return true
}