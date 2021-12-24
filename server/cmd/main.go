package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"server/config"
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
	config, err := config.LoadConfig("config.yml")
	if err != nil {
		return err
	}

	serverPort := flag.String("port", config.Server.Port, "server port")
	dgraphClient := flag.String("dgraph", config.Dgraph.Port, "dgraph port")
	flag.Parse()

	client := newClientDgraph(*dgraphClient)
	storage := repository.NewStorage(client)

	err = storage.SetupRepository()
	if err != nil {
		return err
	}

	programService := a.NewProgramService(storage)
	handler := handlers.NewHandler(programService, config)
	router := chi.NewRouter()
	server := app.NewServer(router, handler)

	// start server
	err = server.Run(*serverPort)
	if err != nil {
		return err
	}
	return nil
}

func newClientDgraph(port string) *dgo.Dgraph {

	target := fmt.Sprintf("localhost:%s", port)
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected to Dgraph at port %s", port)
	return dgo.NewDgraphClient(
		api.NewDgraphClient(conn),
	)
}
