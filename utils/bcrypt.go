package utils

import "golang.org/x/crypto/bcrypt"

func Encrypt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassword(current, encrypted string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(current))
	return err == nil
}
