package commons

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/golang/glog"
	"github.com/purpeltim/scm/util"

	"gopkg.in/yaml.v2"
)

var (
	// Version of PCI DSS Standard
	Version string
	// CfgDir is Config Directory
	CfgDir string
)

// Standard contain information about PCI DSS Standards
type Standard struct {
	ID          string `yaml:"id" json:"id"`
	Description string `yaml:"description" json:"description"`
	Since       string `yaml:"since" json:"since"`
	Author      string `yaml:"author" json:"author"`
}

// ReadStandard is function to read Standard Definition
func (s *Standard) ReadStandard(version string) (*Standard, error) {
	if version == "" {
		version = Version
	}

	if version == "" {
		version = "3.2.1"
	}

	path, err := getStandardDefinitionFilePath(version)
	if err != nil {
		util.ExitWithError(err)
	}

	s, err = getStandard(path)
	if err != nil {
		util.ExitWithError(err)
	}

	return s, nil
}

// getStandardDefinitionFilePath get filepath location
func getStandardDefinitionFilePath(version string) (string, error) {
	filename := "standard.yaml"

	glog.V(2).Info(fmt.Sprintf("Looking for config for version %s", version))

	path := filepath.Join(CfgDir, version)
	file := filepath.Join(path, filename)

	glog.V(2).Info(fmt.Sprintf("Looking for config file: %s\n", file))

	_, err := os.Stat(file)
	if err != nil {
		return "", err
	}

	return file, nil
}

// getStandard get standard content
func getStandard(path string) (*Standard, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	standard := new(Standard)
	err = yaml.Unmarshal(data, standard)
	if err != nil {
		return nil, err
	}

	return standard, err
}
