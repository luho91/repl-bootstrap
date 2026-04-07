package tools

import (
	"fmt"
	"os"
)

func GetConfigPath() (string, error) {
	folderPath, err := GetConfigFolderPath()
	if err != nil {
		return "", err
	}

	return GetConfigFilePath(folderPath), nil
}

func GetConfigFilePath(configDir string) string {

	return fmt.Sprintf("%s/connfig.json", configDir)
}

func GetConfigFolderPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/.config/repl-bootstrap", homeDir), nil
}
