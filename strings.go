package pkg

import (
	"crypto/rand"
	"math/big"
)

func GenerateRandomAlphaNum(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

	b := make([]byte, length)
	for i := range b {
		charsetIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		for err != nil {
			charsetIndex, err = rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		}
		b[i] = charset[int(charsetIndex.Int64())]
	}

	return string(b)
}
