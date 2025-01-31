package random

import (
	"crypto/sha256"
	"encoding/hex"
)

// Sha256 returns random sha256 checksum
func Sha256() string {
	b := Words([]int{10, 20})
	hash := sha256.New()
	_, _ = hash.Write([]byte(b))
	sum := hash.Sum(nil)
	return hex.EncodeToString(sum)
}
