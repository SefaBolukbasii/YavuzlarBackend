package pkg

import (
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

var fixedSalt = []byte("SuperSecretSalt123!")

func HashPassword(password string) (string, error) {
	hash := argon2.IDKey([]byte(password), fixedSalt, 1, 64*1024, 4, 32)
	hashedPassword := base64.StdEncoding.EncodeToString(append(fixedSalt, hash...))
	return hashedPassword, nil
}
