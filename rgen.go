package rgen

import (
	"math/rand"
	"time"
)

const (
	latinLowercaseCharset = "abcdefghijklmnopqrstuvwxyz"
	latinUppercaseCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberCharset         = "0123456789"
	specialCharset        = "!@#$%^&*/+-="
)

var random *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func stringWithCharset(length uint, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[random.Intn(len(charset))]
	}

	return string(b)
}
