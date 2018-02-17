package query

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/irmine/binutils"
)

const (
	GameId = "MINECRAFTPE"
)

var writeIndexes = []string{
	"splitnum",
	"hostname",
	"gametype",
	"game_id",
	"version",
	"server_engine",
	"plugins",
	"map",
	"numplayers",
	"maxplayers",
	"whitelist",
	"hostip",
	"hostport",
}
var readIndexes = map[string]string{
	"hostname":      "MOTD",
	"gametype":      "GameMode",
	"version":       "Version",
	"server_engine": "ServerEngine",
	"plugins":       "PluginNames",
	"map":           "WorldName",
	"numplayers":    "OnlinePlayers",
	"maxplayers":    "MaximumPlayers",
	"whitelist":     "Whitelist",
}

type QueryResult struct {
	MOTD           string
	ListPlugins    bool
	PluginNames    []string
	PlayerNames    []string
	GameMode       string
	Version        string
	ServerEngine   string
	WorldName      string
	OnlinePlayers  int
	MaximumPlayers int
	Whitelist      string
	Port           uint16
	Address        string
}

func (result QueryResult) GetLong() []byte {
	var plugs = result.ServerEngine
	if result.ListPlugins {
		plugs += ":"
		for _, plugin := range result.PluginNames {
			plugs += " " + plugin + ";"
		}
		plugs = strings.TrimRight(plugs, ";")
	}

	var queryData = map[string][]byte{
		"splitnum":      {0x80},
		"hostname":      []byte(result.MOTD),
		"gametype":      []byte(result.GameMode),
		"game_id":       []byte(GameId),
		"version":       []byte(result.Version),
		"server_engine": []byte(result.ServerEngine),
		"plugins":       []byte(plugs),
		"map":           []byte(result.WorldName),
		"numplayers":    []byte(strconv.Itoa(result.OnlinePlayers)),
		"maxplayers":    []byte(strconv.Itoa(result.MaximumPlayers)),
		"whitelist":     []byte(result.Whitelist),
		"hostip":        []byte(result.Address),
		"hostport":      []byte(strconv.Itoa(int(result.Port))),
	}

	var stream = binutils.NewStream()

	// Query data should not have to be ordered, because all values are already prefixed with their keys.
	// Many implementations do not work with unordered data however, so therefore still order it.
	for _, index := range writeIndexes {
		stream.PutBytes([]byte(index))
		stream.PutByte(0)
		stream.PutBytes(queryData[index])
		stream.PutByte(0)
	}

	stream.PutBytes([]byte{0x00, 0x01})
	stream.PutBytes([]byte("player_"))
	stream.PutBytes([]byte{0x00, 0x00})

	for _, name := range result.PlayerNames {
		stream.PutBytes([]byte(name))
		stream.PutByte(0)
	}

	stream.PutByte(0)

	return stream.Buffer
}

func (result QueryResult) ParseLong(data []byte) QueryResult {
	var r = &QueryResult{}
	var str = string(data)
	var arr = strings.Split(str, "\x00")
	var m = map[string]interface{}{}

	for i := 0; i < len(arr); i += 2 {
		key := ""
		ok := true
		if key, ok = readIndexes[arr[i]]; !ok {
			continue
		}
		var val interface{}

		stringVal := arr[i+1]

		if key == "PluginNames" {
			stringVal = strings.Split(stringVal, ":")[1]
			stringVal = strings.TrimLeft(stringVal, " ")
			val = strings.Split(stringVal, "; ")
		} else if key == "OnlinePlayers" || key == "MaximumPlayers" {
			i, _ := strconv.ParseInt(stringVal, 0, 32)
			val = int(i)
		} else {
			val = stringVal
		}
		m[key] = val
	}

	var v = reflect.ValueOf(r).Elem()
	for key, value := range m {
		v.FieldByName(key).Set(reflect.ValueOf(value))
	}

	var players = strings.Split(str, "player_")[1]
	var playerArrayRaw = strings.Split(players, "\x00")
	var playerArray []string
	for _, name := range playerArrayRaw {
		if len(name) < 3 {
			continue
		}
		playerArray = append(playerArray, name)
	}

	r.PlayerNames = playerArray

	return *r
}

func (result QueryResult) GetShort() []byte {
	var strs = []string{result.MOTD, result.GameMode, result.WorldName, strconv.Itoa(result.OnlinePlayers), strconv.Itoa(result.MaximumPlayers)}
	var str = strings.Join(strs, " ")

	var stream = binutils.NewStream()
	stream.Buffer = []byte(str)
	stream.PutByte(0)
	stream.PutLittleShort(int16(result.Port))
	stream.PutByte(0)
	stream.PutBytes([]byte(result.Address))
	stream.PutByte(0)

	return stream.Buffer
}
