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

			resp := w.Result()
			defer resp.Body.Close()

			var body map[string]interface{}

			json.NewDecoder(resp.Body).Decode(&body)
			if err != nil {
				t.Errorf("expected error to be nil, got %v", err)
			}

			if resp.StatusCode != test.want {
				t.Errorf("test: - %s - failed. got: %d, want: %d, body: %v", test.name, resp.StatusCode, test.want, body)
			}
		})

	}

}
