package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"

	a "server/pkg/api"
)

type Storage interface {
	SetupRepository() error
	SaveProgram(program a.Program) (string, error)
	GetProgram(getBy string) ([]a.Program, error)
	GetProgramList() ([]a.ProgramList, error)
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
	op.Schema = a.Schema
	op.RunInBackground = true
	err := s.dgraphClient.Alter(context.Background(), op)
	if err != nil {
		return err
	}
	return nil
}

func (s *storage) SaveProgram(program a.Program) (string, error) {
	ctx := context.Background()
	txn := s.dgraphClient.NewTxn()
	defer txn.Discard(ctx)

	// will be the Key returned if the mutation is successful
	program.Uid = "_:program"

	pb, err := json.Marshal(program)
	if err != nil {
		return "", err
	}
	//TODO: uncommentQuery
	query := `
		query{
			prog as var(func: eq(name, "d"))
		}`
	// query := `
	// 	query P($name: string){
	// 		prog as var(func: eq(name, $name))
	// 	}`

	mu := &api.Mutation{
		Cond:    `@if(eq(len(prog), 0))`,
		SetJson: pb,
	}

	req := &api.Request{
		Query: query,
		//TODO: uncomment vars
		// Vars:      map[string]string{"$name": program.Name},
		Mutations: []*api.Mutation{mu},
		CommitNow: true,
	}

	response, err := txn.Do(ctx, req)
	if err != nil {
		return "", err
	}
	//TODO: delete print
	fmt.Printf("\nRESP: %v\n", response.Uids)

	if len(response.Uids) == 0 {
		return "", errors.New("name already exists")
	}

	//TODO: delete
	responser, err := s.GetProgram("0x77")
	if err != nil {
		fmt.Println("Error:", err)
	}
	if responser[0].Name == "" {
		fmt.Printf("\n####ResponseGet: %+v\n", responser[0].Uid)

	}
	s.GetProgramList()

	uid := response.Uids["program"]

	return uid, nil
}

func (s *storage) GetProgram(getBy string) ([]a.Program, error) {
	ctx := context.Background()
	const query = `query Program($uid: string)
		{
			program(func: uid($uid)) {
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
		Vars:  map[string]string{"$uid": getBy},
	}

	response, err := txn.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	res := &a.GetPrograms{}
	err = json.Unmarshal(response.Json, &res)
	if err != nil {
		return nil, err
	}
	//TODO: delete print
	fmt.Printf("\n####Get: %+v\n", res.Program)
	return res.Program, nil
}

func (s *storage) GetProgramList() ([]a.ProgramList, error) {

	ctx := context.Background()
	const query = `query
		{
			programList(func: has(name)) {
				uid
				name
			}
		}
	`
	txn := s.dgraphClient.NewReadOnlyTxn()
	defer txn.Discard(ctx)

	req := &api.Request{
		Query: query,
	}

	response, err := txn.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	res := &a.GetPrograms{}
	err = json.Unmarshal(response.Json, &res)
	if err != nil {
		return nil, err
	}
	//TODO: delete print
	fmt.Printf("\n####List: %+v\n", res.List)
	return res.List, nil
}
