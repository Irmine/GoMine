package query

import (
	"strconv"
	"strings"

	"github.com/irmine/goraklib/server"
	"github.com/irmine/binutils"
)

const (
	QueryChallenge  = 0x09
	QueryStatistics = 0x00
)

// QueryHeader is the header of each query.
var QueryHeader = []byte{0xfe, 0xfd}

// Query is used to encode/decode queries.
type Query struct {
	*binutils.Stream
	Address string
	Port    uint16

	Header  byte
	QueryId int32
	Token   []byte

	Statistics []byte

	IsShort bool
	Data    []byte
}

// NewQueryFromRaw returns a query from a raw packet.
func NewQueryFromRaw(packet server.RawPacket) *Query {
	var stream = binutils.NewStream()
	stream.Buffer = packet.Buffer
	return &Query{stream, packet.Address, packet.Port, 0, 0, []byte{}, []byte{}, false, []byte{}}
}

// NewQuery returns a new query with an address and port.
func NewQuery(address string, port uint16) *Query {
	return &Query{binutils.NewStream(), address, port, 0, 0, []byte{}, []byte{}, false, []byte{}}
}

// DecodeServer decodes the query sent by the client.
func (query *Query) DecodeServer() {
	query.Offset = 2
	query.Header = query.GetByte()
	query.QueryId = query.GetInt()

	if query.Header == QueryStatistics {
		query.Token = query.Get(4)
		var length = len(query.Get(-1)) + 4 // Token size + padding
		if length != 8 {
			query.IsShort = true
		}
	}
}

// EncodeServer encodes the query to send to the client.
func (query *Query) EncodeServer() {
	query.PutByte(query.Header)
	query.PutInt(query.QueryId)

	switch query.Header {
	case QueryChallenge:
		var token = query.Token
		var offset = 0
		var tokenString = strconv.Itoa(int(binutils.ReadInt(&token, &offset)))

		var padding = 12 - len(tokenString)

		query.PutBytes([]byte(tokenString))
		for i := 0; i < padding; i++ {
			query.PutByte(0)
		}
	case QueryStatistics:
		query.PutBytes(query.Statistics)
		query.PutByte(0)
	}
}

// EncodeClient encodes a query to send to the server.
func (query *Query) EncodeClient() {
	query.PutBytes(QueryHeader)
	query.PutByte(query.Header)
	query.PutInt(query.QueryId)

	if query.Header == QueryStatistics {
		query.PutBytes(query.Token)
		query.PutBytes([]byte{0, 0, 0, 0})
	}
}

// DecodeClient decodes a query sent by the server.
func (query *Query) DecodeClient() {
	query.Header = query.GetByte()
	query.QueryId = query.GetInt()

	switch query.Header {
	case QueryChallenge:
		var buf []byte
		var i, _ = strconv.ParseInt(strings.TrimRight(string(query.Get(-1)), "\x00"), 0, 32)

		binutils.WriteInt(&buf, int32(i))
		query.Token = buf
	case QueryStatistics:
		query.Data = query.Get(-1)
	}
}
