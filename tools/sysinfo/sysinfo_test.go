package sysinfo

import (
	"testing"
)

func TestGetHostInfo(t *testing.T) {
	hi, err := GetHostInfo()

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if hi.OS.Name == "" {
		t.Fail()
	}
}
