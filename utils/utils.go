package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateToken(l int) string {
	b := make([]byte, l)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
