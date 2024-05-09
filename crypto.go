package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

const nonceSize = 12

func Encrypt(key, plainText []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, nonceSize)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return aesGCM.Seal(nil, nonce, plainText, nil), nil
}

func Decrypt(key, cipherText []byte) ([]byte, error) {
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
