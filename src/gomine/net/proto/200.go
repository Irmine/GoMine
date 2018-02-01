package proto

import (
	"gomine/interfaces"
	"gomine/net/info"
	"gomine/net/packets/p200"
	"gomine/vectors"
	"gomine/entities/math"
)

type Protocol200 struct {
	*Protocol
}

func NewProtocol200() Protocol200 {
	var ids = info.PacketIds200
	return Protocol200{NewProtocol(200, map[int]func() interfaces.IPacket {
		ids[info.LoginPacket]:						func() interfaces.IPacket { return p200.NewLoginPacket() },
		ids[info.ClientHandshakePacket]:			func() interfaces.IPacket { return p200.NewClientHandshakePacket() },
		ids[info.ResourcePackClientResponsePacket]:	func() interfaces.IPacket { return p200.NewResourcePackClientResponsePacket() },
		ids[info.RequestChunkRadiusPacket]:			func() interfaces.IPacket { return p200.NewRequestChunkRadiusPacket() },
		ids[info.MovePlayerPacket]:					func() interfaces.IPacket { return p200.NewMovePlayerPacket() },
		ids[info.CommandRequestPacket]:				func() interfaces.IPacket { return p200.NewCommandRequestPacket() },
		ids[info.ResourcePackChunkRequestPacket]:	func() interfaces.IPacket { return p200.NewResourcePackChunkRequestPacket() },
		ids[info.TextPacket]:						func() interfaces.IPacket { return p200.NewTextPacket() },
		ids[info.PlayerListPacket]:					func() interfaces.IPacket { return p200.NewPlayerListPacket() },
	})}
}

func (protocol *Protocol200) GetAddEntity(entity interfaces.IEntity) *p200.AddEntityPacket {
	var pk = p200.NewAddEntityPacket()
	pk.UniqueId = entity.GetUniqueId()
	pk.RuntimeId = entity.GetRuntimeId()
	pk.EntityType = entity.GetEntityId()
	pk.Position = *entity.GetPosition()
	pk.Motion = *entity.GetMotion()
	pk.Rotation = *entity.GetRotation()
	pk.Attributes = entity.GetAttributeMap()
	pk.EntityData = entity.GetEntityData()

	return pk
}

func (protocol *Protocol200) GetAddPlayer(player interfaces.IPlayer) *p200.AddPlayerPacket {
	var pk = p200.NewAddPlayerPacket()
	
	return pk
}

func (protocol *Protocol200) GetChunkRadiusUpdated(radius int32) *p200.ChunkRadiusUpdatedPacket {
	var pk = p200.NewChunkRadiusUpdatedPacket()
	pk.Radius = radius
	return pk
}

func (protocol *Protocol200) GetCraftingData() *p200.CraftingDataPacket {
	var pk = p200.NewCraftingDataPacket()
	return pk
}

func (protocol *Protocol200) GetDisconnect(message string, hideDisconnectScreen bool) *p200.DisconnectPacket {
	var pk = p200.NewDisconnectPacket()
	pk.HideDisconnectionScreen = hideDisconnectScreen
	pk.Message = message
	return pk
}

func (protocol *Protocol200) GetFullChunkData(chunk interfaces.IChunk) *p200.FullChunkDataPacket {
	var pk = p200.NewFullChunkDataPacket()
	pk.Chunk = chunk
	return pk
}

func (protocol *Protocol200) GetMovePlayer(runtimeId uint64, position vectors.TripleVector, rotation math.Rotation, mode byte, onGround bool, ridingRuntimeId uint64) *p200.MovePlayerPacket {
	var pk = p200.NewMovePlayerPacket()
	pk.RuntimeId = runtimeId
	pk.Position = position
	pk.Rotation = rotation
	pk.Mode = mode
	pk.OnGround = onGround
	pk.RidingRuntimeId = ridingRuntimeId

	return pk
}

func (protocol *Protocol200) GetPlayerList(listType byte, players map[string]interfaces.IPlayer) *p200.PlayerListPacket {
	var pk = p200.NewPlayerListPacket()
	pk.ListType = listType
	pk.Players = players
	return pk
}

func (protocol *Protocol200) GetPlayStatus(status int32) *p200.PlayStatusPacket {
	var pk = p200.NewPlayStatusPacket()
	pk.Status = status
	return pk
}

func (protocol *Protocol200) GetRemoveEntity(uniqueId int64) *p200.RemoveEntityPacket {
	var pk = p200.NewRemoveEntityPacket()
	pk.EntityUniqueId = uniqueId
	return pk
}

func (protocol *Protocol200) GetResourcePackChunkData() *p200.ResourcePackChunkDataPacket {

}