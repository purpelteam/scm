package commons

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
)

// getCISCategoriesDefinitionFilePath get filepath location
func getCISCategoriesDefinitionFilePath(version string) (string, error) {
	filename := "categories.yaml"

	glog.V(2).Info(fmt.Sprintf("Looking categories for config for version %s", version))

	path := filepath.Join(CISControlCfgDir, version)
	file := filepath.Join(path, filename)

	glog.V(2).Info(fmt.Sprintf("Looking categories for config file: %s\n", file))

	_, err := os.Stat(file)
	if err != nil {
		return "", err
	}

	return file, nil
}

// getCategories get categories content
func getCategories(path string) (*Categories, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	categories := new(Categories)
	err = yaml.Unmarshal(data, categories)
	if err != nil {
		return nil, err
	}

	return categories, err
}
