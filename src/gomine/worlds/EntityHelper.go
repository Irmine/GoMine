package worlds

import (
	"gomine/interfaces"
	"gomine/net/packets"
)

type EntityHelper struct {
	dimension interfaces.IDimension
}

func NewEntityHelper() *EntityHelper {
	return &EntityHelper{}
}

/**
 * Spawns an entity to a player.
 */
func (manager *EntityHelper) SpawnEntityTo(entity interfaces.IEntity, player interfaces.IPlayer) {
	var pk = packets.NewAddEntityPacket()
	pk.RuntimeId = entity.GetRuntimeId()
	pk.UniqueId = entity.GetUniqueId()
	pk.EntityType = entity.GetEntityId()
	pk.Position = *entity.GetPosition()
	pk.Motion = *entity.GetMotion()

	player.SendPacket(pk)
	entity.AddViewer(player)
}

/**
 * Spawns another player to a player.
 */
func (manager *EntityHelper) SpawnPlayerTo(player interfaces.IPlayer, receiver interfaces.IPlayer) {
	var pk = packets.NewAddPlayerPacket()
	pk.UUID = player.GetUUID()
	pk.Username = player.GetName()
	pk.EntityUniqueId = player.GetUniqueId()
	pk.EntityRuntimeId = player.GetRuntimeId()
	pk.Position = *player.GetPosition()
	pk.Rotation = *player.GetRotation()
	pk.DisplayName = player.GetDisplayName()
	pk.Platform = player.GetPlatform()

	receiver.SendPacket(pk)
}

/**
 * Despawns an entity from a player.
 */
func (manager *EntityHelper) DespawnEntityFrom(entity interfaces.IEntity, player interfaces.IPlayer) {
	var pk = packets.NewRemoveEntityPacket()
	pk.EntityUniqueId = entity.GetUniqueId()
	player.SendPacket(pk)

	entity.RemoveViewer(player)
}

/**
 * Sends entity meta data
 */
func (manager *EntityHelper) SendEntityData(entity interfaces.IEntity, player interfaces.IPlayer) {
	pk := packets.NewSetEntityDataPacket()
	pk.RuntimeId = entity.GetRuntimeId()
	pk.EntityData = entity.GetEntityData()
	player.SendPacket(pk)
}