package p200

import (
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
)

type TransferPacket struct {
	*packets.Packet
	Address string
	Port    uint16
}

func NewTransferPacket() *TransferPacket {
	return &TransferPacket{packets.NewPacket(info.PacketIds200[info.TransferPacket]), "", 0}
}

func (pk *TransferPacket) Encode() {
	pk.PutString(pk.Address)
	pk.PutLittleShort(int16(pk.Port))
}

func (pk *TransferPacket) Decode() {

}
