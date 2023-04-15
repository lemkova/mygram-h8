package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPass(p string) string {
	salt := 8
	hash, _ := bcrypt.GenerateFromPassword([]byte(p), salt)
	return string(hash)
}

func ComparePass(h, p []byte) bool {
	err := bcrypt.CompareHashAndPassword(h, p)
	return err == nil
}
