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

// ReadMappingDefinitionByID is a function to read mapping definition by ID
func ReadMappingDefinitionByID(id string) (*MappingDefinitions, error) {
	md, err := ReadMappingDirectories()
	if err != nil {
		return nil, err
	}

	mp, err := FindMappingInfoByID(md, id)
	if err != nil {
		return nil, err
	}

	mpFilePath, err := getMappingDefinitionFilePath(mp.Location)
	if err != nil {
		return nil, err
	}

	mdef, err := getMappingDefinitions(mpFilePath)
	if err != nil {
		util.ExitWithError(err)
	}

	return mdef, err
}

// getMappingDefinitionFilePath get filepath location
func getMappingDefinitionFilePath(fileLocation string) (string, error) {

	glog.V(2).Info(fmt.Sprintf("Looking for Mapping Definition config for version %s", "current"))

	file := filepath.Join(MappingCfgDir, fileLocation)

	glog.V(2).Info(fmt.Sprintf("Looking for Mapping Definition config file: %s\n", file))

	_, err := os.Stat(file)
	if err != nil {
		return "", err
	}

	return file, nil
}

// getMappingDefinitions get mapping definition content
func getMappingDefinitions(path string) (*MappingDefinitions, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	md := new(MappingDefinitions)
	err = yaml.Unmarshal(data, md)
	if err != nil {
		return nil, err
	}

	return md, err
}
