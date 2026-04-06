package commands

import (
	"fmt"
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
	return nil
}
