package packets

import (
	"gomine/net/info"
	"gomine/vectors"
	"encoding/base64"
)

type StartGamePacket struct {
	*Packet
}

func NewStartGamePacket() *StartGamePacket {
	return &StartGamePacket{NewPacket(info.StartGamePacket)}
}

func (pk *StartGamePacket) Encode()  {
	pk.PutVarLong(1)
	pk.PutVarInt(1)
	pk.PutVarInt(1)

	pk.PutTripleVectorObject(vectors.TripleVector{0, 20, 0})

	pk.PutLittleFloat(0)
	pk.PutLittleFloat(0)

	pk.PutVarInt(321050)
	pk.PutVarInt(0)
	pk.PutVarInt(0)
	pk.PutVarInt(1)
	pk.PutVarInt(1)
	pk.PutBlockPos(0, 80, 0)
	pk.PutBool(false)
	pk.PutVarInt(0)
	pk.PutBool(false)
	pk.PutLittleFloat(0)
	pk.PutLittleFloat(0)
	pk.PutBool(true)
	pk.PutBool(true)
	pk.PutBool(true)
	pk.PutBool(true)
	pk.PutBool(false)
	pk.PutUnsignedVarInt(0)
	pk.PutBool(false)
	pk.PutBool(false)
	pk.PutBool(true)
	pk.PutVarInt(1)
	pk.PutVarInt(0)

	pk.PutString(base64.RawStdEncoding.EncodeToString([]byte("world")))
	pk.PutString("world")
	pk.PutString("")
	pk.PutBool(false)
	pk.PutLittleLong(100)
	pk.PutVarInt(312904)
}

func (pk *StartGamePacket) Decode()  {

}