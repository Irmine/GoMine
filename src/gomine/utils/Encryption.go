package utils

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/cipher"
	"crypto/aes"
)

type EncryptionData struct {
	ClientPublicKey *ecdsa.PublicKey
	ServerPrivateKey *ecdsa.PrivateKey
	ServerToken []byte
	SharedSecret []byte
	SecretKeyBytes [32]byte
	IV []byte
	Cipher cipher.Block
}

func (data *EncryptionData) ComputeSharedSecret() {
	var x, _ = data.ClientPublicKey.Curve.ScalarMult(data.ClientPublicKey.X, data.ClientPublicKey.Y, data.ServerPrivateKey.D.Bytes())
	data.SharedSecret = x.Bytes()
}

func (data *EncryptionData) ComputeSecretKeyBytes() {
	data.SecretKeyBytes = sha256.Sum256(append(data.ServerToken, data.SharedSecret...))

	data.Cipher, _ = aes.NewCipher(data.SecretKeyBytes[:])
	data.IV = data.SecretKeyBytes[:aes.BlockSize]
}

type EncryptionHandler struct {
	Data *EncryptionData
}

func NewEncryptionHandler() *EncryptionHandler {
	return &EncryptionHandler{&EncryptionData{}}
}
