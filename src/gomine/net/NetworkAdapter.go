package net

import (
	"gomine/interfaces"
	server2 "goraklib/server"
	"gomine/net/info"
	"goraklib/protocol"
	"gomine/players/handlers"
)

type NetworkAdapter struct {
	server interfaces.IServer
	rakLibServer *server2.GoRakLibServer
}

/**
 * Returns a new Network adapter to adapt to the RakNet server.
 */
func NewNetworkAdapter(server interfaces.IServer) *NetworkAdapter {
	var rakServer = server2.NewGoRakLibServer(server.GetName(), server.GetAddress(), server.GetPort())
	rakServer.SetMinecraftProtocol(info.LatestProtocol)
	rakServer.SetMinecraftVersion(info.GameVersionNetwork)
	rakServer.SetMaxConnectedSessions(server.GetMaximumPlayers())
	rakServer.SetDefaultGameMode("Creative")
	rakServer.SetMotd(server.GetMotd())

	return &NetworkAdapter{server, rakServer}
}

/**
 * Returns the GoRakLib server.
 */
func (adapter *NetworkAdapter) GetRakLibServer() *server2.GoRakLibServer {
	return adapter.rakLibServer
}

/**
 * Ticks the adapter
 */
func (adapter *NetworkAdapter) Tick() {
	go adapter.rakLibServer.Tick()

	for _, session := range adapter.rakLibServer.GetSessionManager().GetSessions() {
		go func(session *server2.Session) {
			for _, encapsulatedPacket := range session.GetReadyEncapsulatedPackets() {

				player, _ := adapter.server.GetPlayerFactory().GetPlayerBySession(session)

				batch := NewMinecraftPacketBatch(player, adapter.server.GetLogger())
				batch.Buffer = encapsulatedPacket.Buffer
				batch.Decode()

				for _, packet := range batch.GetPackets() {
					packet.DecodeHeader()
					packet.Decode()

					priorityHandlers := GetPacketHandlers(packet.GetId())

					var handled = false
					for _, h := range priorityHandlers {
						for _, handler := range h {
							if packet.IsDiscarded() {
								return
							}

							ret := handler.Handle(packet, player, session, adapter.server)
							if !handled {
								handled = ret
							}
						}
					}

					if !handled {
						adapter.server.GetLogger().Debug("Unhandled Minecraft packet with ID:", packet.GetId())
					}
				}
			}
		}(session)
	}

	for _, pk := range adapter.rakLibServer.GetRawPackets() {
		adapter.server.HandleRaw(pk)
	}

	for _, session := range adapter.rakLibServer.GetSessionManager().GetDisconnectedSessions() {
		player, _ := adapter.server.GetPlayerFactory().GetPlayerBySession(session)
		handler := handlers.NewDisconnectHandler()
		handler.Handle(player, session, adapter.server)
	}
}

func (adapter *NetworkAdapter) GetSession(address string, port uint16) *server2.Session {
	var session, _ = adapter.rakLibServer.GetSessionManager().GetSession(address, port)
	return session
}

func (adapter *NetworkAdapter) SendPacket(pk interfaces.IPacket, session interfaces.IMinecraftSession, priority byte) {
	var b = NewMinecraftPacketBatch(session, adapter.server.GetLogger())
	b.AddPacket(pk)

	adapter.SendBatch(b, session.GetSession(), priority)
}

func (adapter *NetworkAdapter) SendBatch(batch interfaces.IMinecraftPacketBatch, session *server2.Session, priority byte) {
	session.SendConnectedPacket(batch, protocol.ReliabilityReliableOrdered, priority)
}

/**
 * Returns if a packet with the given ID is registered.
 */
func (adapter *NetworkAdapter) IsPacketRegistered(id int) bool {
	return IsPacketRegistered(id)
}

/**
 * Returns a new packet with the given ID and a function that returns that packet.
 */
func (adapter *NetworkAdapter) RegisterPacket(id int, function func() interfaces.IPacket) {
	RegisterPacket(id, function)
}

/**
 * Returns a new packet with the given ID.
 */
func (adapter *NetworkAdapter) GetPacket(id int) interfaces.IPacket {
	return GetPacket(id)
}

/**
 * Registers a new packet handler to listen for packets with the given ID.
 * Returns a bool indicating success.
 */
func (adapter *NetworkAdapter) RegisterPacketHandler(id int, handler interfaces.IPacketHandler, priority int) bool {
	return RegisterPacketHandler(id, handler, priority)
}

/**
 * Returns all packet handlers registered on the given ID.
 */
func (adapter *NetworkAdapter) GetPacketHandlers(id int) [][]interfaces.IPacketHandler {
	return GetPacketHandlers(id)
}

/**
 * Deletes all packet handlers listening for packets with the given ID, on the given priority.
 */
func (adapter *NetworkAdapter) DeregisterPacketHandlers(id int, priority int) {
	DeregisterPacketHandlers(id, priority)
}

/**
 * Deletes a registered packet with the given ID.
 */
func (adapter *NetworkAdapter) DeletePacket(id int) {
	DeregisterPacket(id)
}