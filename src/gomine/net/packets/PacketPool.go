package packets

import "gomine/net"

type PacketPool struct {
	packets map[int]IPacket
}

func NewPacketPool() *PacketPool {
	var pool = PacketPool{}
	pool.packets = make(map[int]IPacket)

	pool.RegisterPacket(net.LoginPacket, NewLoginPacket())
	pool.RegisterPacket(net.PlayStatusPacket, NewPlayStatusPacket())
	pool.RegisterPacket(net.ClientHandshake, NewClientHandshakePacket())
	pool.RegisterPacket(net.ServerHandshake, NewServerHandshakePacket())

	return &pool
}

func (pool *PacketPool) RegisterPacket(id int, packet IPacket) {
	pool.packets[id] = packet
}

func (pool *PacketPool) GetPacket(id int) IPacket {
	return pool.packets[id]
}