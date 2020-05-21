package commons

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
)

// getStandardDefinitionFilePath get filepath location
func getGoalDefinitionFilePath(version string) (string, error) {
	filename := "goals.yaml"

	glog.V(2).Info(fmt.Sprintf("Looking Goal for config for version %s", version))

	path := filepath.Join(CfgDir, version)
	file := filepath.Join(path, filename)

	glog.V(2).Info(fmt.Sprintf("Looking Goal for config file: %s\n", file))

	_, err := os.Stat(file)
	if err != nil {
		return "", err
	}

	return file, nil
}

// getGoals get goal content
func getGoals(path string) (*Goals, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	goals := new(Goals)
	err = yaml.Unmarshal(data, goals)
	if err != nil {
		return nil, err
	}

	return goals, err
}
