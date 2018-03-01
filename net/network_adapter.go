package net

import (
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
	protocol2 "github.com/irmine/gomine/net/protocol"
	"github.com/irmine/gomine/resources"
	"github.com/irmine/gomine/utils"
	"github.com/irmine/goraklib/protocol"
	"github.com/irmine/goraklib/server"
)

type NetworkAdapter struct {
	logger          *utils.Logger
	rakLibServer    *server.GoRakLibServer
	protocolManager protocol2.Manager
	sessionManager  *SessionManager

	DisconnectFunction func(session *MinecraftSession, logger *utils.Logger)
	RawPacketsFunction func(packet server.RawPacket, logger *utils.Logger)
}

// NewNetworkAdapter returns a new Network adapter to adapt to the RakNet server.
func NewNetworkAdapter(logger *utils.Logger, config resources.GoMineConfig, sessionManager *SessionManager) *NetworkAdapter {
	var rakServer = server.NewGoRakLibServer(config.ServerName, config.ServerIp, config.ServerPort)
	rakServer.SetMinecraftProtocol(info.LatestProtocol)
	rakServer.SetMinecraftVersion(info.LatestGameVersionNetwork)
	rakServer.SetMaxConnectedSessions(config.MaximumPlayers)
	rakServer.SetDefaultGameMode("Creative")
	rakServer.SetMotd(config.ServerMotd)

	var adapter = &NetworkAdapter{logger, rakServer, protocol2.NewManager(), sessionManager, func(session *MinecraftSession, logger *utils.Logger) {}, func(packet server.RawPacket, logger *utils.Logger) {}}

	return adapter
}

// GetProtocolPool returns the protocol pool of the network adapter.
func (adapter *NetworkAdapter) GetProtocolManager() protocol2.Manager {
	return adapter.protocolManager
}

// GetRakLibServer returns the GoRakLib server.
func (adapter *NetworkAdapter) GetRakLibServer() *server.GoRakLibServer {
	return adapter.rakLibServer
}

// Tick ticks the adapter, ticking the GoRakLib server and processing packets.
func (adapter *NetworkAdapter) Tick() {
	go adapter.rakLibServer.Tick()

	for _, session := range adapter.rakLibServer.GetSessionManager().GetSessions() {
		go func(session *server.Session) {
			var (
				minecraftSession *MinecraftSession
				ok               bool
			)
			if minecraftSession, ok = adapter.sessionManager.GetSessionByRakNetSession(session); !ok {
				minecraftSession = NewMinecraftSession(adapter, session)
			}

			adapter.HandlePackets(minecraftSession)
		}(session)
	}

	for _, pk := range adapter.rakLibServer.GetRawPackets() {
		adapter.RawPacketsFunction(pk, adapter.logger)
	}

	for _, session := range adapter.rakLibServer.GetSessionManager().GetDisconnectedSessions() {
		if minecraftSession, ok := adapter.sessionManager.GetSessionByRakNetSession(session); ok {
			adapter.DisconnectFunction(minecraftSession, adapter.logger)
		}
	}
}

// HandlePackets handles all packets of the given session + player.
func (adapter *NetworkAdapter) HandlePackets(session *MinecraftSession) {
	for _, encapsulatedPacket := range session.GetSession().GetReadyEncapsulatedPackets() {
		batch := NewMinecraftPacketBatch(session, adapter.logger)
		batch.Buffer = encapsulatedPacket.Buffer
		batch.Decode()

		for _, packet := range batch.GetPackets() {
			if session.GetProtocolNumber() < 120 {
				packet.DecodeId()
			} else {
				packet.DecodeHeader()
			}
			packet.Decode()

			session.HandlePacket(packet)
		}
	}
}

// GetSession returns a GoRakLib session by an address and port.
func (adapter *NetworkAdapter) GetSession(address string, port uint16) *server.Session {
	var session, _ = adapter.rakLibServer.GetSessionManager().GetSession(address, port)
	return session
}

// SendPacket sends a packet to the given Minecraft session with the given priority.
func (adapter *NetworkAdapter) SendPacket(pk packets.IPacket, session *MinecraftSession, priority byte) {
	var b = NewMinecraftPacketBatch(session, adapter.logger)
	b.AddPacket(pk)

	adapter.SendBatch(b, session.GetSession(), priority)
}

// SendBatch sends a Minecraft packet batch to the given GoRakLib session with the given priority.
func (adapter *NetworkAdapter) SendBatch(batch *MinecraftPacketBatch, session *server.Session, priority byte) {
	session.SendConnectedPacket(batch, protocol.ReliabilityReliableOrdered, priority)
}

// GetLogger returns the logger of the network adapter.
func (adapter *NetworkAdapter) GetLogger() *utils.Logger {
	return adapter.logger
}
