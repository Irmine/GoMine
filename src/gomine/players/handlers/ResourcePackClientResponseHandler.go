package handlers

import (
	"gomine/net/info"
	"gomine/interfaces"
	"goraklib/server"
	"gomine/net/packets"
	"gomine/vectors"
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
		if response.Status == 3 {
			var pk = packets.NewResourcePackStackPacket()
			server.GetRakLibAdapter().SendPacket(pk, session)
			return true
		}

		var pk4 = packets.NewStartGamePacket()
		pk4.PlayerGameMode = 1
		pk4.PlayerPosition = vectors.TripleVector{0, 20, 0}
		pk4.LevelSeed = 12345
		pk4.Generator = 12345
		pk4.LevelGameMode = 1
		pk4.LevelSpawnPosition = vectors.TripleVector{0, 20, 0}
		pk4.MultiPlayerGame = true
		pk4.BroadcastToXbox = true
		pk4.BroadcastToLan = true
		pk4.CommandsEnabled = true
		pk4.GameRules = server.GetDefaultLevel().GetGameRules()
		pk4.BonusChest = true
		pk4.StartMap = true
		pk4.TrustPlayers = true
		pk4.DefaultPermissionLevel = 2
		pk4.XboxBroadcastMode = 1
		pk4.LevelName = server.GetDefaultLevel().GetName()
		pk4.CurrentTick = int64(server.GetCurrentTick())
		pk4.EnchantmentSeed = 312904

		server.GetRakLibAdapter().SendPacket(pk4, session)

		var pk = packets.NewPlayerListPacket()
		pk.Players = append(pk.Players, player)
		pk.ListType = packets.ListTypeAdd
		server.GetRakLibAdapter().SendPacket(pk, session)

		var pk3 = packets.NewCraftingDataPacket()
		server.GetRakLibAdapter().SendPacket(pk3, session)
	}

	return true
}