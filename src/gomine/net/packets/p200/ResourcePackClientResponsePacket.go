package p200

import (
	"gomine/net/info"
	"gomine/net/packets"
)

const (
	StatusRefused = iota + 1
	StatusSendPacks
	StatusHaveAllPacks
	StatusCompleted
)

type ResourcePackClientResponsePacket struct {
	*packets.Packet
	Status byte
	PackUUIDs []string
}

func NewResourcePackClientResponsePacket() *ResourcePackClientResponsePacket {
	return &ResourcePackClientResponsePacket{packets.NewPacket(info.ResourcePackClientResponsePacket), 0, []string{}}
}

func (pk *ResourcePackClientResponsePacket) Encode()  {

}

func (pk *ResourcePackClientResponsePacket) Decode()  {
	pk.Status = pk.GetByte()
	var idCount = pk.GetLittleShort()
	for i := int16(0); i < idCount; i++ {
		pk.PackUUIDs = append(pk.PackUUIDs, pk.GetString())
	}
}