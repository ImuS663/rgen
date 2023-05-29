package rgen

import (
	"bytes"
	"github.com/ImuS663/rgen"
	"strings"
	"testing"
)

const (
	latinLowercaseCharset = "abcdefghijklmnopqrstuvwxyz"
	latinUppercaseCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberCharset         = "0123456789"
	specialCharset        = "!@#$%^&*/+-="
)

func TestBytesLength(t *testing.T) {
	b := rgen.Bytes(15)

	if len(b) != 15 {
		t.Fatalf("incorrect byte slyce length: %v", len(b))
	}
}

func TestBytesEmpty(t *testing.T) {
	b := rgen.Bytes(15)
	be := make([]byte, 15)

	if bytes.Equal(be, b) {
		t.Fatalf("byte slyce empty")
	}
}

func TestStringLength(t *testing.T) {
	s := rgen.String(17)

	if len(s) != 17 {
		t.Fatalf("incorrect string length: %v", len(s))
	}
}

func TestStringLatinOnlyLength(t *testing.T) {
	s := rgen.StringLatinOnly(28)

	if len(s) != 28 {
		t.Fatalf("incorrect string length: %v", len(s))
	}
}

func TestStringSpecialLength(t *testing.T) {
	s := rgen.StringSpecial(8)

	if len(s) != 8 {
		t.Fatalf("incorrect string length: %v", len(s))
	}
}

func TestStringCharset(t *testing.T) {
	charset := latinLowercaseCharset + latinUppercaseCharset + numberCharset
	s := rgen.String(15)

	for i := range s {
		if !strings.Contains(charset, string(s[i])) {
			t.Fatalf("string contains a incorrect char: %v", string(s[i]))
		}
	}
}

func TestStringLatinOnlyCharset(t *testing.T) {
	charset := latinLowercaseCharset + latinUppercaseCharset
	s := rgen.StringLatinOnly(28)

	for i := range s {
		if !strings.Contains(charset, string(s[i])) {
			t.Fatalf("string contains a incorrect char: %v", string(s[i]))
		}
	}
}

func TestStringSpecialCharset(t *testing.T) {
	charset := latinLowercaseCharset + latinUppercaseCharset + numberCharset + specialCharset
	s := rgen.StringSpecial(8)

	for i := range s {
		if !strings.Contains(charset, string(s[i])) {
			t.Fatalf("string contains a incorrect char: %v", string(s[i]))
		}
	}
}
