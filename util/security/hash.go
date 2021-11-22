package security

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func GenerateHash(info ...string) (hash string) {
	newDna := strings.Join(info, ",")
	bytes := sha256.Sum256([]byte(newDna))
	hash = fmt.Sprintf("%X", bytes)
	return
}
