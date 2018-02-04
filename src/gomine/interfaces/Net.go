package interfaces

import (
	"goraklib/server"
	"gomine/utils"
	"gomine/net/info"
	"gomine/net/packets/p200"
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

	GetAddEntity(IEntity) *p200.AddEntityPacket
	GetAddPlayer(IPlayer) *p200.AddPlayerPacket
	GetChunkRadiusUpdated(int32) *p200.ChunkRadiusUpdatedPacket
	GetCraftingData() *p200.CraftingDataPacket
	GetDisconnect(string, bool) *p200.DisconnectPacket
	GetFullChunkData(IChunk) *p200.FullChunkDataPacket
	GetMovePlayer(uint64, vectors.TripleVector, math.Rotation, byte, bool, uint64) *p200.MovePlayerPacket
	GetPlayerList(byte, map[string]IPlayer) *p200.PlayerListPacket
	GetPlayStatus(int32) *p200.PlayStatusPacket
	GetRemoveEntity(int64) *p200.RemoveEntityPacket
	GetResourcePackChunkData(string, int32, int64, []byte) *p200.ResourcePackChunkDataPacket
	GetResourcePackDataInfo(IPack) *p200.ResourcePackDataInfoPacket
	GetResourcePackInfo(bool, []IPack, []IPack) *p200.ResourcePackInfoPacket
	GetResourcePackStack(bool, []IPack, []IPack) *p200.ResourcePackStackPacket
	GetServerHandshake(string) *p200.ServerHandshakePacket
	GetSetEntityData(IEntity, map[uint32][]interface{}) *p200.SetEntityDataPacket
	GetStartGame(IPlayer) *p200.StartGamePacket
	GetText(types.Text) *p200.TextPacket
	GetTransfer(string, uint16) *p200.TransferPacket
	GetUpdateAttributes(IEntity, *data.AttributeMap) *p200.UpdateAttributesPacket
}

type IProtocolPool interface {
	GetProtocol(int32) IProtocol
	RegisterProtocol(IProtocol)
	IsProtocolRegistered(int32) bool
	DeregisterProtocol(int32)
}