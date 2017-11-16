package players

import "gomine/interfaces"

type PlayerFactory struct {
	server interfaces.IServer
	players map[string]Player
}

func NewPlayerFactory(server interfaces.IServer) *PlayerFactory {
	return &PlayerFactory{server, make(map[string]Player)}
}