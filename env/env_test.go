package env

import (
	"os"
	"testing"
)

func TestGetEnvString(t *testing.T) {
	res := getEnvString("ABC", "abc")
	if res != "abc" {
		t.Errorf("Error on setting default value %v", res)
	}

	if err := os.Setenv("ABC", "def"); err != nil {
		t.Errorf("Error on setting env %v", err)
	}

	res = getEnvString("ABC", "abc")
	if res != "def" {
		t.Errorf("Error on setting default value %v", res)
	}
}

func TestGetEnvInt(t *testing.T) {
	res := getEnvInt("ABC", 10)
	if res != 10 {
		t.Errorf("Error on setting default value %v", res)
	}

	if err := os.Setenv("ABC", "2"); err != nil {
		t.Errorf("Error on setting env %v", err)
	}

	res = getEnvInt("ABC", 10)
	if res != 2 {
		t.Errorf("Error on setting default value %v", res)
	}
}
