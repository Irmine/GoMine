package handlers

import (
	"gomine/net/info"
	"gomine/interfaces"
	"gomine/net/packets"
	"goraklib/server"
)

type LoginHandler struct {
	*PacketHandler
}

func NewLoginHandler() LoginHandler {
	return LoginHandler{NewPacketHandler(info.LoginPacket)}
}

/**
 * Handles the main login process.
 */
func (handler LoginHandler) Handle(packet interfaces.IPacket, player interfaces.IPlayer, session *server.Session, server interfaces.IServer) bool {
	if loginPacket, ok := packet.(*packets.LoginPacket); ok {
		_, online := server.GetPlayerFactory().GetPlayerByName(loginPacket.Username)
		if online == nil {
			//return false // Players can't 'quit' currently, so don't check for double names.
		}

		var player = player.New(server, session, loginPacket.Username, loginPacket.ClientUUID, loginPacket.ClientXUID, loginPacket.ClientId)
		player.SetLanguage(loginPacket.Language)
		player.SetSkinId(loginPacket.SkinId)
		player.SetSkinData(loginPacket.SkinData)
		player.SetCapeData(loginPacket.CapeData)
		player.SetGeometryName(loginPacket.GeometryName)
		player.SetGeometryData(string(loginPacket.GeometryData))

		playStatus := packets.NewPlayStatusPacket()
		playStatus.Status = 0
		player.SendPacket(playStatus)

		resourceInfo := packets.NewResourcePackInfoPacket()
		resourceInfo.MustAccept = server.GetConfiguration().ForceResourcePacks

		resourceInfo.ResourcePacks = server.GetPackHandler().GetResourceStack().GetPacks()
		resourceInfo.BehaviorPacks = server.GetPackHandler().GetBehaviorStack().GetPacks()

		player.SendPacket(resourceInfo)

		server.GetPlayerFactory().AddPlayer(player, session)
	}

	return true
}