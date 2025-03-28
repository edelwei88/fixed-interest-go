package lib

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashString(s string) string {
	hash := sha256.Sum256([]byte(s))
	hexHash := hex.EncodeToString(hash[:])

	return hexHash
}
