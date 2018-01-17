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

	SendIV []byte
	ReceiveIV []byte
	Cipher cipher.Block

	SendCounter int64
}

func (data *EncryptionData) ComputeSharedSecret() {
	var x, _ = data.ClientPublicKey.Curve.ScalarMult(data.ClientPublicKey.X, data.ClientPublicKey.Y, data.ServerPrivateKey.D.Bytes())
	data.SharedSecret = x.Bytes()
}

func (data *EncryptionData) ComputeSecretKeyBytes() {
	data.SecretKeyBytes = sha256.Sum256(append(data.ServerToken, data.SharedSecret...))

	data.Cipher, _ = aes.NewCipher(data.SecretKeyBytes[:])
	data.SendIV = data.SecretKeyBytes[:aes.BlockSize]
	data.ReceiveIV = data.SecretKeyBytes[:aes.BlockSize]
}

type EncryptionHandler struct {
	Data *EncryptionData
}

func NewEncryptionHandler() *EncryptionHandler {
	return &EncryptionHandler{&EncryptionData{}}
}

func (handler *EncryptionHandler) ComputeSendChecksum(d []byte) []byte {
	var buffer []byte
	WriteLittleLong(&buffer, handler.Data.SendCounter)

	var sum = sha256.Sum256(append(buffer, append(d, handler.Data.SecretKeyBytes[:]...)...))
	handler.Data.SendCounter++
	return sum[:8]
}