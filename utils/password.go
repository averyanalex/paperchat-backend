package utils

import (
	"crypto/sha512"
	"fmt"
)

// HashPassword will return SHA512 string hash
func HashPassword(password string) string {
	h := sha512.New()
	h.Write([]byte(password))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}


// CheckPassword will return true, if given password have same hash
func CheckPassword(password string, hash string) bool {
	return HashPassword(password) == hash
}