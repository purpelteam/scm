package commons

import (
	"os"
	"testing"
)

var (
	cfgDir  = "../../../../../../standards/definitions/pcissc/pcidss/standards/"
	version = "3.2.1"
)

func TestReadStandard(t *testing.T) {
	CfgDir = cfgDir
	Version = version

	s, err := ReadStandard(version)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if s.ID != Version {
		t.Fail()
	}
}

func TestGetStandardDefinitionFilePath(t *testing.T) {
	CfgDir = cfgDir
	Version = version

	filename, err := getStandardDefinitionFilePath(Version)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		t.Fail()
	}

	if info.Name() != "standard.yaml" {
		t.Fail()
	}
}

func TestGetStandard(t *testing.T) {
	CfgDir = cfgDir
	Version = version

	filename, err := getStandardDefinitionFilePath(Version)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	standard, err := getStandard(filename)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if standard.ID != Version {
		t.Fail()
	}
}
