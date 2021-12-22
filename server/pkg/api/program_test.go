package api

import (
	"errors"
	"testing"
)

type mockProgramRepository struct{}

func (m mockProgramRepository) SaveProgram(program Program) (string, error) {

	// mock name already exists
	if program.Name == "test program name already created" {
		return "", errors.New("name already exists")
	}

	// mock successfully saved and return uid
	return "uid test", nil
}

func (m mockProgramRepository) GetProgram(getBy string) ([]Program, error) {

	//mock resource found successfully
	p := GetPrograms{
		Program: []Program{
			{
				Uid:     "uid response db",
				Name:    "name response db",
				Program: "program response db",
			},
		},
	}

	// mock resource not found
	if getBy == "no uid exists in db" {
		p.Program[0].Name = ""
		return p.Program, nil
	}

	return p.Program, nil
}

func (m mockProgramRepository) GetProgramList() ([]ProgramList, error) {
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
		{
			name: "should return an error because the name already exists",
			program: Program{
				Name:    "test program name already created",
				Program: "code to compile",
			},
			want: errors.New("name already exists"),
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

func TestGet(t *testing.T) {
	mockDB := mockProgramRepository{}
	mockProgramService := NewProgramService(&mockDB)

	tests := []struct {
		name  string
		getBy string
		want  string
	}{
		{
			name:  "run query successfully and return error nil",
			getBy: "uid test",
			want:  "name response db",
		},
		{
			name:  "should return error because no resource found",
			getBy: "no uid exists in db",
			want:  "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response, err := mockProgramService.Get(test.getBy)
			if err != nil {
				t.Errorf("expected error to be nil, got %v", err)
			}

			if response[0].Name != test.want {
				t.Errorf("test: %s failed. got: %s, wanted: %v", test.name, response[0].Name, test.want)
			}
		})
	}

}

func TestGetList(t *testing.T) {
	mockDB := mockProgramRepository{}
	mockProgramService := NewProgramService(&mockDB)

	_, err := mockProgramService.GetList()
	if err != nil {
		t.Errorf("got: %v, want: nil ", err)
	}
}
