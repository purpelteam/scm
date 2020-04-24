package main

import (
	"encoding/json"
	"fmt"

	pci_dss_commons "github.com/purpeltim/scm/standards/readers/pcissc/pcidss/standards/commons"
	"github.com/purpeltim/scm/standards/util"
)

func main() {
	// fmt.Println("test")

	std := pci_dss_commons.Standard{}
	pci_dss_commons.CfgDir = "standards/definitions/pcissc/pcidss/standards/"
	result, err := std.ReadStandard("3.2.1")
	if err != nil {
		util.ExitWithError(err)
	}

	pciStd, err := json.Marshal(result)
	if err != nil {
		util.ExitWithError(err)
	}
	fmt.Printf("%s", pciStd)

}
