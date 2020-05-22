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

// ReadMappingDirectories is function to read Mapping Directories
func ReadMappingDirectories() (*MappingDirectories, error) {

	// Mapping Directories
	mdFilePath, err := getMappingDirectoriesFilePath()
	if err != nil {
		util.ExitWithError(err)
	}

	mp, err := getMappingDirectories(mdFilePath)
	if err != nil {
		util.ExitWithError(err)
	}

	return mp, err
}

// FindMappingInfoByID is function to find Mapping Information by ID
func FindMappingInfoByID(md *MappingDirectories, ID string) (*Mapping, error) {
	var result *Mapping
	var err error

	for _, mapping := range md.Mappings {
		if mapping.ID == ID {
			result = mapping
		}
	}

	return result, err
}

// getMappingDirectoriesFilePath get filepath location
func getMappingDirectoriesFilePath() (string, error) {
	filename := "index.yaml"

	glog.V(2).Info(fmt.Sprintf("Looking for Mapping Directories config for version %s", "current"))

	file := filepath.Join(MappingCfgDir, filename)

	glog.V(2).Info(fmt.Sprintf("Looking for Mapping Directories config file: %s\n", file))

	_, err := os.Stat(file)
	if err != nil {
		return "", err
	}

	return file, nil
}

// getMappingDirectories get mapping directories content
func getMappingDirectories(path string) (*MappingDirectories, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	md := new(MappingDirectories)
	err = yaml.Unmarshal(data, md)
	if err != nil {
		return nil, err
	}

	return md, err
}
