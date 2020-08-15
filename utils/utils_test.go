package utils

import (
	"os"
	"testing"
)

func TestGetEnvPanic(t *testing.T) {
	k, v := "KEY1", "VALUE1"
	os.Setenv(k, v)
	if GetEnvFatal(k) != "VALUE1" {
		t.Error()
	}
}
