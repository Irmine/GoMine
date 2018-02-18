package data

const (
	MojangPublicKey = "MHYwEAYHKoZIzj0CAQYFK4EEACIDYgAE8ELkixyLcwlZryUQcu1TvPOmI2B7vX83ndnWRUaXm74wFfa5f/lwQNTfrLVHa2PmenpGI6JhIMUJaWZrjmMj90NoKNFSNBuKdm8rYiXsfaz3K36x/1U26HpG0ZxK/V1V"
)

const (
	StatusLoginSuccess = iota
	StatusLoginFailedClient
	StatusLoginFailedServer
	StatusSpawn
	StatusLoginFailedInvalidTenant
	StatusLoginFailedVanillaEdu
	StatusLoginFailedEduVanilla
)

const (
	MoveNormal = iota
	MoveReset
	MoveTeleport
	MovePitch
)

const (
	StatusRefused = iota + 1
	StatusSendPacks
	StatusHaveAllPacks
	StatusCompleted
)

const (
	TextRaw = iota
	TextChat
	TextTranslation
	TextPopup
	TextJukeboxPopup
	TextTip
	TextSystem
	TextWhisper
	TextAnnouncement
)

const (
	ResourcePackChunkSize = 1048576
)

const (
	ListTypeAdd = iota
	ListTypeRemove
)
