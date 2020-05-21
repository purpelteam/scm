package main

import (
	"encoding/json"
	"fmt"

	mapping_commons "github.com/purpeltim/scm/mappings/readers/commons"
	cis_control_commons "github.com/purpeltim/scm/standards/readers/cis/controls/commons"
	pci_dss_commons "github.com/purpeltim/scm/standards/readers/pcissc/pcidss/standards/commons"

	"github.com/purpeltim/scm/standards/util"
	"github.com/purpeltim/scm/tools/sysinfo"
)

func main() {
	// fmt.Println("test")

	// TestSCM Type
	type TestSCM struct {
		HostInfo   *sysinfo.HostInfo
		PciDssStd  *pci_dss_commons.Standard
		CisControl *cis_control_commons.CISControl
		Mappings   *mapping_commons.MappingDirectories
	}

	// Variable
	var err error
	T := TestSCM{}

	// Host Info
	T.HostInfo, err = sysinfo.GetHostInfo()
	if err != nil {
		util.ExitWithError(err)
	}

	// PCI DSS
	pci_dss_commons.CfgDir = "standards/definitions/pcissc/pcidss/standards/"

	stdPCIDss, err := pci_dss_commons.ReadStandard("3.2.1")
	if err != nil {
		util.ExitWithError(err)
	}
	T.PciDssStd = stdPCIDss

	// CIS Control
	cis_control_commons.CISControlCfgDir = "standards/definitions/cis/controls/"

	cisControls, err := cis_control_commons.ReadCISControl("7.1")
	if err != nil {
		util.ExitWithError(err)
	}
	T.CisControl = cisControls

	// Mapping Directories
	mapping_commons.MappingCfgDir = "mappings/definitions/"

	md, err := mapping_commons.ReadMappingDirectories()
	if err != nil {
		util.ExitWithError(err)
	}
	T.Mappings = md

	// Verbose TestSCM
	testSCM, err := json.Marshal(T)
	if err != nil {
		util.ExitWithError(err)
	}
	fmt.Printf("%s", testSCM)

}
