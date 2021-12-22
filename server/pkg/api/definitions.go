package api

// Program represents de data of a program in Dgraph.
// If Uid field is empty when running a mutation for
// a new node, Dgraph sets one.
type Program struct {
	Uid     string `json:"uid,omitempty"`
	Name    string `json:"name,omitempty"`
	Program string `json:"program,omitempty"`
	// DType   []string `json:"dgraph.type,omitempty"`
}

type ProgramList struct {
	Uid  string `json:"uid"`
	Name string `json:"name"`
}

// GetPrograms returns an array for the query type.
// Json tag must be used for the query name
type GetPrograms struct {
	Program []Program     `json:"program"`
	List    []ProgramList `json:"programList"`
}

var (
	Schema = `
		name: string @index(hash) @upsert .
		program: string .
	`
)
