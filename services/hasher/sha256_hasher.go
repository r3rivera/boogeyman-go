package hasher

import (
	"crypto/sha256"
	"encoding/hex"
)

type Sha256HashItem string

func (v *Sha256HashItem) HashItem() (string, error) {
	hash := sha256.Sum256([]byte(*v))
	return hex.EncodeToString(hash[:]), nil
}
