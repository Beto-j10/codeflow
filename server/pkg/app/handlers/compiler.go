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
			m := m{
				"error":      http.StatusText(http.StatusUnsupportedMediaType),
				"statusCode": 415,
			}
			wJSON(w, m, http.StatusUnsupportedMediaType)
			return
		}

		dataCompiler := &CompilerClient{}

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		err := decoder.Decode(&dataCompiler)
		if err != nil {
			log.Printf("Error Decode: %v", err)
			m := m{
				"error":      http.StatusText(http.StatusBadRequest),
				"statusCode": 400,
			}
			wJSON(w, m, http.StatusBadRequest)
			return
		}

		response, err := dataCompiler.CompilerClient()
		if err != nil {
			if err.Error() == "bad request: data is missing" {
				m := m{
					"error":      err.Error(),
					"statusCode": 400,
				}
				wJSON(w, m, http.StatusBadRequest)
				return
			}
			log.Printf("Error request compiler: %v", err)
			m := m{
				"error":      http.StatusText(http.StatusInternalServerError),
				"statusCode": 500,
			}
			wJSON(w, m, http.StatusInternalServerError)
			return
		}

		//check text/plain
		if response.Header.Get("Content-Type") == "text/plain" {
			body := fmt.Sprintf("%v", response.Body)
			wJSON(w, m{"compiler": body}, response.StatusCode)
			return
		}

		wJSON(w, m{"compiler": response.Body}, response.StatusCode)
	}
}
