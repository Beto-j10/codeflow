package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"

	_api "server/pkg/api"
)

type Storage interface {
	SetupRepository() error
	SaveProgram(program _api.Program) error
	// GetProgram(program _api.Program) error
	// GetListPrograms(program _api.Program) error
}

//TODO check variables name // check log.Fatal
type storage struct {
	dgraphClient *dgo.Dgraph
}

func NewStorage(dgraphClient *dgo.Dgraph) Storage {
	return &storage{
		dgraphClient: dgraphClient,
	}
}

func (s *storage) SetupRepository() error {
	op := &api.Operation{}
	op.Schema = _api.Schema
	op.RunInBackground = true
	err := s.dgraphClient.Alter(context.Background(), op)
	if err != nil {
		return err
	}
	return nil
}

//TODO: add upsert name to the names already exist
func (s *storage) SaveProgram(program _api.Program) error {
	ctx := context.Background()
	txn := s.dgraphClient.NewTxn()
	defer txn.Discard(ctx)

	p := _api.Program{
		// Uid:     "_:program1",
		Name:    "Program1",
		Program: `{"node": {"pp": "tres"}}`,
	}

	pb, err := json.Marshal(p)
	if err != nil {
		return err
	}

	mu := &api.Mutation{
		SetJson: pb,
	}

	req := &api.Request{
		Mutations: []*api.Mutation{mu},
		CommitNow: true,
	}

	resp, err := txn.Do(ctx, req)
	if err != nil {
		return err
	}

	fmt.Println("RESP: ", resp)
	return nil
}

func (s *storage) GetPrograms() {
	ctx := context.Background()
	const query = `query programs($a: string)
		{
			programs(func: anyoftext(name, $a)) {
				uid
				name
				program
			}
		}
	`
	txn := s.dgraphClient.NewReadOnlyTxn()
	defer txn.Discard(ctx)
	req := &api.Request{
		Query: query,
		Vars:  map[string]string{"$a": "programa"},
	}
	resp, err := txn.Do(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s\n", resp)
	var Decode struct {
		Programs []struct {
			Uid     string
			Name    string
			Program string
		}
	}
	if err := json.Unmarshal(resp.GetJson(), &Decode); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", Decode)

}
