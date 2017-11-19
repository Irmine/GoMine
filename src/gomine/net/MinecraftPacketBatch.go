package net

import (
	"gomine/utils"
	"bytes"
	"compress/zlib"
	"io/ioutil"
	"gomine/net/packets"
	"goraklib/protocol"
)

const McpeFlag = 0xFE

type MinecraftPacketBatch struct {
	stream *utils.BinaryStream

	raw []byte
	packets []packets.IPacket
}

func NewMinecraftPacketBatch(packet protocol.EncapsulatedPacket) MinecraftPacketBatch {
	var batch = MinecraftPacketBatch{}
	batch.stream = utils.NewStream()
	batch.stream.Buffer = packet.GetBuffer()

	return batch
}

func (batch *MinecraftPacketBatch) Decode() bool {
	var mcpeFlag = batch.stream.GetByte()
	if mcpeFlag != McpeFlag {
		return false
	}


	var reader = bytes.NewReader(batch.stream.Buffer[batch.stream.Offset:])
	var zlibReader, _ = zlib.NewReader(reader)
	defer zlibReader.Close()

	batch.raw, _ = ioutil.ReadAll(zlibReader)

	batch.stream.ResetStream()
	batch.stream.SetBuffer(batch.raw)

	var packetData [][]byte

	for !batch.stream.Feof() {
		packetData = append(packetData, []byte(batch.stream.GetString()))
	}

	for _, data := range packetData {
		packetId := int(data[0])
		packet := packets.GetPacket(packetId)

		packet.SetBuffer(data)
		batch.packets = append(batch.packets, packet)
	}

	return true
}

func (batch *MinecraftPacketBatch) GetPackets() []packets.IPacket {
	return batch.packets
}
