package rgen

import (
	"bytes"
	"github.com/ImuS663/rgen"
	"testing"
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
