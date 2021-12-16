package handlers

import (
	"net/http"
	"server/pkg/api"
)

type Handler interface {
	Compiler() http.HandlerFunc
	SaveProgram() http.HandlerFunc
}

type handler struct {
	program api.ProgramService
}

func NewHandler(programService api.ProgramService) Handler {
	return &handler{
		program: programService,
	}
}

//TODO: CHECK ERRORS REQUEST //DisallowUnknownFields
