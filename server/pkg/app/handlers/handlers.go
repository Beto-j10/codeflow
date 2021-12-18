package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"server/pkg/api"
)

// m type builds a map structure quickly to send a Responder
type m map[string]interface{}

// WJSON marshals 'M' to JSON and setting the
// Content-Type as application/json.
func wJSON(w http.ResponseWriter, response m, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	r, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error Marshal: %v", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)
	w.Write(r)
}

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
