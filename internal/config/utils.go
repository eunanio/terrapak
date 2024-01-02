package config

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashSecret(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}