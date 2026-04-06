package tools

import (
	"fmt"
	"os"
)


func GetConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/.config/repl-bootstrap/config.json", homeDir), nil
}
