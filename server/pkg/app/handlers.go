package app

import (
	"encoding/json"
	"log"
	"net/http"
	"server/pkg/api"
	"server/pkg/app/client"

	"github.com/go-chi/render"
)

//TODO: CHECK ERRORS REQUEST //DisallowUnknownFields
func (s *Server) compiler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// check content-type
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}

		dataCompiler := &client.Compiler{}
		err := json.NewDecoder(r.Body).Decode(&dataCompiler)
		if err != nil {
			log.Printf("data read error: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		response, err := dataCompiler.CompilerClient()
		if err != nil {
			log.Printf("Error request compiler: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, response.Body)
	}
}

func (s *Server) saveProgram() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// check content-type
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}

		program := &api.Program{}
		err := json.NewDecoder(r.Body).Decode(&program)
		if err != nil {
			log.Printf("data read error: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// TODO: check error
		s.program.New(*program)
	}
}
