package net

import (
	"gomine/interfaces"
	server2 "goraklib/server"
	"gomine/net/packets"
	"gomine/net/info"
	"goraklib/protocol"
	"fmt"
)

type GoRakLibAdapter struct {
	server interfaces.IServer
	rakLibServer *server2.GoRakLibServer
}

func NewGoRakLibAdapter(server interfaces.IServer) *GoRakLibAdapter {
	var rakServer = server2.NewGoRakLibServer(server.GetName(), server.GetAddress(), server.GetPort())
	rakServer.SetMinecraftProtocol(info.LatestProtocol)
	rakServer.SetMinecraftVersion(info.GameVersionNetwork)
	rakServer.SetServerName(server.GetName())
	rakServer.SetMaxConnectedSessions(server.GetMaximumPlayers())
	rakServer.SetConnectedSessionCount(0)
	rakServer.SetDefaultGameMode("Creative")
	rakServer.SetMotd(server.GetMotd())

	packets.InitPacketPool()

	return &GoRakLibAdapter{server, rakServer}
}

func (adapter *GoRakLibAdapter) Tick() {
	go adapter.rakLibServer.Tick()

	go func() {
		for _, session := range adapter.rakLibServer.GetSessionManager().GetSessions() {
			for _, encapsulatedPacket := range session.GetReadyEncapsulatedPackets() {
				batch := NewMinecraftPacketBatch()
				batch.stream.Buffer = encapsulatedPacket.Buffer
				batch.Decode()
				for _, packet := range batch.GetPackets() {
					//packet.Decode()
					if packet.GetId() == info.LoginPacket {
						fmt.Println("LoginPacket received... Answering...")
						var pk = packets.NewPlayStatusPacket()
						pk.Status = 1

						pk.EncodeHeader()
						pk.Encode()

						var b = NewMinecraftPacketBatch()
						b.AddPacket(pk)

						pk.Status = 0
						b.AddPacket(pk)
						adapter.SendBatch(&b, session.GetAddress(), session.GetPort())
					}
				}
			}
		}
	}()
}

func (adapter *GoRakLibAdapter) SendBatch(batch *MinecraftPacketBatch, ip string, port uint16) {
	batch.Encode()

	var encPacket = protocol.NewEncapsulatedPacket()
	encPacket.SetBuffer(batch.stream.GetBuffer())

	var datagram = protocol.NewDatagram()
	datagram.AddPacket(&encPacket)
	datagram.Encode()

	adapter.rakLibServer.GetSessionManager().SendPacket(datagram, ip, port)
}