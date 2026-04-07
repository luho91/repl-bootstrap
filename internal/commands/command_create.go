package commands

import (
	tools "github.com/luho91/repl-bootstrap/internal/tools"
	"fmt"
	"os"
)

type createArg string

const (
	Config createArg = "config"
)

func CommandCreate(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("No argument provided, expected 1 argument")
	}

	switch createArg(args[0]) {
	case Config: return createConfig()
	default: return fmt.Errorf("Unknown argument %s", args[0])
	}
}

func createConfig() error {
	configFolder, err := tools.GetConfigFolderPath()

	if err != nil {
		return err
	}

	err = os.MkdirAll(configFolder, 0o775)

	if err != nil {
		return err
	}

	configPath := tools.GetConfigFilePath(configFolder)

	f, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}
