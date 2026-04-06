package tools

import (
	"os"
)

type Facts struct {
	ConfigExists	bool
}

func CollectFacts() Facts {
	f := Facts{}

	f.ConfigExists = checkConfigExists()

	return f
}

func checkConfigExists() bool {
	configPath, err := GetConfigPath()
	if err != nil {
		return false
	}
	_, err = os.Stat(configPath)

	return err == nil
}
