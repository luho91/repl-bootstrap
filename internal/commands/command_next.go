package commands

import (
	tools "github.com/luho91/repl-bootstrap/internal/tools"
	"fmt"
)

func commandNext(args []string) error {
	fmt.Printf("%s\n", tools.CleanInput("test eins zwo"))
	return nil
}
