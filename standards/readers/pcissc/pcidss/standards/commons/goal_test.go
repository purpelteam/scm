package commons

import (
	"os"
	"testing"
)

func TestGetGoalDefinitionFilePath(t *testing.T) {
	CfgDir = cfgDir
	Version = version

	filename, err := getGoalDefinitionFilePath(Version)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		t.Fail()
	}

	if info.Name() != "goals.yaml" {
		t.Fail()
	}
}

func TestGetGoal(t *testing.T) {
	CfgDir = cfgDir
	Version = version

	filename, err := getGoalDefinitionFilePath(Version)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	goals, err := getGoals(filename)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if goals.Values[0].ID != "1" {
		t.Fail()
	}
}
