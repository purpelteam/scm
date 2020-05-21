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

// ReadStandard is function to read Standard Definition
func ReadStandard(version string) (*Standard, error) {
	if version == "" {
		version = Version
	}

	if version == "" {
		version = "3.2.1"
	}
	if Version == "" {
		Version = "3.2.1"
	}

	// standard
	stdFilePath, err := getStandardDefinitionFilePath(version)
	if err != nil {
		util.ExitWithError(err)
	}

	// s := Standard{}

	s, err := getStandard(stdFilePath)
	if err != nil {
		util.ExitWithError(err)
	}

	// goals
	goalFilePath, err := getGoalDefinitionFilePath(version)
	if err != nil {
		util.ExitWithError(err)
	}

	goals, err := getGoals(goalFilePath)
	if err != nil {
		util.ExitWithError(err)
	}
	s.Goals = goals.Values

	// Requirements
	reqParentFilePath, err := getRequirementsDefinitionFilePath(version)
	if err != nil {
		util.ExitWithError(err)
	}

	reqParents, err := getRequirements(reqParentFilePath)
	if err != nil {
		util.ExitWithError(err)
	}
	s.Requirements = reqParents.Values

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
