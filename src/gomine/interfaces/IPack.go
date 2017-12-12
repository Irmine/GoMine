package interfaces

type IPack interface {
	GetUUID() string
	GetVersion() string
	GetFileSize() int64
	GetSha256() string
	GetChunk(offset int, length int) []byte
}