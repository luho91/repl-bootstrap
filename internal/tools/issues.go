package tools

type StateIssue string

const (
	IssueMissingConfig		StateIssue = "missing_config"
	IssueEmptyConfig		StateIssue = "empty_config"
)
 
func EvaluateIssues(f Facts) []StateIssue {
	var issues []StateIssue

	if !f.ConfigExists {
		issues = append(issues, IssueMissingConfig)
	}

	if f.ConfigIsEmpty {
		issues = append(issues, IssueEmptyConfig)
	}

	return issues
}
