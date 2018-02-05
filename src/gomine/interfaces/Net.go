package interfaces

import (
	"goraklib/server"
	"gomine/utils"
	"gomine/net/info"
	"gomine/entities/data"
	"gomine/vectors"
	"gomine/entities/math"
	"gomine/net/packets/types"
)

type IPacketHandler interface {
	Handle(IPacket, IPlayer, *server.Session, IServer) bool
	SetPriority(int) bool
	GetPriority() int
}

type IPacket interface {
	SetBuffer([]byte)
	GetBuffer() []byte
	GetId() int
	SetId(int)
	EncodeHeader()
	Encode()
	DecodeHeader()
	Decode()
	ResetStream()
	GetOffset() int
	SetOffset(int)
	Discard()
	IsDiscarded() bool
	EncodeId()
	DecodeId()
}

type INetworkAdapter interface {
	GetSession(string, uint16) *server.Session
	SendBatch(IMinecraftPacketBatch, *server.Session, byte)
	SendPacket(IPacket, IMinecraftSession, byte)
	Tick()
	GetRakLibServer() *server.GoRakLibServer
	GetProtocolPool() IProtocolPool
}

type IMinecraftPacketBatch interface {
	GetPackets() []IPacket
	AddPacket(IPacket)
	Encode()
	Decode()
	GetBuffer() []byte
}

type IMinecraftSession interface {
	GetServer() IServer
	GetSession() *server.Session
	GetPing() uint64
	GetUUID() utils.UUID
	GetXUID() string
	SetLanguage(string)
	GetLanguage() string
	GetClientId() int
	GetEncryptionHandler() *utils.EncryptionHandler
	UsesEncryption() bool
	EnableEncryption()
	IsXBOXLiveAuthenticated() bool
	SetXBOXLiveAuthenticated(bool)
	SendPacket(IPacket)
	SendBatch(IMinecraftPacketBatch)
	IsInitialized() bool
	GetPlatform() int32
	GetProtocolNumber() int32
	GetProtocol() IProtocol
	SetProtocol(IProtocol)
	GetGameVersion() string
	HandlePacket(IPacket, IPlayer)

	SendAddEntity(IEntity)
	SendAddPlayer(IPlayer)
	SendChunkRadiusUpdated(int32)
	SendCraftingData()
	SendDisconnect(string, bool)
	SendFullChunkData(IChunk)
	SendMovePlayer(IPlayer, vectors.TripleVector, math.Rotation, byte, bool, uint64)
	SendPlayerList(byte, map[string]IPlayer)
	SendPlayStatus(int32)
	SendRemoveEntity(IEntity)
	SendResourcePackChunkData(string, int32, int64, []byte)
	SendResourcePackDataInfo(IPack)
	SendResourcePackInfo(bool, []IPack, []IPack)
	SendResourcePackStack(bool, []IPack, []IPack)
	SendServerHandshake(string)
	SendSetEntityData(IEntity, map[uint32][]interface{})
	SendStartGame(IPlayer)
	SendText(types.Text)
	Transfer(string, uint16)
	SendUpdateAttributes(IEntity, *data.AttributeMap)
}

type IProtocol interface {
	GetProtocolNumber() int32
	GetPackets() map[int]func() IPacket
	RegisterPacket(int, func() IPacket)
	GetPacket(int) IPacket
	IsPacketRegistered(int) bool
	GetHandlers(info.PacketName) [][]IPacketHandler
	RegisterHandler(info.PacketName, IPacketHandler, int) bool
	DeregisterPacketHandlers(info.PacketName, int)
	GetIdList() info.PacketIdList
	GetHandlersById(int) [][]IPacketHandler

	GetAddEntity(IEntity) IPacket
	GetAddPlayer(IPlayer) IPacket
	GetChunkRadiusUpdated(int32) IPacket
	GetCraftingData() IPacket
	GetDisconnect(string, bool) IPacket
	GetFullChunkData(IChunk) IPacket
	GetMovePlayer(uint64, vectors.TripleVector, math.Rotation, byte, bool, uint64) IPacket
	GetPlayerList(byte, map[string]IPlayer) IPacket
	GetPlayStatus(int32) IPacket
	GetRemoveEntity(int64) IPacket
	GetResourcePackChunkData(string, int32, int64, []byte) IPacket
	GetResourcePackDataInfo(IPack) IPacket
	GetResourcePackInfo(bool, []IPack, []IPack) IPacket
	GetResourcePackStack(bool, []IPack, []IPack) IPacket
	GetServerHandshake(string) IPacket
	GetSetEntityData(IEntity, map[uint32][]interface{}) IPacket
	GetStartGame(IPlayer) IPacket
	GetText(types.Text) IPacket
	GetTransfer(string, uint16) IPacket
	GetUpdateAttributes(IEntity, *data.AttributeMap) IPacket
}

type IProtocolPool interface {
	GetProtocol(int32) IProtocol
	RegisterProtocol(IProtocol)
	IsProtocolRegistered(int32) bool
	DeregisterProtocol(int32)
}