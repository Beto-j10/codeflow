package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Compiler handles POST /compiler
func (h *handler) Compiler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Check content-type
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

		response, err := dataCompiler.compilerClient(&h.config.Compiler)
		if err != nil {
			if strings.HasPrefix(err.Error(), "bad request:") {
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

		// Check text/plain
		if response.Header.Get("Content-Type") == "text/plain" {
			body := fmt.Sprintf("%v", response.Body)
			wJSON(w, m{"compiler": body}, response.StatusCode)
			return
		}

		wJSON(w, m{"compiler": response.Body}, response.StatusCode)
	}
}
