package packets

import "gomine/net/info"

type ResourcePackClientResponsePacket struct {
	*Packet
	Status byte
	ResourcePackIds []string
}

func NewResourcePackClientResponsePacket() ResourcePackClientResponsePacket {
	return ResourcePackClientResponsePacket{NewPacket(info.ResourcePackClientResponsePacket), 0, []string{}}
}

func (pk ResourcePackClientResponsePacket) Encode()  {

}

func (pk ResourcePackClientResponsePacket) Decode()  {
	pk.Status = pk.GetByte()
	var idCount = pk.GetLittleShort()
	for i := int16(0); i < idCount; i++ {
		pk.ResourcePackIds = append(pk.ResourcePackIds, pk.GetString())
	}
}

