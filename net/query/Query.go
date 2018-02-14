package query

import (
	"strconv"
	"strings"

	"github.com/irmine/gomine/utils"
	"github.com/irmine/goraklib/server"
)

const (
	QueryChallenge  = 0x09
	QueryStatistics = 0x00
)

var QueryHeader = []byte{0xfe, 0xfd}

type Query struct {
	*utils.BinaryStream
	Address string
	Port    uint16

	Header  byte
	QueryId int32
	Token   []byte

	Statistics []byte

	IsShort bool
	Data    []byte
}

func NewQueryFromRaw(packet server.RawPacket) *Query {
	var stream = utils.NewStream()
	stream.Buffer = packet.Buffer
	return &Query{stream, packet.Address, packet.Port, 0, 0, []byte{}, []byte{}, false, []byte{}}
}

func NewQuery(address string, port uint16) *Query {
	return &Query{utils.NewStream(), address, port, 0, 0, []byte{}, []byte{}, false, []byte{}}
}

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

func (query *Query) EncodeServer() {
	query.PutByte(query.Header)
	query.PutInt(query.QueryId)

	switch query.Header {
	case QueryChallenge:
		var token = query.Token
		var offset = 0
		var tokenString = strconv.Itoa(int(utils.ReadInt(&token, &offset)))

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

func (query *Query) EncodeClient() {
	query.PutBytes(QueryHeader)
	query.PutByte(query.Header)
	query.PutInt(query.QueryId)

	if query.Header == QueryStatistics {
		query.PutBytes(query.Token)
		query.PutBytes([]byte{0, 0, 0, 0})
	}
}

func (query *Query) DecodeClient() {
	query.Header = query.GetByte()
	query.QueryId = query.GetInt()

	switch query.Header {
	case QueryChallenge:
		var buf []byte
		var i, _ = strconv.ParseInt(strings.TrimRight(string(query.Get(-1)), "\x00"), 0, 32)

		utils.WriteInt(&buf, int32(i))
		query.Token = buf
	case QueryStatistics:
		query.Data = query.Get(-1)
	}
}
