package commands

import (
	tools "github.com/luho91/repl-bootstrap/internal/tools"
	"fmt"
)

func CommandNext(args []string) error {
	facts := tools.CollectFacts()
	issues := tools.EvaluateIssues(facts)
	suggestions := tools.GetSuggestion(issues)

	if len(suggestions) == 0 {
		fmt.Println("Nothing to do!")
	}

	for i, s := range suggestions {
		fmt.Println(fmt.Sprintf("======[ Suggestion #%v ]======", i + 1))
		fmt.Println("Title:", s.Title)
		fmt.Println("Description:", s.Description)
		fmt.Println("Command:", s.Command)
		if i > 4 {
			break
		}
	}
	return nil
}
