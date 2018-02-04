package types

import "gomine/utils"

type PlayerListEntry struct {
	UUID utils.UUID
	XUID string
	EntityUniqueId int64
	Username string
	DisplayName string
	Platform int32
	SkinId string
	SkinData []byte
	CapeData []byte
	GeometryName string
	GeometryData string
}

type SessionData struct {
	ClientUUID utils.UUID
	ClientXUID string
	ClientId int
	ProtocolNumber int32
	GameVersion string
	Language string
	DeviceOS int
}

type Text struct {
	Message string
	SourceName string
	SourceDisplayName string
	SourcePlatform int32
	SourceXUID string
	TextType byte
	IsTranslation bool
	TranslationParameters []string
}