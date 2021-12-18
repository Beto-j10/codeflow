package handlers

import (
	"encoding/json"
	"net/http"
	"server/pkg/api"
)

func (h *handler) SaveProgram() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// check content-type
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Server - Unsupported Media Type", http.StatusUnsupportedMediaType)
			return
		}

		program := &api.Program{}
		// err := json.NewDecoder(r.Body).Decode(&program)
		// if err != nil {
		// 	log.Printf("data read error: %v", err)
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	return
		// }

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&program)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// TODO: check error
		h.program.New(*program)
	}
}
