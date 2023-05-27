package bhash

import (
	"crypto/sha256"
	"encoding/hex"
)

type Hasher interface {
	DataToHash() string
}

func GenerateHash(h Hasher) string {
	data := h.DataToHash()
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
