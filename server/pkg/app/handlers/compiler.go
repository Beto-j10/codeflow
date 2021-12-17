package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (h *handler) Compiler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// check content-type
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}

		dataCompiler := &CompilerClient{}
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

		//response as JSON
		if response.Header.Get("Content-Type") == "application/json" {
			body, err := json.Marshal(response.Body)
			if err != nil {
				log.Printf("Error encoding data: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(response.StatusCode)
			w.Write(body)
			return
		}

		//reponse as text/plain
		body := fmt.Sprintf("%v", response.Body)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(response.StatusCode)
		w.Write([]byte(body))
	}
}
