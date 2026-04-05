package commands

type cliCommand struct {
	Name		string
	Description	string
	Callback	func() error
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
	}
}
