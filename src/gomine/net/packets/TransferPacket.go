package packets

import "gomine/net/info"

type TransferPacket struct {
	*Packet
	Address string
	Port uint16
}

func NewTransferPacket() *TransferPacket {
	return &TransferPacket{NewPacket(info.TransferPacket), "", 0}
}

func (pk *TransferPacket) Encode() {
	pk.PutString(pk.Address)
	pk.PutLittleShort(int16(pk.Port))
}

func (pk *TransferPacket) Decode() {

}