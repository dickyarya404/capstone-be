package helper

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(plainText string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainText), 12)
	if err != nil {
		return "", errors.New("error when hashing")
	}

	return string(hash), nil
}

func ComparePassword(hashed string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))

	return err == nil
}
