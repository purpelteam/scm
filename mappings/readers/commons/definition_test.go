package commons

import (
	"testing"
)

func TestReadMappingDefinitionByID(t *testing.T) {
	MappingCfgDir = cfgDir

	md, err := ReadMappingDefinitionByID("1")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if md.From != "pci-dss-3.2.1" {
		t.Fail()
	}
}
