package utils

import (
	"crypto/sha256"
	"fmt"
)

func Hash(data []byte) string {
	hash := sha256.New()
	hash.Write(data)

	hashed := hash.Sum(nil)
	return fmt.Sprintf("%x", hashed)
}
