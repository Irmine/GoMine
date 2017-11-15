package net

import (
	"gomine/interfaces"
	server2 "goraklib/server"
)

type GoRakLibAdapter struct {
	server interfaces.IServer
	rakLibServer *server2.GoRakLibServer
}

func NewGoRakLibAdapter(server interfaces.IServer) *GoRakLibAdapter {
	var rakServer = server2.NewGoRakLibServer(server.GetName(), server.GetAddress(), server.GetPort())
	rakServer.SetMinecraftProtocol(LatestProtocol)
	rakServer.SetMinecraftVersion(GameVersionNetwork)
	rakServer.SetServerName(server.GetName())
	rakServer.SetMaxConnectedSessions(server.GetMaximumPlayers())
	rakServer.SetConnectedSessionCount(0)
	rakServer.SetDefaultGameMode("Creative")
	rakServer.SetMotd(server.GetMotd())

	return &GoRakLibAdapter{server, rakServer}
}

func (adapter *GoRakLibAdapter) Tick() {
	go adapter.rakLibServer.Tick()
}