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
func (m *mockProgramService) Get(api.Program) ([]api.Program, error) {
	return nil, nil
}
func (m *mockProgramService) GetList() ([]api.ProgramList, error) {
	return nil, nil
}

type SaveReq struct {
	Name    string `json:"name,omitempty"`
	Program string `json:"program,omitempty"`
}

func TestSave(t *testing.T) {

	mockProgram := mockProgramService{}
	mockHandler := NewHandler(&mockProgram)

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
