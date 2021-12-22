package main

import (
	"fmt"
	"log"
	"os"
	a "server/pkg/api"
	"server/pkg/app"
	"server/pkg/app/handlers"
	"server/pkg/repository"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Startup error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	client := newClientDgraph()
	storage := repository.NewStorage(client)

	err := storage.SetupRepository()
	if err != nil {
		return err
	}

	programService := a.NewProgramService(storage)
	handler := handlers.NewHandler(programService)
	router := chi.NewRouter()
	server := app.NewServer(router, handler)

	// start server
	err = server.Run()
	if err != nil {
		return err
	}
	return nil
}

func newClientDgraph() *dgo.Dgraph {
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(conn),
	)
}
