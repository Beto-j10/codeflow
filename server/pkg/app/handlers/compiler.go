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
			wJSON(w, m{"error": http.StatusText(http.StatusUnsupportedMediaType), "statusCode": 415}, http.StatusUnsupportedMediaType)
			return
		}

		dataCompiler := &CompilerClient{}
		err := json.NewDecoder(r.Body).Decode(&dataCompiler)
		if err != nil {
			log.Printf("data read error: %v", err)
			wJSON(w, m{"error": http.StatusText(http.StatusBadRequest), "statusCode": 400}, http.StatusBadRequest)
			return
		}

		response, err := dataCompiler.CompilerClient()
		if err != nil {
			log.Printf("Error request compiler: %v", err)
			wJSON(w, m{"error": http.StatusText(http.StatusInternalServerError), "statusCode": 500}, http.StatusInternalServerError)
			return
		}

		//reponse as text/plain
		if response.Header.Get("Content-Type") == "text/plain" {
			body := fmt.Sprintf("%v", response.Body)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(response.StatusCode)
			w.Write([]byte(body))
			return
		}

		//response as JSON
		wJSON(w, m{"compiler": response.Body}, response.StatusCode)
	}
}
