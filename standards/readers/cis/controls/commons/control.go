package commons

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/golang/glog"
	"github.com/purpeltim/scm/standards/util"
	"gopkg.in/yaml.v2"
)

// getControlsDefinitionFilePath get filepath location
func getControlsDefinitionFilePath(version string) (string, error) {
	filename := "control-parent.yaml"

	glog.V(2).Info(fmt.Sprintf("Looking Controls for config for version %s", version))

	path := filepath.Join(CISControlCfgDir, version)
	file := filepath.Join(path, filename)

	glog.V(2).Info(fmt.Sprintf("Looking Controls for config file: %s\n", file))

	_, err := os.Stat(file)
	if err != nil {
		return "", err
	}

	return file, nil
}

// getControls get CIS Control content
func getControls(path string) (*Controls, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	reqParents := new(Controls)
	err = yaml.Unmarshal(data, reqParents)
	if err != nil {
		return nil, err
	}

	for i, parent := range reqParents.Values {

		subControlFilePath, err := getSubControlDefinitionFilePath(CISControlVersion, parent.ID)
		if err != nil {
			//util.ExitWithError(err)
		} else {
			itemRoot, err := getSubControls(subControlFilePath)
			if err != nil {
				util.ExitWithError(err)
			}

			reqParents.Values[i].SubControls = itemRoot.Items
		}
	}

	return reqParents, err
}

// getSubControlDefinitionFilePath
func getSubControlDefinitionFilePath(version, parent string) (string, error) {
	filename := "control-" + parent + ".yaml"

	glog.V(2).Info(fmt.Sprintf("Looking Sub Control for config for version %s", version))

	path := filepath.Join(CISControlCfgDir, version)
	file := filepath.Join(path, filename)

	glog.V(2).Info(fmt.Sprintf("Looking Sub Control for config file: %s\n", file))

	_, err := os.Stat(file)
	if err != nil {
		return "", err
	}

	return file, nil
}

// getSubControls
func getSubControls(path string) (*ItemRoot, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	reqItems := new(ItemRoot)
	err = yaml.Unmarshal(data, reqItems)
	if err != nil {
		return nil, err
	}

	return reqItems, err
}
