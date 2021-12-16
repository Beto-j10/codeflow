package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"server/pkg/api"
)

func (h *handler) SaveProgram() http.HandlerFunc {
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
		h.program.New(*program)
	}
}
