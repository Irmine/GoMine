package info

const (
	LatestProtocol    		 = 200
	LatestGameVersion        = "v1.2.10.1"
	LatestGameVersionNetwork = "1.2.10.1"
)

type PacketIdList map[PacketName]int

type PacketName string

const (
	LoginPacket PacketName = "LoginPacket"
	PlayStatusPacket PacketName = "PlayStatusPacket"
	ServerHandshakePacket PacketName = "ServerHandshakePacket"
	ClientHandshakePacket PacketName = "ClientHandshakePacket"
	DisconnectPacket PacketName = "DisconnectPacket"
	ResourcePackInfoPacket PacketName = "ResourcePackInfoPacket"
	ResourcePackStackPacket PacketName = "ResourcePackStackPacket"
	ResourcePackClientResponsePacket PacketName = "ResourcePackClientResponsePacket"
	TextPacket PacketName = "TextPacket"
	SetTimePacket PacketName = "SetTimePacket"
	StartGamePacket PacketName = "StartGamePacket"
	AddPlayerPacket PacketName = "AddPlayerPacket"
	AddEntityPacket PacketName = "AddEntityPacket"
	RemoveEntityPacket PacketName = "RemoveEntityPacket"
	AddItemEntityPacket PacketName = "AddItemEntityPacket"
	AddHangingEntityPacket PacketName = "AddHangingEntityPacket"
	TakeItemEntityPacket PacketName = "TakeItemEntityPacket"
	MoveEntityPacket PacketName = "MoveEntityPacket"
	MovePlayerPacket PacketName = "MovePlayerPacket"
	RiderJumpPacket PacketName = "RiderJumpPacket"
	UpdateBlockPacket PacketName = "UpdateBlockPacket"
	AddPaintingPacket PacketName = "AddPaintingPacket"
	ExplodePacket PacketName = "ExplodePacket"
	LevelSoundEventPacket PacketName = "LevelSoundEventPacket"
	LevelEventPacket PacketName = "LevelEventPacket"
	BlockEventPacket PacketName = "BlockEventPacket"
	EntityEventPacket PacketName = "EntityEventPacket"
	MobEffectPacket PacketName = "MobEffectPacket"
	UpdateAttributesPacket PacketName = "UpdateAttributesPacket"
	InventoryTransactionPacket PacketName = "InventoryTransactionPacket"
	MobEquipmentPacket PacketName = "MobEquipmentPacket"
	MobArmorEquipmentPacket PacketName = "MobArmorEquipmentPacket"
	InteractPacket PacketName = "InteractPacket"
	BlockPickRequestPacket PacketName = "BlockPickRequestPacket"
	EntityPickRequestPacket PacketName = "EntityPickRequestPacket"
	PlayerActionPacket PacketName = "PlayerActionPacket"
	EntityFallPacket PacketName = "EntityFallPacket"
	HurtArmorPacket PacketName = "HurtArmorPacket"
	SetEntityDataPacket PacketName = "SetEntityDataPacket"
	SetEntityMotionPacket PacketName = "SetEntityMotionPacket"
	SetEntityLinkPacket PacketName = "SetEntityLinkPacket"
	SetHealthPacket PacketName = "SetHealthPacket"
	SetSpawnPositionPacket PacketName = "SetSpawnPositionPacket"
	AnimatePacket PacketName = "AnimatePacket"
	RespawnPacket PacketName = "RespawnPacket"
	ContainerOpenPacket PacketName = "ContainerOpenPacket"
	ContainerClosePacket PacketName = "ContainerClosePacket"
	PlayerHotbarPacket PacketName = "PlayerHotbarPacket"
	InventoryContentPacket PacketName = "InventoryContentPacket"
	InventorySlotPacket PacketName = "InventorySlotPacket"
	ContainerSetDataPacket PacketName = "ContainerSetDataPacket"
	CraftingDataPacket PacketName = "CraftingDataPacket"
	CraftingEventPacket PacketName = "CraftingEventPacket"
	GuiDataPickItemPacket PacketName = "GuiDataPickItemPacket"
	AdventureSettingsPacket PacketName = "AdventureSettingsPacket"
	BlockEntityDataPacket PacketName = "BlockEntityDataPacket"
	PlayerInputPacket PacketName = "PlayerInputPacket"
	FullChunkDataPacket PacketName = "FullChunkDataPacket"
	SetCommandsEnabledPacket PacketName = "SetCommandsEnabledPacket"
	SetDifficultyPacket PacketName = "SetDifficultyPacket"
	ChangeDimensionPacket PacketName = "ChangeDimensionPacket"
	SetPlayerGameTypePacket PacketName = "SetPlayerGameTypePacket"
	PlayerListPacket PacketName = "PlayerListPacket"
	SimpleEventPacket PacketName = "SimpleEventPacket"
	EventPacket PacketName = "EventPacket"
	SpawnExperienceOrbPacket PacketName = "SpawnExperienceOrbPacket"
	ClientboundMapItemDataPacket PacketName = "ClientboundMapItemDataPacket"
	MapInfoRequestPacket PacketName = "MapInfoRequestPacket"
	RequestChunkRadiusPacket PacketName = "RequestChunkRadiusPacket"
	ChunkRadiusUpdatedPacket PacketName = "ChunkRadiusUpdatedPacket"
	ItemFrameDropItemPacket PacketName = "ItemFrameDropItemPacket"
	GameRulesChangedPacket PacketName = "GameRulesChangedPacket"
	CameraPacket PacketName ="CameraPacket"
	BossEventPacket PacketName = "BossEventPacket"
	ShowCreditsPacket PacketName = "ShowCreditsPacket"
	AvailableCommandsPacket PacketName = "AvailableCommandsPacket"
	CommandRequestPacket PacketName = "CommandRequestPacket"
	CommandBlockUpdatePacket PacketName = "CommandBlockUpdatePacket"
	CommandOutputPacket PacketName = "CommandOutputPacket"
	UpdateTradePacket PacketName = "UpdateTradePacket"
	UpdateEquipPacket PacketName = "UpdateEquipPacket"
	ResourcePackDataInfoPacket PacketName = "ResourcePackDataInfoPacket"
	ResourcePackChunkDataPacket PacketName = "ResourcePackChunkDataPacket"
	ResourcePackChunkRequestPacket PacketName = "ResourcePackChunkRequestPacket"
	TransferPacket PacketName = "TransferPacket"
	PlaySoundPacket PacketName = "PlaySoundPacket"
	StopSoundPacket PacketName = "StopSoundPacket"
	SetTitlePacket PacketName = "SetTitlePacket"
	AddBehaviorTreePacket PacketName = "AddBehaviorTreePacket"
	StructureBlockUpdatePacket PacketName = "StructureBlockUpdatePacket"
	ShowStoreOfferPacket PacketName = "ShowStoreOfferPacket"
	PurchaseReceiptPacket PacketName = "PurchaseReceiptPacket"
	PlayerSkinPacket PacketName = "PlayerSkinPacket"
	SubClientLoginPacket PacketName = "SubClientLoginPacket"
	WSConnectPacket PacketName = "WSConnectPacket"
	SetLastHurtByPacket PacketName = "SetLastHurtByPacket"
	BookEditPacket PacketName = "BookEditPacket"
	NpcRequestPacket PacketName = "NpcRequestPacket"
	PhotoTransferPacket PacketName = "PhotoTransferPacket"
	ModalFormRequestPacket PacketName = "ModalFormRequestPacket"
	ModalFormResponsePacket PacketName = "ModalFormResponsePacket"
	ServerSettingsRequestPacket PacketName = "ServerSettingsRequestPacket"
	ServerSettingsResponsePacket PacketName = "ServerSettingsResponsePacket"
	ShowProfilePacket PacketName = "ShowProfilePacket"
	SetDefaultGameTypePacket PacketName = "SetDefaultGameTypePacket"
)