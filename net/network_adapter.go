package net

import (
	"github.com/irmine/gomine/interfaces"
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/players/handlers/p200"
	"github.com/irmine/goraklib/protocol"
	server2 "github.com/irmine/goraklib/server"
)

type NetworkAdapter struct {
	server       interfaces.IServer
	rakLibServer *server2.GoRakLibServer
	protocolPool *ProtocolPool
}

// NewNetworkAdapter returns a new Network adapter to adapt to the RakNet server.
func NewNetworkAdapter(server interfaces.IServer) *NetworkAdapter {
	var rakServer = server2.NewGoRakLibServer(server.GetName(), server.GetAddress(), server.GetPort())
	rakServer.SetMinecraftProtocol(info.LatestProtocol)
	rakServer.SetMinecraftVersion(info.LatestGameVersionNetwork)
	rakServer.SetMaxConnectedSessions(server.GetMaximumPlayers())
	rakServer.SetDefaultGameMode("Creative")
	rakServer.SetMotd(server.GetMotd())

	var adapter = &NetworkAdapter{server, rakServer, NewProtocolPool()}
	adapter.protocolPool.RegisterDefaults()

	return adapter
}

// GetProtocolPool returns the protocol pool of the network adapter.
func (adapter *NetworkAdapter) GetProtocolPool() interfaces.IProtocolPool {
	return adapter.protocolPool
}

// GetRakLibServer returns the GoRakLib server.
func (adapter *NetworkAdapter) GetRakLibServer() *server2.GoRakLibServer {
	return adapter.rakLibServer
}

// Tick ticks the adapter, ticking the GoRakLib server and processing packets.
func (adapter *NetworkAdapter) Tick() {
	go adapter.rakLibServer.Tick()

	for _, session := range adapter.rakLibServer.GetSessionManager().GetSessions() {
		go func(session *server2.Session) {
			var player, _ = adapter.server.GetPlayerFactory().GetPlayerBySession(session)
			if !adapter.server.GetPlayerFactory().PlayerExistsBySession(session) {
				player = player.New(adapter.server, &MinecraftSession{server: adapter.server, session: session}, "")
			}

			adapter.HandlePackets(session, player)
		}(session)
	}

	for _, pk := range adapter.rakLibServer.GetRawPackets() {
		adapter.server.HandleRaw(pk)
	}

	for _, session := range adapter.rakLibServer.GetSessionManager().GetDisconnectedSessions() {
		player, _ := adapter.server.GetPlayerFactory().GetPlayerBySession(session)
		handler := p200.NewDisconnectHandler()
		handler.Handle(player, session, adapter.server)
	}
}

// HandlePackets handles all packets of the given session + player.
func (adapter *NetworkAdapter) HandlePackets(session *server2.Session, player interfaces.IPlayer) {
	for _, encapsulatedPacket := range session.GetReadyEncapsulatedPackets() {
		batch := NewMinecraftPacketBatch(player, adapter.server.GetLogger())
		batch.Buffer = encapsulatedPacket.Buffer
		batch.Decode()

		for _, packet := range batch.GetPackets() {
			if player.GetProtocolNumber() < 120 {
				packet.DecodeId()
			} else {
				packet.DecodeHeader()
			}
			packet.Decode()

			player.HandlePacket(packet, player)
		}
	}
}

// GetSession returns a GoRakLib session by an address and port.
func (adapter *NetworkAdapter) GetSession(address string, port uint16) *server2.Session {
	var session, _ = adapter.rakLibServer.GetSessionManager().GetSession(address, port)
	return session
}

// SendPacket sends a packet to the given Minecraft session with the given priority.
func (adapter *NetworkAdapter) SendPacket(pk interfaces.IPacket, session interfaces.IMinecraftSession, priority byte) {
	var b = NewMinecraftPacketBatch(session, adapter.server.GetLogger())
	b.AddPacket(pk)

	adapter.SendBatch(b, session.GetSession(), priority)
}

// SendBatch sends a Minecraft packet batch to the given GoRakLib session with the given priority.
func (adapter *NetworkAdapter) SendBatch(batch interfaces.IMinecraftPacketBatch, session *server2.Session, priority byte) {
	session.SendConnectedPacket(batch, protocol.ReliabilityReliableOrdered, priority)
}
