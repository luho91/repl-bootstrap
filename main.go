package main

import(
	tools "github.com/luho91/repl-bootstrap/internal/tools"
	commands "github.com/luho91/repl-bootstrap/internal/commands"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("REB > Welcome to repl-bootstrap!")
	fmt.Println("REB > Type \"help\" to see an overview of possible commands.")

	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("REB > ")
		scanner.Scan()
		input := tools.CleanInput(scanner.Text())

		if len(input) == 0 {
			continue
		}

		cmdString := input[0]
		command, ok := commands.Commands[cmdString]

		if ok {
			command.Callback()
		} else {
			fmt.Printf("Unknown command: %s\n", cmdString)
		}
	}
}
