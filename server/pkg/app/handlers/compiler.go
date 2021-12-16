package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"server/pkg/app/client"

	"github.com/go-chi/render"
)

func (h *handler) Compiler() http.HandlerFunc {
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
