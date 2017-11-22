package utils

import (
	"fmt"
	"strings"
	"encoding/base64"
	"encoding/json"
)

func DecodeJwt(v string, t interface{}) {
	v = strings.Split(v, ".")[1]
	v = strings.Replace(v, "-_", "+/", -1)
	str, err := base64.StdEncoding.DecodeString(v)
	if err != nil {
		fmt.Printf("An error occurred while decoding base64 chain data : %v", err)
		return
	}
	json.Unmarshal(str, t)
}