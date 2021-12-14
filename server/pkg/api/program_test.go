package api

import (
	"errors"
	"testing"
)

//TODO: study test mock
type mockProgramRepository struct{}

func (m mockProgramRepository) SaveProgram(request Program) error {
	if request.Name == "test program name already created" {
		return errors.New("DB - name already exists")
	}
	return nil
}

func TestSaveProgram(t *testing.T) {
	mockDB := mockProgramRepository{}
	mockProgramService := NewProgramService(&mockDB)

	tests := []struct {
		name    string
		request Program
		want    error
	}{
		{
			name: "should save a program succesfully",
			request: Program{
				Name:    "test program",
				Program: "code to compile",
			},
			want: nil,
		},
		{
			name: "should return an error because the name is missing",
			request: Program{
				Name:    "",
				Program: "code to compile",
			},
			want: errors.New("Program service - name required"),
		},
		{
			name: "should return an error because the name already exists",
			request: Program{
				Name:    "test program name already created",
				Program: "code to compile",
			},
			want: errors.New("DB - name already exists"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := mockProgramService.New(test.request)
			if !errors.Is(err, test.want) && err.Error() != test.want.Error() {
				t.Errorf("test: %v failed. got: %v, wanted: %v", test.name, err, test.want)
			}
		})
	}
}
