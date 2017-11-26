package interfaces

type IPlayerFactory interface {
	CreatePlayer()
	GetPlayers() []IPlayer
	GetPlayerByName(string) (IPlayer, error)
	GetPlayerByAddress(string, uint16) (IPlayer, error)
}
