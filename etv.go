package pkg

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

// AES encryption ETV uses before passing username/password to the login request
func EtvEncrypt(s string) (string, error) {
	keyHex := "79697368656e67746563686e6f6c6f6779636f6d70616e793230313930343131"
	ivHex := "79697368656e67746563684032303139"

	key, err := hex.DecodeString(keyHex)
	if err != nil {
		return "", err
	}

	iv, err := hex.DecodeString(ivHex)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	paddedE := pkcs7Pad([]byte(s), aes.BlockSize)
	ciphertext := make([]byte, len(paddedE))

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, paddedE)

	return hex.EncodeToString(ciphertext), nil
}

func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}
