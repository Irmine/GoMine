package utils

import (
	"fmt"
	"strings"
	"encoding/base64"
	"encoding/json"
)

func DecodeJwtPayload(v string, t interface{}) {
	v = strings.Split(v, ".")[1]
	v = strings.Replace(v, "-_", "+/", -1)
	str, err := base64.RawURLEncoding.DecodeString(v)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.Unmarshal(str, t)
}

func DecodeJwt(v string) []string {
	var splits = strings.Split(v, ".")
	var jwt []string
	for _, split := range splits {
		base64string := strings.Replace(split, "-_", "+/", -1)
		str, err := base64.RawURLEncoding.DecodeString(base64string)

		if err != nil {
			println(err)
			continue
		}

		jwt = append(jwt, string(str))
	}
	return jwt
}