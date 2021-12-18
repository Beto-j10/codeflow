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
