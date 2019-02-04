package net

import (
	"github.com/irmine/gomine/net/packets"
	protocol2 "github.com/irmine/gomine/net/protocol"
	"github.com/irmine/gomine/text"
	"github.com/irmine/goraklib/protocol"
	"github.com/irmine/goraklib/server"
	"net"
)

type NetworkAdapter struct {
	rakLibManager   *server.Manager
	packetManager   protocol2.IPacketManager
	sessionManager  *SessionManager
}

// NewNetworkAdapter returns a new Network adapter to adapt to the RakNet server.
func NewNetworkAdapter(packetManager protocol2.IPacketManager, sessionManager *SessionManager) *NetworkAdapter {
	var manager = server.NewManager()
	var adapter = &NetworkAdapter{manager, packetManager, sessionManager}

	manager.PacketFunction = func(packet []byte, session *server.Session) {
		var minecraftSession *MinecraftSession
		var ok bool
		if minecraftSession, ok = adapter.sessionManager.GetSessionByRakNetSession(session); !ok {
			minecraftSession = NewMinecraftSession(adapter, session)
		}
		adapter.HandlePacket(minecraftSession, packet)
	}
	manager.DisconnectFunction = func(session *server.Session) {
		text.DefaultLogger.Debug(session, "disconnected!")

	}
	manager.ConnectFunction = func(session *server.Session) {
		text.DefaultLogger.Debug(session, "connected!")
	}
	return adapter
}

// GetRakLibManager returns the GoRakLib manager of the network adapter.
func (adapter *NetworkAdapter) GetRakLibManager() *server.Manager {
	return adapter.rakLibManager
}

// HandlePackets handles all packets of the given session + player.
func (adapter *NetworkAdapter) HandlePacket(session *MinecraftSession, buffer []byte) {
	batch := NewMinecraftPacketBatch(session)
	batch.Buffer = buffer
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

// GetSession returns a GoRakLib session by an address and port.
func (adapter *NetworkAdapter) GetSession(address string, port uint16) *server.Session {
	var session, _ = adapter.rakLibManager.Sessions.GetSession(&net.UDPAddr{IP: net.ParseIP(address), Port: int(port)})
	return session
}

// SendPacket sends a packet to the given Minecraft session with the given priority.
func (adapter *NetworkAdapter) SendPacket(pk packets.IPacket, session *MinecraftSession, priority server.Priority) {
	var b = NewMinecraftPacketBatch(session)
	b.AddPacket(pk)

	adapter.SendBatch(b, session.GetSession(), priority)
}

// SendBatch sends a Minecraft packet batch to the given GoRakLib session with the given priority.
func (adapter *NetworkAdapter) SendBatch(batch *MinecraftPacketBatch, session *server.Session, priority server.Priority) {
	session.SendPacket(batch, protocol.ReliabilityReliableOrdered, priority)
}
