package api

import (
	"errors"
	"testing"
)

type mockProgramRepository struct{}

// SaveProgram is a mock function for testing
func (m mockProgramRepository) SaveProgram(program Program) (string, error) {
	// mock successfully saved and return uid
	return "uid test", nil
}

// GetProgram is a mock function for testing
func (m mockProgramRepository) GetProgram(getBy string) ([]Program, error) {

	//mock resource return
	p := GetPrograms{
		Program: []Program{
			{
				Uid:     "uid response db",
				Name:    "name response db",
				Program: "program response db",
			},
		},
	}

	return p.Program, nil
}

// GetProgramsList is a mock function for testing
func (m mockProgramRepository) GetProgramList() ([]ProgramList, error) {

	//mock resource return
	p := GetPrograms{
		List: []ProgramList{
			{
				Uid:  "uid1 response db",
				Name: "name1 response db",
			},
			{
				Uid:  "uid2 response db",
				Name: "name2 response db",
			},
		},
	}

	return p.List, nil
}

// TestNew tests the New function
func TestNew(t *testing.T) {
	mockDB := mockProgramRepository{}
	mockProgramService := NewProgramService(&mockDB)

	tests := []struct {
		name    string
		program Program
		want    error
	}{
		{
			name: "should save a program succesfully",
			program: Program{
				Name:    "test program",
				Program: "code to compile",
			},
			want: nil,
		},
		{
			name: "should return an error because the name is missing",
			program: Program{
				Name:    "",
				Program: "code to compile",
			},
			want: errors.New("name required"),
		},
		{
			name: "should return an error because the program is missing",
			program: Program{
				Name:    "test program",
				Program: "",
			},
			want: errors.New("program required"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response, err := mockProgramService.New(test.program)
			if !errors.Is(err, test.want) && err.Error() != test.want.Error() {
				t.Errorf("test: %s failed. got: %v, wanted: %v, response: %v", test.name, err, test.want, response)
			}
		})
	}
}

// TestGet tests the Get function
func TestGet(t *testing.T) {
	mockDB := mockProgramRepository{}
	mockProgramService := NewProgramService(&mockDB)

	_, err := mockProgramService.Get("uid")
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}

}

// TestGetList tests the GetList function
func TestGetList(t *testing.T) {
	mockDB := mockProgramRepository{}
	mockProgramService := NewProgramService(&mockDB)

	_, err := mockProgramService.GetList()
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}
}
