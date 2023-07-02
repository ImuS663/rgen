package rgen

import (
	"math/rand"
	"strings"
	"time"
)

const (
	latinLowercaseCharset = "abcdefghijklmnopqrstuvwxyz"
	latinUppercaseCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberCharset         = "0123456789"
	specialCharset        = "!@#$%^&*/+-="
)

var random *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func Bytes(length uint) []byte {
	b := make([]byte, length)

	random.Read(b)

	return b
}

func String(length uint) string {
	charset := latinLowercaseCharset + latinUppercaseCharset + numberCharset

	return stringWithCharset(length, charset)
}

func StringLatinOnly(length uint) string {
	charset := latinLowercaseCharset + latinUppercaseCharset

	return stringWithCharset(length, charset)
}

func StringSpecial(length uint) string {
	charset := latinLowercaseCharset + latinUppercaseCharset + numberCharset + specialCharset

	return stringWithCharset(length, charset)
}

func StringByPreset(preset string) string {
	for true {
		preset = strings.Replace(preset, "%a", stringWithCharset(1, latinLowercaseCharset+latinUppercaseCharset+numberCharset), 1)
		preset = strings.Replace(preset, "%o", stringWithCharset(1, latinLowercaseCharset+latinUppercaseCharset), 1)
		preset = strings.Replace(preset, "%l", stringWithCharset(1, latinLowercaseCharset), 1)
		preset = strings.Replace(preset, "%u", stringWithCharset(1, latinUppercaseCharset), 1)
		preset = strings.Replace(preset, "%n", stringWithCharset(1, numberCharset), 1)

		if !cont(preset) {
			break
		}
	}

	return preset
}

func stringWithCharset(length uint, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[random.Intn(len(charset))]
	}

	return string(b)
}

func cont(preset string) bool {
	return strings.Contains(preset, "%a") || strings.Contains(preset, "%c") ||
		strings.Contains(preset, "%l") || strings.Contains(preset, "%u") ||
		strings.Contains(preset, "%n")
}
