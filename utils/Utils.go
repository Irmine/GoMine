package utils

import (
	"fmt"
	"strings"
	"encoding/base64"
	"encoding/json"
	"crypto/ecdsa"
	"crypto/sha512"
	"crypto/rand"
	"crypto/x509"
)

type EncryptionHeader struct {
	Algorithm string `json:"alg"`
	X5u string `json:"x5u"`
}

type EncryptionPayload struct {
	Token string `json:"salt"`
}

func DecodeJwtPayload(v string, t interface{}) {
	v = strings.Split(v, ".")[1]
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
		str, err := base64.RawURLEncoding.DecodeString(split)

		if err != nil {
			println(err)
			continue
		}

		jwt = append(jwt, string(str))
	}
	return jwt
}

func ConstructEncryptionJwt(key *ecdsa.PrivateKey, token []byte) string {
	var header = EncryptionHeader{}
	header.Algorithm = "ES384"
	var b, _ = x509.MarshalPKIXPublicKey(&key.PublicKey)

	header.X5u = base64.RawStdEncoding.EncodeToString(b)

	var payload = EncryptionPayload{}
	payload.Token = base64.RawStdEncoding.EncodeToString(token)

	var headerData, _ = json.Marshal(header)
	var headerStr = base64.RawURLEncoding.EncodeToString(headerData)
	var payloadData, _ = json.Marshal(payload)
	var payloadStr = base64.RawURLEncoding.EncodeToString(payloadData)

	var hash = sha512.New384()
	hash.Write([]byte(headerStr + "." + payloadStr))

	var r, s, err = ecdsa.Sign(rand.Reader, key, hash.Sum(nil))
	if err != nil {
		fmt.Println(err)
	}

	var signature = base64.RawURLEncoding.EncodeToString(append(r.Bytes(), s.Bytes()...))

	return headerStr + "." + payloadStr + "." + signature
}