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

// ReadCISControl is function to read CIS Control Definition
func ReadCISControl(version string) (*CISControl, error) {
	if version == "" {
		version = CISControlVersion
	}

	if version == "" {
		version = "7.1"
	}
	if CISControlVersion == "" {
		CISControlVersion = "7.1"
	}

	// Control
	stdFilePath, err := getCISControlDefinitionFilePath(version)
	if err != nil {
		util.ExitWithError(err)
	}

	// s := CISControl{}
	s, err := getCISControl(stdFilePath)
	if err != nil {
		util.ExitWithError(err)
	}

	// Categories
	catFilePath, err := getCISCategoriesDefinitionFilePath(version)
	if err != nil {
		util.ExitWithError(err)
	}

	categories, err := getCategories(catFilePath)
	if err != nil {
		util.ExitWithError(err)
	}
	s.Categories = categories.Values

	// Implementation Groups
	igsFilePath, err := getCISIGsDefinitionFilePath(version)
	if err != nil {
		util.ExitWithError(err)
	}

	igs, err := getImplemetationGroups(igsFilePath)
	if err != nil {
		util.ExitWithError(err)
	}
	s.IGs = igs.Values

	// Controls
	controlsFilePath, err := getControlsDefinitionFilePath(version)
	if err != nil {
		util.ExitWithError(err)
	}

	controls, err := getControls(controlsFilePath)
	if err != nil {
		util.ExitWithError(err)
	}
	s.Controls = controls.Values

	return s, nil
}

// getCISControlDefinitionFilePath get filepath location
func getCISControlDefinitionFilePath(version string) (string, error) {
	filename := "control.yaml"

	glog.V(2).Info(fmt.Sprintf("Looking for CIS Control config for version %s", version))

	path := filepath.Join(CISControlCfgDir, version)
	file := filepath.Join(path, filename)

	glog.V(2).Info(fmt.Sprintf("Looking for CIS Control config file: %s\n", file))

	_, err := os.Stat(file)
	if err != nil {
		return "", err
	}

	return file, nil
}

// getCISControl get standard content
func getCISControl(path string) (*CISControl, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	standard := new(CISControl)
	err = yaml.Unmarshal(data, standard)
	if err != nil {
		return nil, err
	}

	return standard, err
}
