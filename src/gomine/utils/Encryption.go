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
	DecryptSecretKeyBytes [32]byte
	EncryptSecretKeyBytes [32]byte

	DecryptIV []byte
	EncryptIV []byte
	DecryptCipher cipher.Block
	EncryptCipher cipher.Block

	SendCounter int64
}

func (data *EncryptionData) ComputeSharedSecret() {
	var x, _ = data.ClientPublicKey.Curve.ScalarMult(data.ClientPublicKey.X, data.ClientPublicKey.Y, data.ServerPrivateKey.D.Bytes())
	data.SharedSecret = x.Bytes()
}

func (data *EncryptionData) ComputeSecretKeyBytes() {
	var secret = sha256.Sum256(append(data.ServerToken, data.SharedSecret...))
	data.DecryptSecretKeyBytes = secret
	data.EncryptSecretKeyBytes = secret

	data.DecryptCipher, _ = aes.NewCipher(data.DecryptSecretKeyBytes[:])
	data.EncryptCipher, _ = aes.NewCipher(data.EncryptSecretKeyBytes[:])

	data.DecryptIV = data.DecryptSecretKeyBytes[:aes.BlockSize]
	data.EncryptIV = data.DecryptSecretKeyBytes[:aes.BlockSize]
}

type EncryptionHandler struct {
	Data *EncryptionData
}

func NewEncryptionHandler() *EncryptionHandler {
	return &EncryptionHandler{&EncryptionData{}}
}

func (handler *EncryptionHandler) ComputeSendChecksum(d []byte) []byte {
	var buffer []byte
	var secret = handler.Data.EncryptSecretKeyBytes[:]

	WriteLittleLong(&buffer, handler.Data.SendCounter)
	handler.Data.SendCounter++

	var hash = sha256.New()
	hash.Write(buffer)
	hash.Write(d)
	hash.Write(secret)

	var sum = hash.Sum(nil)
	return sum[:8]
}