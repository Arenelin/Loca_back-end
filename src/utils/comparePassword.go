package utils

import (
	"bytes"
	"encoding/base64"
	"golang.org/x/crypto/argon2"
	"strings"
)

func ComparePassword(password, hashedPassword string) bool {
	parts := strings.Split(hashedPassword, "$")
	if len(parts) != 2 {
		return false
	}

	salt, err := base64.StdEncoding.DecodeString(parts[0])
	if err != nil {
		return false
	}

	expectedHash, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return false
	}

	newHash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	return bytes.Equal(newHash, expectedHash)
}
