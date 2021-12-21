package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"server/pkg/api"
)

// m type builds a map structure quickly to send a Responder
type m map[string]interface{}

// wJSON marshals 'm' to JSON and setting the
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
	GetProgram() http.HandlerFunc
	GetProgramList() http.HandlerFunc
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

func (h *handler) SaveProgram() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// check content-type
		if r.Header.Get("Content-Type") != "application/json" {
			wJSON(w, m{"error": http.StatusText(http.StatusUnsupportedMediaType), "statusCode": 415}, http.StatusUnsupportedMediaType)
			return
		}

		program := &api.Program{}

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&program)
		if err != nil {
			log.Printf("Error Decode: %v", err)
			wJSON(w, m{"error": http.StatusText(http.StatusBadRequest), "statusCode": 400}, http.StatusBadRequest)
			return
		}
		// TODO: check error
		h.program.New(*program)
	}
}

func (h *handler) GetProgram() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

	}
}

func (h *handler) GetProgramList() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

	}
}
