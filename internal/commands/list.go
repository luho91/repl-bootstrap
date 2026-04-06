package commands

type CliCommand struct {
	Name		string
	Description	string
	Callback	func([]string) error
}

var Commands map[string]CliCommand

func init() {
	Commands = map[string]CliCommand {
		"exit": {
			Name:			"exit",
			Description:	"Exit repl-bootstrap",
			Callback:		CommandExit,
		},
		"help": {
			Name:			"help",
			Description:	"Display possible commands",
			Callback:		CommandHelp,
		},
		"next": {
			Name:			"next",
			Description:	"Show tips for what to do next, from the state your repl is currently in",
			Callback:		CommandNext,
		},
		"create": {
			Name:			"create",
			Description:	"Create anything - help create for more details",
			Callback:		CommandCreate,
		},
	}
}
