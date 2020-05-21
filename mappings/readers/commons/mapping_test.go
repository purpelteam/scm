package commons

import (
	"os"
	"testing"
)

var (
	cfgDir = "../../../mappings/definitions/"
)

func TestReadMappingDirectories(t *testing.T) {
	MappingCfgDir = cfgDir

	mp, err := ReadMappingDirectories()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if mp.Indexs[0].ID != "pci-dss-3.2.1" {
		t.Fail()
	}
}

func TestGetMappingDirectoriesFilePath(t *testing.T) {
	MappingCfgDir = cfgDir

	filename, err := getMappingDirectoriesFilePath()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		t.Fail()
	}

	if info.Name() != "index.yaml" {
		t.Fail()
	}
}

func TestGetMappingDirectories(t *testing.T) {
	MappingCfgDir = cfgDir

	filename, err := getMappingDirectoriesFilePath()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	mp, err := getMappingDirectories(filename)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if mp.Indexs[0].ID != "pci-dss-3.2.1" {
		t.Fail()
	}
}
