package tools

import (
	"os"
)

type Facts struct {
	ConfigExists		bool
	ConfigIsEmpty		bool
	ConfigIsValidJSON	bool
}

func CollectFacts() Facts {
	f := Facts{}

	f.ConfigExists = checkConfigExists()
	f.ConfigIsEmpty = checkConfigIsEmpty()
	f.ConfigIsValidJSON = checkConfigIsValidJSON()

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

func checkConfigIsEmpty() bool {
	configPath, err := GetConfigPath()
	if err != nil {
		return true
	}

	content, err := ReadFile(configPath)

	if content == "" || err != nil {
		return true
	}

	return false
}

func checkConfigIsValidJSON() bool {
	return false
}
