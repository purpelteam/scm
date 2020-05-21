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

// getRequirementsDefinitionFilePath get filepath location
func getRequirementsDefinitionFilePath(version string) (string, error) {
	filename := "requirement-parent.yaml"

	glog.V(2).Info(fmt.Sprintf("Looking Requirement for config for version %s", version))

	path := filepath.Join(CfgDir, version)
	file := filepath.Join(path, filename)

	glog.V(2).Info(fmt.Sprintf("Looking Requirement for config file: %s\n", file))

	_, err := os.Stat(file)
	if err != nil {
		return "", err
	}

	return file, nil
}

// getRequirements get Requirement content
func getRequirements(path string) (*Requirements, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	reqParents := new(Requirements)
	err = yaml.Unmarshal(data, reqParents)
	if err != nil {
		return nil, err
	}

	for i, parent := range reqParents.Values {

		reqItemFilePath, err := getReqItemsDefinitionFilePath(Version, parent.ID)
		if err != nil {
			//util.ExitWithError(err)
		} else {
			itemRoot, err := getReqItems(reqItemFilePath)
			if err != nil {
				util.ExitWithError(err)
			}

			reqParents.Values[i].Items = itemRoot.Items
		}
	}

	return reqParents, err
}

// getReqItemsDefinitionFilePath
func getReqItemsDefinitionFilePath(version, parent string) (string, error) {
	filename := "requirement-" + parent + ".yaml"

	glog.V(2).Info(fmt.Sprintf("Looking Requirement for config for version %s", version))

	path := filepath.Join(CfgDir, version)
	file := filepath.Join(path, filename)

	glog.V(2).Info(fmt.Sprintf("Looking Requirement for config file: %s\n", file))

	_, err := os.Stat(file)
	if err != nil {
		return "", err
	}

	return file, nil
}

// getReqItems
func getReqItems(path string) (*ItemRoot, error) {
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
