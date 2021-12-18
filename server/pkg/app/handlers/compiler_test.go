package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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
			http.Error(w, "Server - method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		// check request content-type
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Server - Unsupported Media Type", http.StatusUnsupportedMediaType)
			return
		}

		var c CompilerResp

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&c)
		if err != nil {
			http.Error(w, "Server - "+err.Error(), http.StatusBadRequest)
			return
		}

		// check all data is filled in
		// helps support the compiler data structure
		v := reflect.ValueOf(c)
		for i := 0; i < v.NumField(); i++ {
			if reflect.DeepEqual(v.Field(i).Interface(), reflect.Zero(v.Field(i).Type()).Interface()) {
				msg := fmt.Sprintf("Server - empty field: %v", v.Type().Field(i).Tag)
				http.Error(w, msg, http.StatusBadRequest)
				return
			}
		}
		http.Error(w, "Ok", http.StatusOK)
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

			resp := w.Result()
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("expected error to be nil, got %v", err)
			}
			if resp.StatusCode != test.want {
				t.Errorf("test: - %s - failed. got: %d, want: %d, msg: %s", test.name, resp.StatusCode, test.want, string(body))
			}
		})

	}

}
