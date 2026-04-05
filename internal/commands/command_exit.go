package commands

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Good bye :)")
	os.Exit(0)
	return nil
}
