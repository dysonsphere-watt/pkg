package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	b64 "encoding/base64"
	"io"
)

const nonceSize = 12

func Encrypt(key, plainText []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, nonceSize)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := aesGCM.Seal(nil, nonce, plainText, nil)

	b64CipherText := b64.URLEncoding.EncodeToString(cipherText)
	return b64CipherText, nil
}

func Decrypt(key []byte, b64CipherText string) ([]byte, error) {
	cipherText, err := b64.URLEncoding.DecodeString(b64CipherText)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]
	return aesGCM.Open(nil, nonce, cipherText, nil)
}
