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

func String(length uint) string {
	charset := latinLowercaseCharset + latinUppercaseCharset + numberCharset

	return stringWithCharset(length, charset)
}

func StringLatinOnly(length uint) string {
	charset := latinLowercaseCharset + latinUppercaseCharset

	return stringWithCharset(length, charset)
}

func stringWithCharset(length uint, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[random.Intn(len(charset))]
	}

	return string(b)
}
