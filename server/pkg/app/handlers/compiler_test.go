package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type CompilerResp struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	Script       string `json:"script"`
	Language     string `json:"language"`
	VersionIndex string `json:"versionIndex"`
}
type CompilerReq struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	Script       string `json:"script"`
	Language     string `json:"language"`
	VersionIndex string `json:"versionIndex"`
}

func TestCompilerClient(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//check request method
		if r.Method != http.MethodPost {
			wJSON(w, m{"error": http.StatusText(http.StatusMethodNotAllowed), "statusCode": 405}, http.StatusMethodNotAllowed)

			return
		}

		// check request content-type
		if r.Header.Get("Content-Type") != "application/json" {
			wJSON(w, m{"error": http.StatusText(http.StatusUnsupportedMediaType), "statusCode": 415}, http.StatusUnsupportedMediaType)
			return
		}

		var c CompilerResp

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&c)
		if err != nil {
			wJSON(w, m{"error": err.Error(), "statusCode": 400}, http.StatusBadRequest)
			return
		}

		// check all data is filled in
		// helps support the compiler data structure
		v := reflect.ValueOf(c)
		for i := 0; i < v.NumField(); i++ {
			if reflect.DeepEqual(v.Field(i).Interface(), reflect.Zero(v.Field(i).Type()).Interface()) {
				msg := fmt.Sprintf("Server - empty field: %v", v.Type().Field(i).Tag)
				wJSON(w, m{"error": msg, "statusCode": "400"}, http.StatusBadRequest)
				return
			}
		}
		wJSON(w, m{"output": "Compiled code", "statusCode": "200", "memory": "7884", "cpuTime": "0.01"}, http.StatusOK)
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

			URL = ts.URL
			w := httptest.NewRecorder()
			h := handler{}
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
