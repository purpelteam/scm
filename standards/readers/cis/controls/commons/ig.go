package commons

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
)

// getCISIGsDefinitionFilePath get filepath location
func getCISIGsDefinitionFilePath(version string) (string, error) {
	filename := "igs.yaml"

	glog.V(2).Info(fmt.Sprintf("Looking IGs for config for version %s", version))

	path := filepath.Join(CISControlCfgDir, version)
	file := filepath.Join(path, filename)

	glog.V(2).Info(fmt.Sprintf("Looking IGs for config file: %s\n", file))

	_, err := os.Stat(file)
	if err != nil {
		return "", err
	}

	return file, nil
}

// getImplemetationGroups get Imlementation Groups content
func getImplemetationGroups(path string) (*ImplementaionGroups, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	igs := new(ImplementaionGroups)
	err = yaml.Unmarshal(data, igs)
	if err != nil {
		return nil, err
	}

	return igs, err
}
