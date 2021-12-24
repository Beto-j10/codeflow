package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"server/config"
	"testing"
)

// CompilerResp is the response from the compiler
type CompilerResp struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	Script       string `json:"script"`
	Language     string `json:"language"`
	VersionIndex string `json:"versionIndex"`
}

// CompilerReq is the request body for the compiler
type CompilerReq struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	Script       string `json:"script"`
	Language     string `json:"language"`
	VersionIndex string `json:"versionIndex"`
}

// TestCompiler tests the compiler handler and compiler client
func TestCompilerClient(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//check request method
		if r.Method != http.MethodPost {
			m := m{
				"error":      http.StatusText(http.StatusMethodNotAllowed),
				"statusCode": 405,
			}
			wJSON(w, m, http.StatusMethodNotAllowed)

			return
		}

		// check request content-type
		if r.Header.Get("Content-Type") != "application/json" {
			m := m{
				"error":      http.StatusText(http.StatusUnsupportedMediaType),
				"statusCode": 415,
			}
			wJSON(w, m, http.StatusUnsupportedMediaType)
			return
		}

		var c CompilerResp

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&c)
		if err != nil {
			m := m{
				"error":      err.Error(),
				"statusCode": 400,
			}
			wJSON(w, m, http.StatusBadRequest)
			return
		}

		m := m{
			"output":     "Compiled code",
			"statusCode": "200",
			"memory":     "7884",
			"cpuTime":    "0.01",
		}
		wJSON(w, m, http.StatusOK)
	}))
	defer ts.Close()

	tests := []struct {
		name        string
		compiler    CompilerReq
		contentType string
		want        int
	}{
		{
			name: "successful request",
			compiler: CompilerReq{
				ClientID:     "test 123",
				ClientSecret: "test 456",
				Script:       "test script",
				Language:     "python",
				VersionIndex: "test 1",
			},
			contentType: "application/json",
			want:        200,
		},
		{
			name: "invalid because data is missing",
			compiler: CompilerReq{
				ClientID:     "test ID",
				ClientSecret: "test Secret",
				Script:       "test script",
				Language:     "test Language",
				VersionIndex: "",
			},
			contentType: "application/json",
			want:        400,
		},
		{
			name: "invalid due Content-Type",
			compiler: CompilerReq{
				ClientID:     "test ID",
				ClientSecret: "test Secret",
				Script:       "test script",
				Language:     "test Language",
				VersionIndex: "test Index",
			},
			contentType: "text/plain",
			want:        415,
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			dataReq, err := json.Marshal(test.compiler)
			if err != nil {
				t.Errorf("expected error to be nil, got %v", err)
				return
			}
			req := httptest.NewRequest(http.MethodPost, "/compiler", bytes.NewBuffer(dataReq))
			req.Header.Set("Content-Type", test.contentType)

			w := httptest.NewRecorder()
			h := handler{
				config: &config.Config{
					Compiler: config.CompilerConfig{
						URL:          ts.URL,
						ClientId:     test.compiler.ClientID,
						ClientSecret: test.compiler.ClientSecret,
					},
				},
			}
			hf := http.HandlerFunc(h.Compiler())
			hf.ServeHTTP(w, req)

			response := w.Result()
			defer response.Body.Close()

			var body map[string]interface{}

			json.NewDecoder(response.Body).Decode(&body)
			if err != nil {
				t.Errorf("expected error to be nil, got %v", err)
			}

			if response.StatusCode != test.want {
				t.Errorf("test: - %s - failed. got: %d, want: %d, msg: %v", test.name, response.StatusCode, test.want, body)
			}
		})

	}

}
