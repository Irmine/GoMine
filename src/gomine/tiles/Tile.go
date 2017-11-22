package tiles

var TId uint64 = 0

type Tile struct {
	Name string
	closed bool
	tId uint64
}

func NewTile(Name string) Tile {
	TId++
	return Tile{Name, false, TId}
}

func (tile *Tile) GetName() string {
	return tile.Name
}

func (tile *Tile) SetName(name string) {
	tile.Name = name
}

func (tile *Tile) Close() {
	tile.closed = true
	//todo
}

func (tile *Tile) IsClosed() bool {
	return tile.closed
}

func (tile *Tile) GetId() uint64 {
	return tile.tId
}