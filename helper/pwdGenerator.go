package helper

import (
	"math/rand"
	"time"
)

func GeneratePassword(length int, includeUppercase, includeLowercase, includeNumbers, includeSpecialChars bool) string {
	charset := ""
	if includeUppercase {
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if includeLowercase {
		charset += "abcdefghijklmnopqrstuvwxyz"
	}
	if includeNumbers {
		charset += "0123456789"
	}
	if includeSpecialChars {
		charset += "!@#$%^&*()-_=+[]{}|;:'\",.<>?/\\"
	}

	password := make([]byte, length)
	rand.Seed(time.Now().UnixNano())
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	return string(password)
}
