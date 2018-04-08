package items

import (
	"fmt"
	"github.com/irmine/gomine/text"
	"strconv"
	"strings"
)

// IdToState is a map used to convert
// an ID + item data combination to item type.
// The keys of these maps are created using the
// getKey method.
var IdToType = map[string]Type{
	GetKey(0, 0): DefaultManager.stringIds["minecraft:air"],
	GetKey(1, 0): DefaultManager.stringIds["minecraft:stone"],
}

// TypeToId is a map used to convert
// a block state to an ID + data combination.
var TypeToId = map[string]string{
	fmt.Sprint(DefaultManager.stringIds["minecraft:air"]):   GetKey(0, 0),
	fmt.Sprint(DefaultManager.stringIds["minecraft:stone"]): GetKey(1, 0),
}

// getKey returns the key of an ID + data combination,
// which is used in both maps.
func GetKey(id int16, data int16) string {
	return fmt.Sprint(id, ":", data)
}

// FromKey attempts to retrieve an ID + data combination,
// from a string created with getKey.
// Any errors that occur are logged to the default logger.
func FromKey(key string) (int16, int16) {
	fragments := strings.Split(key, ":")
	idFrag, dataFrag := fragments[0], fragments[1]
	i, err := strconv.Atoi(idFrag)
	text.DefaultLogger.LogError(err)
	d, err := strconv.Atoi(dataFrag)
	text.DefaultLogger.LogError(err)
	return int16(i), int16(d)
}
