package packets

type DataPacket interface {
	Encode()
	Decode()
}
