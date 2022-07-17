package models

//define TestCase struct
type TestCase struct {
	Identity string `json:"identity"`
	ProblemIdentity string `json:"problem_identity"`
	Input   string `json:"input"`
	Output  string `json:"output"`
}

func (table *TestCase)TableName() string {
	return "test_case"
}