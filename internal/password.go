package internal

import (
	"crypto/sha256"
	"github.com/anaskhan96/go-password-encoder"
)

func PasswordEncode(rawPassword string) (string, string) {
	salt, encode := password.Encode(rawPassword, PasswordOption())
	return salt, encode
}

func PasswordOption() *password.Options {
	return &password.Options{
		SaltLen:      256,
		Iterations:   128,
		KeyLen:       128,
		HashFunction: sha256.New,
	}
}

func PasswordVerify(rawPassword string, salt string, encode string) bool {
	return password.Verify(rawPassword, salt, encode, PasswordOption())
}
