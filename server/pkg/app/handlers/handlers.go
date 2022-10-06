package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"server/config"
	"server/pkg/api"
	"strconv"
	"strings"
)

// m type builds a map structure quickly to send to a Responder
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

// Handler is the interface that wraps the http.HandlerFunc methods.
type Handler interface {
	Compiler() http.HandlerFunc
	SaveProgram() http.HandlerFunc
	GetProgram() http.HandlerFunc
	GetProgramList() http.HandlerFunc
}

type handler struct {
	program api.ProgramService
	config  *config.Config
}

// NewHandler returns a new Handler
func NewHandler(programService api.ProgramService, config *config.Config) Handler {
	return &handler{
		program: programService,
		config:  config,
	}
}

// SaveProgram handles POST /program
func (h *handler) SaveProgram() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// check content-type
		if r.Header.Get("Content-Type") != "application/json" {
			m := m{
				"error":      http.StatusText(http.StatusUnsupportedMediaType),
				"statusCode": 415,
			}
			wJSON(w, m, http.StatusUnsupportedMediaType)
			return
		}

		// check there is no query in path
		urlValues := r.URL.Query()
		if len(urlValues) > 0 {
			m := m{
				"error":      "No query allowed in URL",
				"statusCode": 400,
			}
			wJSON(w, m, http.StatusBadRequest)
			return
		}

		program := &api.Program{}

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		err := decoder.Decode(&program)
		if err != nil {
			log.Printf("Error Decode: %v", err)
			m := m{
				"error":      http.StatusText(http.StatusBadRequest),
				"statusCode": 400,
			}
			wJSON(w, m, http.StatusBadRequest)
			return
		}

		response, err := h.program.New(*program)
		if err != nil {
			if err.Error() == "name already exists" {
				m := m{
					"error":      err.Error(),
					"statusCode": 409,
				}
				wJSON(w, m, http.StatusConflict)
				return
			}
			if err.Error() == "name required" || err.Error() == "program required" {
				m := m{
					"error":      err.Error(),
					"statusCode": 400,
				}
				wJSON(w, m, http.StatusBadRequest)
				return
			}
			log.Printf("Error save program: %v", err)
			m := m{
				"error":      http.StatusText(http.StatusInternalServerError),
				"statusCode": 500,
			}
			wJSON(w, m, http.StatusInternalServerError)
			return
		}

		m := m{
			"successful": program.Name + " created",
			"uid":        response,
			"statusCode": 201,
		}
		wJSON(w, m, http.StatusCreated)
	}
}

// GetProgram handles GET /program/{uid}
func (h *handler) GetProgram() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		urlValues := r.URL.Query()
		uid := urlValues.Get("uid")

		//check valid query
		if len(urlValues) != 1 || uid == "" {
			m := m{
				"error":      http.StatusText(http.StatusBadRequest),
				"statusCode": 400,
			}
			wJSON(w, m, http.StatusBadRequest)
			return
		}

		uidL := strings.ToLower(uid)

		// check if uid is hex in 0x format or is int type
		if !strings.HasPrefix(uidL, "0x") {

			// check if uid is int
			_, err := strconv.Atoi(uidL)
			if err != nil {
				m := m{
					"error":      uid + " is invalid",
					"statusCode": 400,
				}
				wJSON(w, m, http.StatusBadRequest)
				return
			}

		} else {

			//chech if uid is hex
			uidSplit := strings.Split(uidL, "x")[1]
			_, err := strconv.ParseUint(uidSplit, 16, 32)
			if err != nil {
				m := m{
					"error":      uid + " is invalid",
					"statusCode": 400,
				}
				wJSON(w, m, http.StatusBadRequest)
				return
			}
		}

		response, err := h.program.Get(uid)
		if err != nil {
			log.Printf("error get program: %v", err)
			m := m{
				"error":      http.StatusText(http.StatusInternalServerError),
				"statusCode": 500,
			}
			wJSON(w, m, http.StatusInternalServerError)
			return
		}

		if response[0].Name == "" {
			m := m{
				"error":      uid + " Not Found",
				"statusCode": 404,
			}
			wJSON(w, m, http.StatusNotFound)
			return
		}
		m := m{
			"program":    response[0],
			"statusCode": 200,
		}
		wJSON(w, m, http.StatusOK)
	}
}

// GetProgramList handles GET /program
func (h *handler) GetProgramList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// check there is no query in path
		urlValues := r.URL.Query()
		if len(urlValues) > 0 {
			m := m{
				"error":      "No query allowed in URL",
				"statusCode": 400,
			}
			wJSON(w, m, http.StatusBadRequest)
			return
		}

		response, err := h.program.GetList()
		if err != nil {
			log.Printf("error get programs: %v", err)
			m := m{
				"error":      http.StatusText(http.StatusInternalServerError),
				"statusCode": 500,
			}
			wJSON(w, m, http.StatusInternalServerError)
			return
		}

		if len(response) == 0 {
			m := m{
				"message":    "No saved programs",
				"statusCode": 200,
			}
			wJSON(w, m, http.StatusOK)
			return
		}

		m := m{
			"programs":   response,
			"statusCode": 200,
		}
		wJSON(w, m, http.StatusOK)
	}

}
