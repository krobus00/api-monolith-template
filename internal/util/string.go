package util

import (
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/scrypt"
)

func HashPassword(password string, salt []byte) (string, error) {
	hash, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, 32)
	if err != nil {
		return "", err
	}
	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)
	encodedHash := base64.RawStdEncoding.EncodeToString(hash)
	return fmt.Sprintf("%s.%s", encodedSalt, encodedHash), nil
}

func ComparePassword(hashedPassword, password string) (bool, error) {
	parts := strings.Split(hashedPassword, ".")
	if len(parts) != 2 {
		return false, fmt.Errorf("invalid hashed password format")
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return false, err
	}

	hash, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return false, err
	}

	newHash, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, 32)
	if err != nil {
		return false, err
	}

	return subtle.ConstantTimeCompare(hash, newHash) == 1, nil
}
