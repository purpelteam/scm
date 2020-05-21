package commons

import (
	"os"
	"testing"
)

var (
	cfgDir  = "../../../../../standards/definitions/cis/controls/"
	version = "7.1"
)

func TestReadCISControl(t *testing.T) {
	CISControlCfgDir = cfgDir
	CISControlVersion = version

	s, err := ReadCISControl(version)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if s.ID != CISControlVersion {
		t.Fail()
	}
}

func TestGetCISControlDefinitionFilePath(t *testing.T) {
	CISControlCfgDir = cfgDir
	CISControlVersion = version

	filename, err := getCISControlDefinitionFilePath(CISControlVersion)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		t.Fail()
	}

	if info.Name() != "control.yaml" {
		t.Fail()
	}
}

func TestGetCISControl(t *testing.T) {
	CISControlCfgDir = cfgDir
	CISControlVersion = version

	filename, err := getCISControlDefinitionFilePath(CISControlVersion)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	standard, err := getCISControl(filename)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if standard.ID != CISControlVersion {
		t.Fail()
	}
}
