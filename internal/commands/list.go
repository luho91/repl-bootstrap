package commands

type cliCommand struct {
	Name		string
	Description	string
	Callback	func([]string) error
}

var Commands map[string]cliCommand

func init() {
	Commands = map[string]cliCommand {
		"exit": {
			Name:			"exit",
			Description:	"Exit repl-bootstrap",
			Callback:		commandExit,
		},
		"help": {
			Name:			"help",
			Description:	"Display possible commands",
			Callback:		commandHelp,
		},
		"next": {
			Name:			"next",
			Description:	"Show tips for what to do next, from the state your repl is currently in",
			Callback:		commandNext,
		},
	}
}
