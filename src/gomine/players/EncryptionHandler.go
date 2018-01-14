package players

import "crypto/ecdsa"

type EncryptionData struct {
	ClientPublicKey *ecdsa.PublicKey
	ServerPrivateKey *ecdsa.PrivateKey
	SharedSecret []byte
	SecretBytes [32]byte
}

type EncryptionHandler struct {
	data EncryptionData
}

func NewEncryptionHandler() *EncryptionHandler {
	return &EncryptionHandler{EncryptionData{}}
}
