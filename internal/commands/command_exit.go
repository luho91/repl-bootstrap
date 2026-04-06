package commands

import (
	"fmt"
	"os"
)

func CommandExit(args []string) error {
	fmt.Println("Good bye :)")
	os.Exit(0)
	return nil
}
