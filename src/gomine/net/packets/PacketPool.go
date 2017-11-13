package packets

import "gomine/net"

type PacketPool struct {
	pool map[int]DataPacket
}

func NewPacketPool() PacketPool {
	pool := PacketPool{}
	pool.pool[net.ClientHandshake] = ClientHandshakePacket{}
	pool.pool[net.ServerHandshake] = ClientHandshakePacket{}
	return pool
}
