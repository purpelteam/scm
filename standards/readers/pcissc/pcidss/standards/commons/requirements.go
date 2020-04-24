package commons

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
)

// Requirement contain information about PCI DSS Standard Requirement Information
type Requirement struct {
	ID          string `yaml:"id" json:"id"`
	Description string `yaml:"description" json:"description"`
}

// Requirements contain list of PCI DSS Requirement information
type Requirements struct {
	Values []*Requirement `yaml:"requirements" json:"requirements"`
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

	return reqParents, err
}
