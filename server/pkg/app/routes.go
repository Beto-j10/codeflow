package app

import (
	"github.com/go-chi/chi/v5"
)

func (s *Server) Routes() *chi.Mux {
	router := s.router
	router.Route("/v1/api", func(router chi.Router) {
		router.Post("/compiler", s.handler.Compiler())
		router.Post("/program", s.handler.SaveProgram())
		router.Get("/program", s.handler.GetProgram())
		router.Get("/program/list", s.handler.GetProgramList())
	})
	return router
}
