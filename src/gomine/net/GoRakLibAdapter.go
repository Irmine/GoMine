package net

import (
	"gomine/interfaces"
	server2 "goraklib/server"
	"gomine/net/info"
	"goraklib/protocol"
)

type GoRakLibAdapter struct {
	server interfaces.IServer
	rakLibServer *server2.GoRakLibServer
}

/**
 * Returns a new GoRakLib adapter to adapt to the RakNet server.
 */
func NewGoRakLibAdapter(server interfaces.IServer) *GoRakLibAdapter {
	var rakServer = server2.NewGoRakLibServer(server.GetName(), server.GetAddress(), server.GetPort())
	rakServer.SetMinecraftProtocol(info.LatestProtocol)
	rakServer.SetMinecraftVersion(info.GameVersionNetwork)
	rakServer.SetServerName(server.GetServerName())
	rakServer.SetMaxConnectedSessions(server.GetMaximumPlayers())
	rakServer.SetDefaultGameMode("Creative")
	rakServer.SetMotd(server.GetMotd())

	return &GoRakLibAdapter{server, rakServer}
}

/**
 * Returns the GoRakLib server.
 */
func (adapter *GoRakLibAdapter) GetRakLibServer() *server2.GoRakLibServer {
	return adapter.rakLibServer
}

/**
 * Ticks the adapter
 */
func (adapter *GoRakLibAdapter) Tick() {
	go adapter.rakLibServer.Tick()

	for _, session := range adapter.rakLibServer.GetSessionManager().GetSessions() {
		go func(session *server2.Session) {
			for _, encapsulatedPacket := range session.GetReadyEncapsulatedPackets() {

				batch := NewMinecraftPacketBatch()
				batch.Buffer = encapsulatedPacket.Buffer
				batch.Decode(adapter.server.GetLogger())

				for _, packet := range batch.GetPackets() {
					packet.DecodeHeader()
					packet.Decode()

					var player, _ = adapter.server.GetPlayerFactory().GetPlayerBySession(session.GetAddress(), session.GetPort())

					priorityHandlers := GetPacketHandlers(packet.GetId())

					for _, handlers := range priorityHandlers {
						for _, handler := range handlers {
							handler.Handle(packet, player, session, adapter.server)
						}
					}

					if len(priorityHandlers) == 0 {
						adapter.server.GetLogger().Debug("Unhandled Minecraft packet with ID:", packet.GetId())
					}
				}
			}
		}(session)
	}
}

func (adapter *GoRakLibAdapter) GetSession(address string, port uint16) *server2.Session {
	var session, _ = adapter.rakLibServer.GetSessionManager().GetSession(address, port)
	return session
}

func (adapter *GoRakLibAdapter) SendPacket(pk interfaces.IPacket, session *server2.Session, priority byte) {
	var b = NewMinecraftPacketBatch()
	b.AddPacket(pk)

	adapter.SendBatch(b, session, priority)
}

func (adapter *GoRakLibAdapter) SendBatch(batch interfaces.IMinecraftPacketBatch, session *server2.Session, priority byte) {
	session.SendConnectedPacket(batch, protocol.ReliabilityReliableOrdered, priority)
}