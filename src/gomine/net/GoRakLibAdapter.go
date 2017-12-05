package net

import (
	"gomine/interfaces"
	server2 "goraklib/server"
	"gomine/net/info"
	"goraklib/protocol"
	"strconv"
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
	rakServer.SetConnectedSessionCount(0)
	rakServer.SetDefaultGameMode("Creative")
	rakServer.SetMotd(server.GetMotd())

	InitPacketPool()
	InitHandlerPool()

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

	go func() {
		for _, session := range adapter.rakLibServer.GetSessionManager().GetSessions() {
			for _, encapsulatedPacket := range session.GetReadyEncapsulatedPackets() {

				batch := NewMinecraftPacketBatch()
				batch.stream.Buffer = encapsulatedPacket.Buffer
				batch.Decode(adapter.server.GetLogger())

				for _, packet := range batch.GetPackets() {
					packet.DecodeHeader()
					packet.Decode()

					var player, _ = adapter.server.GetPlayerFactory().GetPlayerBySession(session.GetAddress(), session.GetPort())

					handlers := GetPacketHandlers(packet.GetId())

					for _, handler := range handlers {
						handler.Handle(packet, player, session, adapter.server)
					}
					if len(handlers) == 0 {
						adapter.server.GetLogger().Debug("Unhandled Minecraft packet with ID: " + strconv.Itoa(packet.GetId()))
					}
				}
			}
		}
	}()
}

func (adapter *GoRakLibAdapter) GetSession(address string, port uint16) *server2.Session {
	var session, _ = adapter.rakLibServer.GetSessionManager().GetSession(address, port)
	return session
}

func (adapter *GoRakLibAdapter) SendPacket(pk interfaces.IPacket, session *server2.Session) {
	pk.EncodeHeader()
	pk.Encode()

	var b = NewMinecraftPacketBatch()
	b.AddPacket(pk)

	adapter.SendBatch(b, session)
}

func (adapter *GoRakLibAdapter) SendBatch(batch interfaces.IMinecraftPacketBatch, session *server2.Session) {
	batch.Encode()

	var encPacket = protocol.NewEncapsulatedPacket()
	encPacket.SetBuffer(batch.GetStream().GetBuffer())

	var datagram = protocol.NewDatagram()
	datagram.AddPacket(&encPacket)

	adapter.rakLibServer.SendPacket(datagram, session)
}