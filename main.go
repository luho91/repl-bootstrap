package main

import(
	tools "github.com/luho91/repl-bootstrap/internal/tools"
	commands "github.com/luho91/repl-bootstrap/internal/commands"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to repl-bootstrap!")
	fmt.Println("Type \"help\" to see an overview of possible commands.")
	fmt.Println("Type \"next\" to see tips for what to do next (or how to begin).")

	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("REB > ")
		scanner.Scan()
		input := tools.CleanInput(scanner.Text())

		if len(input) == 0 {
			continue
		}

		cmdString := input[0]
		args := input[1:]
		command, ok := commands.Commands[cmdString]

		if ok {
			err := command.Callback(args)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Printf("Unknown command: %s\n", cmdString)
		}
	}
}
