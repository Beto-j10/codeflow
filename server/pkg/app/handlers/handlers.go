package handlers

import (
	"encoding/json"
	"fmt"
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

		// check there is no query in path
		urlValues := r.URL.Query()
		if len(urlValues) > 0 {
			wJSON(w, m{"error": "No query string allowed in URL", "statusCode": 400}, http.StatusBadRequest)
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

		err = h.program.New(*program)
		if err != nil {
			if err.Error() == "name already exists" {
				wJSON(w, m{program.Name: "Already exists", "statusCode": 409}, http.StatusConflict)
				return
			}
			log.Printf("Error save program: %v", err)
			wJSON(w, m{"error": http.StatusText(http.StatusInternalServerError), "statusCode": 500}, http.StatusInternalServerError)
			return
		}

		wJSON(w, m{program.Name: "Created", "statusCode": 201}, http.StatusCreated)
	}
}

func (h *handler) GetProgram() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		urlValues := r.URL.Query()
		uid := urlValues.Get("uid")

		//check valid query
		if len(urlValues) == 1 && uid != "" {
			response, err := h.program.Get(uid)
			if err != nil {
				log.Printf("error get program: %v", err)
				wJSON(w, m{"error": http.StatusText(http.StatusInternalServerError), "statusCode": 500}, http.StatusInternalServerError)
				return
			}

			if response[0].Name == "" {
				wJSON(w, m{uid: "Not Found", "statusCode": 404}, http.StatusNotFound)
				return
			}
			wJSON(w, m{"program": response[0]}, http.StatusOK)
			//TODO: delete print
			fmt.Printf("#####%v\n", urlValues)
			return
		}

		wJSON(w, m{"error": http.StatusText(http.StatusBadRequest), "statusCode": 400}, http.StatusBadRequest)
	}
}

func (h *handler) GetProgramList() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

	}
}
