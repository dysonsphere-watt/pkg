package pkg

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
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

func EtvDecrypt(cipherTextHex string) (string, error) {
	keyHex := "79697368656e67746563686e6f6c6f6779636f6d70616e793230313930343131"
	ivHex := "79697368656e67746563684032303139"

	key, err := hex.DecodeString(keyHex)
	if err != nil {
		return "", fmt.Errorf("failed to decode key: %w", err)
	}

	iv, err := hex.DecodeString(ivHex)
	if err != nil {
		return "", fmt.Errorf("failed to decode IV: %w", err)
	}

	ciphertext, err := hex.DecodeString(cipherTextHex)
	if err != nil {
		return "", fmt.Errorf("failed to decode ciphertext: %w", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %w", err)
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	plaintext, err := pkcs7Unpad(ciphertext, aes.BlockSize)
	if err != nil {
		return "", fmt.Errorf("failed to unpad plaintext: %w", err)
	}

	return string(plaintext), nil
}

func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("data is empty")
	}
	if length%blockSize != 0 {
		return nil, errors.New("data is not a multiple of the block size")
	}
	paddingLen := int(data[length-1])
	if paddingLen > blockSize || paddingLen == 0 {
		return nil, errors.New("invalid padding size")
	}
	for _, v := range data[length-paddingLen:] {
		if int(v) != paddingLen {
			return nil, errors.New("invalid padding")
		}
	}
	return data[:length-paddingLen], nil
}
