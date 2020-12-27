package utils

import "crypto/sha512"

func HashPassword(password string) {
	return sha512.New()
}