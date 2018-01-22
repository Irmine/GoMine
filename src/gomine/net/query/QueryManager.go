package query

import (
	"gomine/interfaces"
	"goraklib/server"
	"math/rand"
	"time"
	"net"
	"strconv"
	"fmt"
)

type QueryManager struct {
	server interfaces.IServer
	token []byte
}

func NewQueryManager(server interfaces.IServer) QueryManager {
	var b = make([]byte, 4)
	rand.Read(b)
	return QueryManager{server, b}
}

func (manager *QueryManager) HandleQuery(query *Query) {
	switch query.Header {
	case QueryChallenge:
		var q = NewQuery(query.Address, query.Port)
		q.Header = QueryChallenge
		q.QueryId = query.QueryId
		q.Token = manager.token

		manager.sendQuery(q)

	case QueryStatistics:
		if string(manager.token) != string(query.Token) {
			return
		}

		var q = NewQuery(query.Address, query.Port)
		q.Header = QueryStatistics
		q.QueryId = query.QueryId
		q.Statistics = manager.server.GenerateQueryResult(query.IsShort)

		manager.sendQuery(q)
	}
}

/**
 * Sends a query to the address and port set in it.
 */
func (manager *QueryManager) sendQuery(query *Query) {
	query.EncodeServer()
	var raw = server.NewRawPacket()
	raw.Buffer = query.Buffer
	raw.Address = query.Address
	raw.Port = query.Port
	manager.server.GetRakLibAdapter().GetRakLibServer().SendRaw(raw)
}

/**
 * Queries a server with the given address and port.
 * The call times out after the given timeout duration if no response is given.
 */
func QueryServer(address string, port uint16, timeout time.Duration) (QueryResult, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var realAddresses, err = net.LookupIP(address)
	if err == nil {
		address = realAddresses[0].String()
	}

	var connection, err2 = net.Dial("udp", address + ":" + strconv.Itoa(int(port)))
	connection.SetReadDeadline(time.Now().Add(timeout))

	if err2 != nil {
		return QueryResult{}, err2
	}

	var q = NewQuery(address, port)
	q.Header = QueryChallenge
	q.QueryId = int32(time.Now().Unix())
	q.EncodeClient()

	connection.Write(q.Buffer)

	var buf = make([]byte, 128)
	var bytesRead, err3 = connection.Read(buf)
	if err3 != nil {
		return QueryResult{}, err3
	}

	buf = buf[:bytesRead]

	q = NewQuery(address, port)
	q.Buffer = buf
	q.DecodeClient()

	var statQuery = NewQuery(address, port)
	statQuery.Header = QueryStatistics
	statQuery.QueryId = int32(time.Now().Unix())
	statQuery.Token = q.Token
	statQuery.EncodeClient()

	connection.Write(statQuery.Buffer)

	buf = make([]byte, 4096)
	var byteCount, err4 = connection.Read(buf)
	if err4 != nil {
		return QueryResult{}, err4
	}
	connection.Close()

	buf = buf[:byteCount]

	q = NewQuery(address, port)
	q.Buffer = buf
	q.DecodeClient()

	var res = QueryResult{}.ParseLong(q.Data)
	res.Port = port
	res.Address = address

	return res, nil
}