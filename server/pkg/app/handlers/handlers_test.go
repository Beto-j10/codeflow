package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"server/pkg/api"
	"testing"
)

type mockProgramService struct{}

func (m *mockProgramService) New(program api.Program) (string, error) {

	if program.Name == "" {
		return "", errors.New("name required")
	}
	if program.Program == "" {
		return "", errors.New("program required")
	}

	// mock name already exists
	if program.Name == "test program name already created" {
		return "", errors.New("name already exists")
	}

	// mock successfully saved and return uid
	return "uid test", nil
}
func (m *mockProgramService) Get(getBy string) ([]api.Program, error) {

	//mock resource found successfully
	p := api.GetPrograms{
		Program: []api.Program{
			{
				Uid:     "uid response db",
				Name:    "name response db",
				Program: "program response db",
			},
		},
	}

	// mock resource not found
	if getBy == "0xb" {
		p.Program[0].Name = ""
		return p.Program, nil
	}

	return p.Program, nil
}
func (m *mockProgramService) GetList() ([]api.ProgramList, error) {
	return nil, nil
}

func TestSaveProgram(t *testing.T) {

	mockProgram := mockProgramService{}
	mockHandler := NewHandler(&mockProgram)

	type SaveReq struct {
		Name    string `json:"name"`
		Program string `json:"program"`
	}

	tests := []struct {
		name        string
		save        SaveReq
		path        string
		contentType string
		want        int
	}{
		{
			name: "save a program successfully",
			save: SaveReq{
				Name:    "name test",
				Program: "program test",
			},
			path:        `/program`,
			contentType: "application/json",
			want:        201,
		},
		{
			name: "Error due Content-Type",
			save: SaveReq{
				Name:    "name test",
				Program: "program test",
			},
			path:        `/program`,
			contentType: "text/plain",
			want:        415,
		},
		{
			name: "Error due to query in url ",
			save: SaveReq{
				Name:    "name test",
				Program: "program test",
			},
			path:        `/program?uid=123`,
			contentType: "application/json",
			want:        400,
		},
		{
			name: "Error because the name already exists",
			save: SaveReq{
				Name:    "test program name already created",
				Program: "program test",
			},
			path:        `/program`,
			contentType: "application/json",
			want:        409,
		},
		{
			name: "Error because the name is missing",
			save: SaveReq{
				Name:    "",
				Program: "program test",
			},
			path:        `/program`,
			contentType: "application/json",
			want:        400,
		},
		{
			name: "Error because the program is missing",
			save: SaveReq{
				Name:    "program test",
				Program: "",
			},
			path:        `/program`,
			contentType: "application/json",
			want:        400,
		},
	}
	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			dataReq, err := json.Marshal(test.save)
			if err != nil {
				t.Errorf("expected error to be nil, got %v", err)
				return
			}
			req := httptest.NewRequest(http.MethodPost, (test.path), bytes.NewBuffer(dataReq))
			req.Header.Set("Content-Type", test.contentType)

			w := httptest.NewRecorder()
			hf := http.HandlerFunc(mockHandler.SaveProgram())
			hf.ServeHTTP(w, req)

			response := w.Result()
			defer response.Body.Close()

			var body map[string]interface{}

			err = json.NewDecoder(response.Body).Decode(&body)
			if err != nil {
				t.Errorf("expected error to be nil, got %v", err)
			}

			if response.StatusCode != test.want {
				t.Errorf("test: - %s - failed. got: %d, want: %d, body: %v", test.name, response.StatusCode, test.want, body)
			}
		})

	}

}

func TestGetProgram(t *testing.T) {
	mockProgram := mockProgramService{}
	mockHandler := NewHandler(&mockProgram)

	tests := []struct {
		name string
		path string
		want int
	}{
		{
			name: "return a program succesfully",
			path: `/program?uid=0x1`,
			want: 200,
		},
		{
			name: "error due to resource not found",
			path: "/program?uid=0xb",
			want: 404,
		},
		{
			name: "error due to bad request, bad query in url",
			path: "/program?ui=0x1",
			want: 400,
		},
		{
			name: "error due to bad request, good query key but value is missing",
			path: "/program?uid=",
			want: 400,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, test.path, nil)
			req.Header.Set("Accept", "application/json")

			w := httptest.NewRecorder()
			//TODO: delete
			// h := handler{}
			// hf := http.HandlerFunc(h.GetProgram())
			hf := http.HandlerFunc(mockHandler.GetProgram())
			hf.ServeHTTP(w, req)

			response := w.Result()
			defer response.Body.Close()

			var body map[string]interface{}

			err := json.NewDecoder(response.Body).Decode(&body)
			if err != nil {
				t.Errorf("expected error to be nil, got %v", err)
			}

			if response.StatusCode != test.want {
				t.Errorf("test: - %s - failed. got: %d, want: %d, body: %v", test.name, response.StatusCode, test.want, body)

			}
		})

	}
}
