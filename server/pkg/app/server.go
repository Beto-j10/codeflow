package app

import (
	"flag"
	"log"
	"net/http"
	"time"

	"server/pkg/app/handlers"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Server struct {
	router  *chi.Mux
	handler handlers.Handler
}

func NewServer(router *chi.Mux, handler handlers.Handler) *Server {
	return &Server{
		router:  router,
		handler: handler,
	}
}

//TODO: add auto documentation
func (s *Server) Run() error {

	port := flag.String("port", "8080", "server port")
	flag.Parse()

	s.config()
	router := s.Routes()

	err := http.ListenAndServe(":"+*port, router)
	if err != nil {
		log.Printf("Error running serve: %v", err)
		return err
	}
	return nil
}

func (s *Server) config() {

	// stack middleware
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.Timeout(15 * time.Second))

	s.router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type"},
	}))
}
