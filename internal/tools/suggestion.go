package tools


type Suggestion struct {
	Title		string
	Description	string
	Command		string
}

func GetSuggestion(issues []StateIssue) []Suggestion {
	suggestionMap := getSuggestionMap()
	s := []Suggestion{}

	for _, i := range issues {
		s = append(s, suggestionMap[i])
	}

	return s
}

func getSuggestionMap() map[StateIssue]Suggestion {
	suggestions := map[StateIssue]Suggestion{}

	suggestions[IssueMissingConfig] = Suggestion {
		Title:			"Config file is missing",
		Description:	"Couldn't find a config file for this repl in your ~/.config. You should create one!",
		Command:		"create config",
	}

	return suggestions
}
