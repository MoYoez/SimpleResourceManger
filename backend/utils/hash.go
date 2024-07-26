package tools

import (
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

// Hasher Hash Password.
func Hasher(LoginInPassword string) string {
	hasherMixed, err := bcrypt.GenerateFromPassword([]byte(LoginInPassword), 10)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(hasherMixed)
}
