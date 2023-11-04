package utils

import (
	"bytes"
	"testing"
)

func TestHash256(t *testing.T) {
	if !bytes.Equal(Hash256("1000"), Hash256("1000")){
		t.Error("Not equal")
	}
}
