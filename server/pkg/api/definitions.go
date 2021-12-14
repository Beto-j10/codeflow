package api

type Program struct {
	Uid     string `json:"uid,omitempty"`
	Name    string `json:"name,omitempty"`
	Program string `json:"program,omitempty"`
}

var (
	Schema = `
		name: string @index(fulltext) .
		program: string .
	`
)
