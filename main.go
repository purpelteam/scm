package main

import (
	"encoding/json"
	"fmt"

	pci_dss_commons "github.com/purpeltim/scm/standards/readers/pcissc/pcidss/standards/commons"
	"github.com/purpeltim/scm/standards/util"
	"github.com/purpeltim/scm/tools/sysinfo"
)

func main() {
	// fmt.Println("test")

	type TestProject struct {
		HostInfo  *sysinfo.HostInfo
		PciDssStd *pci_dss_commons.Standard
	}

	T := TestProject{}

	std := pci_dss_commons.Standard{}
	pci_dss_commons.CfgDir = "standards/definitions/pcissc/pcidss/standards/"

	result, err := std.ReadStandard("3.2.1")
	if err != nil {
		util.ExitWithError(err)
	}

	T.PciDssStd = result
	T.HostInfo, err = sysinfo.GetHostInfo()
	if err != nil {
		util.ExitWithError(err)
	}

	TestPciStd, err := json.Marshal(T)
	if err != nil {
		util.ExitWithError(err)
	}
	fmt.Printf("%s", TestPciStd)

}
