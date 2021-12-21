package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"server/pkg/api"
	"testing"
)

type mockProgramService struct{}

func (m *mockProgramService) New(api.Program) error {
	return nil
}
func (m *mockProgramService) Get(string) ([]api.Program, error) {
	return nil, nil
}
func (m *mockProgramService) GetList() ([]api.ProgramList, error) {
	return nil, nil
}

func TestSave(t *testing.T) {

	mockProgram := mockProgramService{}
	mockHandler := NewHandler(&mockProgram)

	type SaveReq struct {
		Name    string `json:"name"`
		Program string `json:"program"`
	}

	tests := []struct {
		name        string
		save        SaveReq
		contentType string
		want        int
	}{
		{
			name: "save a program successfully",
			save: SaveReq{
				Name:    "name test",
				Program: "program test",
			},
			contentType: "application/json",
			want:        200,
		},
		{
			name: "invalid due Content-Type",
			save: SaveReq{
				Name:    "name test",
				Program: "program test",
			},
			contentType: "text/plain",
			want:        415,
		},
	}
	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			dataReq, err := json.Marshal(test.save)
			if err != nil {
				t.Errorf("expected error to be nil, got %v", err)
				return
			}
			req := httptest.NewRequest(http.MethodPost, ("/save"), bytes.NewBuffer(dataReq))
			req.Header.Set("Content-Type", test.contentType)

			w := httptest.NewRecorder()
			hf := http.HandlerFunc(mockHandler.SaveProgram())
			hf.ServeHTTP(w, req)

			response := w.Result()
			defer response.Body.Close()

			var body map[string]interface{}

			json.NewDecoder(response.Body).Decode(&body)
			if err != nil {
				t.Errorf("expected error to be nil, got %v", err)
			}

			if response.StatusCode != test.want {
				t.Errorf("test: - %s - failed. got: %d, want: %d, body: %v", test.name, response.StatusCode, test.want, body)
			}
		})

	}

}

func TestGet(t *testing.T) {
	mockProgram := mockProgramService{}
	mockHandler := NewHandler(&mockProgram)

	tests := []struct {
		name string
		path string
		want int
	}{
		{
			name: "return a program succesfully",
			path: `/compiler?uid=2&yy=4`,
			want: 200,
		},
		{
			name: "error due to resource not found",
			path: "/compiler?uid=0x1",
			want: 404,
		},
		{
			name: "error due to bad request",
			path: "/compiler?ui=uid1",
			want: 400,
		},
		{
			name: "error due to bad request",
			path: "/compiler?",
			want: 400,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, test.path, nil)
			req.Header.Set("Accept", "application/json")

			w := httptest.NewRecorder()
			// h := handler{}
			// hf := http.HandlerFunc(h.GetProgram())
			hf := http.HandlerFunc(mockHandler.GetProgram())
			hf.ServeHTTP(w, req)

			// response := w.Result()
			// defer response.Body.Close()
			// t.Fatal("kkkkk")
		})
	}
}
