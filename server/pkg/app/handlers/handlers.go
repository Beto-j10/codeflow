package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"server/pkg/api"
	"strconv"
	"strings"
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
			wJSON(w, m{"error": "No query allowed in URL", "statusCode": 400}, http.StatusBadRequest)
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

		response, err := h.program.New(*program)
		if err != nil {
			if err.Error() == "name already exists" {
				wJSON(w, m{"error": err.Error(), "statusCode": 409}, http.StatusConflict)
				return
			}
			if err.Error() == "name required" || err.Error() == "program required" {
				wJSON(w, m{"error": err.Error(), "statusCode": 400}, http.StatusBadRequest)
				return
			}
			log.Printf("Error save program: %v", err)
			wJSON(w, m{"error": http.StatusText(http.StatusInternalServerError), "statusCode": 500}, http.StatusInternalServerError)
			return
		}

		wJSON(w, m{"successful": program.Name + " created", "uid": response, "statusCode": 201}, http.StatusCreated)
	}
}

func (h *handler) GetProgram() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		urlValues := r.URL.Query()
		uid := urlValues.Get("uid")

		//check valid query
		if len(urlValues) != 1 || uid == "" {
			wJSON(w, m{"error": http.StatusText(http.StatusBadRequest), "statusCode": 400}, http.StatusBadRequest)
			return
		}

		uidL := strings.ToLower(uid)

		//TODO: add test cases

		// check if uid is a hex type in 0x format or an int type
		if !strings.HasPrefix(uidL, "0x") {

			// check if uidL is int
			_, err := strconv.Atoi(uidL)
			if err != nil {
				wJSON(w, m{"error": uid + " is invalid", "statusCode": 400}, http.StatusBadRequest)
				return
			}

		} else {

			//chech if it is hex
			uidSplit := strings.Split(uidL, "x")[1]
			_, err := strconv.ParseUint(uidSplit, 16, 32)
			if err != nil {
				wJSON(w, m{"error": uid + " is invalid", "statusCode": 400}, http.StatusBadRequest)
				return
			}
		}

		response, err := h.program.Get(uid)
		if err != nil {
			log.Printf("error get program: %v", err)
			wJSON(w, m{"error": http.StatusText(http.StatusInternalServerError), "statusCode": 500}, http.StatusInternalServerError)
			return
		}

		if response[0].Name == "" {
			wJSON(w, m{"error": uid + " Not Found", "statusCode": 404}, http.StatusNotFound)
			return
		}
		wJSON(w, m{"program": response[0], "statusCode": 200}, http.StatusOK)
	}
}

func (h *handler) GetProgramList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// check there is no query in path
		urlValues := r.URL.Query()
		if len(urlValues) > 0 {
			wJSON(w, m{"error": "No query allowed in URL", "statusCode": 400}, http.StatusBadRequest)
			return
		}

		response, err := h.program.GetList()
		if err != nil {
			log.Printf("error get program: %v", err)
			wJSON(w, m{"error": http.StatusText(http.StatusInternalServerError), "statusCode": 500}, http.StatusInternalServerError)
			return
		}

		if len(response) == 0 {
			wJSON(w, m{"message": "No saved programs", "statusCode": 200}, http.StatusOK)
			return
		}

		wJSON(w, m{"programs": response, "statusCode": 200}, http.StatusOK)
	}

}
