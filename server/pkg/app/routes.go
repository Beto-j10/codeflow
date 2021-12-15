package app

import (
	"github.com/go-chi/chi/v5"
)

func (s *Server) Routes() *chi.Mux {
	router := s.router
	router.Route("/v1/api", func(router chi.Router) {
		router.Post("/compiler", s.compiler())
		router.Post("/program", s.saveProgram())
	})
	return router
}
