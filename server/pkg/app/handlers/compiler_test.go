package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type CompilerTest struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	Script       string `json:"script"`
	Language     string `json:"language"`
	VersionIndex string `json:"versionIndex"`
}

func TestCompilerClient(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		//check request method
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			res, _ := json.Marshal(map[string]string{"msg": "Method Not Allowed"})
			w.Write([]byte(res))
			return
		}

		// check request content-type
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			res, _ := json.Marshal(map[string]string{"msg": "Unsupported Media Type"})
			w.Write([]byte(res))
			return
		}

		var c CompilerTest

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&c)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			res, _ := json.Marshal(map[string]string{"msg": err.Error()})
			w.Write([]byte(res))
			return
		}
		// check all data is filled in
		// helps support the compiler data structure
		v := reflect.ValueOf(c)
		for i := 0; i < v.NumField(); i++ {
			if reflect.DeepEqual(v.Field(i).Interface(), reflect.Zero(v.Field(i).Type()).Interface()) {
				msg := fmt.Sprintf("empty field: %v", v.Type().Field(i).Tag)
				w.WriteHeader(http.StatusBadRequest)
				res, _ := json.Marshal(map[string]string{"msg": msg})
				w.Write([]byte(res))
				return
			}
		}
		res, _ := json.Marshal(map[string]string{"msg": "Ok"})
		w.Write([]byte(res))
	}))
	defer ts.Close()

	test := struct {
		name     string
		compiler Compiler
		want     int
	}{

		name: "successful request",
		compiler: Compiler{
			ClientID:     "test 123",
			ClientSecret: "test 456",
			Script:       "test script",
			Language:     "python",
			VersionIndex: "test 1",
		},
		want: 200,
	}

	response, err := test.compiler.CompilerClientURL(ts.URL)
	if err != nil {
		t.Errorf("Expected error to be nil, got: %v", err)
	}

	if response.StatusCode != test.want {
		t.Errorf("test: - %s - failed. got: %d, wanted: %d, msg: %s", test.name, response.StatusCode, test.want, response.Body["msg"])
	}

}
