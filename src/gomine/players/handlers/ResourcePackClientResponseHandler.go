package handlers

import (
	"gomine/net/info"
	"gomine/interfaces"
	"goraklib/server"
	"gomine/net/packets"
	"gomine/vectors"
	"gomine/entities/math"
	math2 "math"
	"gomine/permissions"
)

const (
	ChunkSize = 1048576
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
			stack.ResourcePacks = server.GetPackHandler().GetResourceStack().GetPacks()
			stack.BehaviorPacks = server.GetPackHandler().GetBehaviorStack().GetPacks()
			stack.MustAccept = server.GetConfiguration().ForceResourcePacks
			player.SendPacket(stack)

		case packets.StatusCompleted:
			player.PlaceInWorld(vectors.NewTripleVector(0, 20, 0), math.NewRotation(0, 0, 0), server.GetDefaultLevel(), server.GetDefaultLevel().GetDefaultDimension())
			player.SetFinalized()

			var startGame = packets.NewStartGamePacket()
			startGame.Generator = 1
			startGame.LevelSeed = 312402
			startGame.TrustPlayers = true
			startGame.DefaultPermissionLevel = permissions.LevelOperator
			startGame.EntityRuntimeId = player.GetRuntimeId()
			startGame.EntityUniqueId = player.GetUniqueId()
			startGame.PlayerGameMode = 1
			startGame.PlayerPosition = vectors.TripleVector{20, 20, 20}
			startGame.LevelGameMode = 1
			startGame.LevelSpawnPosition = vectors.TripleVector{0, 20, 0}
			startGame.CommandsEnabled = true
			startGame.GameRules = server.GetDefaultLevel().GetGameRules()
			startGame.LevelName = server.GetDefaultLevel().GetName()
			startGame.CurrentTick = int64(server.GetCurrentTick())
			startGame.Time = 0
			startGame.AchievementsDisabled = true
			startGame.BroadcastToXbox = true
			startGame.XBOXBroadcastMode = 0

			player.SendPacket(startGame)

			var craftingData = packets.NewCraftingDataPacket()
			player.SendPacket(craftingData)
		}
		return true
	}

	return false
}