package commands

import (
	"fmt"
)

func commandHelp() error {
	fmt.Println("The following commands are available:")
	fmt.Println()
	for _, command := range Commands {
		fmt.Println(command.Name, ":", command.Description)
	}
	return nil
}
