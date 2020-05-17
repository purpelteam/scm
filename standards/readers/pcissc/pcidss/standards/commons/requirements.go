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

// Requirement contain information about PCI DSS Standard Requirement Information
type Requirement struct {
	ID              string  `yaml:"id" json:"id"`
	Description     string  `yaml:"description" json:"description"`
	Explanation     string  `yaml:"explanation" json:"explanation"`
	ExplanationNote string  `yaml:"explanation_note" json:"explanation_note"`
	Goal            string  `yaml:"goal" json:"goal"`
	Items           []*Item `yaml:"items" json:"items"`
}

// Requirements contain list of PCI DSS Requirement information
type Requirements struct {
	Values []*Requirement `yaml:"requirements" json:"requirements"`
}

// Item containt Item Information
type Item struct {
	ID               string              `yaml:"id" json:"id"`
	Description      string              `yaml:"description" json:"description"`
	DescriptionNote  string              `yaml:"description_note" json:"description_note"`
	TestingProcedure []*TestingProcedure `yaml:"testing_procedures" json:"testing_procedures"`
	Guidance         string              `yaml:"guidance" json:"guidance"`
	GuidanceNote     string              `yaml:"guidance_note" json:"guidance_note"`
	Items            []*Item             `yaml:"items" json:"items"`
}

// TestingProcedure contain Testing Procedure information
type TestingProcedure struct {
	ID              string `yaml:"id" json:"id"`
	Description     string `yaml:"description" json:"description"`
	DescriptionNote string `yaml:"description_note" json:"description_note"`
}

// ItemRoot containt ItemRoot Information
type ItemRoot struct {
	Parent string  `yaml:"parent" json:"parent"`
	Items  []*Item `yaml:"items" json:"items"`
}

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
