package api

import (
	"errors"
	"testing"
)

//TODO: study test mock
type mockProgramRepository struct{}

func (m mockProgramRepository) SaveProgram(program Program) error {
	if program.Name == "test program name already created" {
		return errors.New("name already exists")
	}
	return nil
}

func (m mockProgramRepository) GetProgram(getBy Program) ([]Program, error) {
	if getBy.Uid == "no uid exists in db" {
		return nil, errors.New("not found")
	}
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

func TestSaveProgram(t *testing.T) {
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
			err := mockProgramService.New(test.program)
			if !errors.Is(err, test.want) && err.Error() != test.want.Error() {
				t.Errorf("test: %s failed. got: %v, wanted: %v", test.name, err, test.want)
			}
		})
	}
}

func TestGetProgram(t *testing.T) {
	mockDB := mockProgramRepository{}
	mockProgramService := NewProgramService(&mockDB)

	tests := []struct {
		name  string
		getBy Program
		want  error
	}{
		{
			name: "run query successfully and return error nil",
			getBy: Program{
				Uid: "uid test",
			},
			want: nil,
		},
		{
			name: "should return an error because Uid is missing",
			getBy: Program{
				Uid: "",
			},
			want: errors.New("uid required"),
		},
		{
			name: "should return error because no resource found",
			getBy: Program{
				Uid: "no uid exists in db",
			},
			want: errors.New("not found"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response, err := mockProgramService.Get(test.getBy)
			if !errors.Is(err, test.want) && err.Error() != test.want.Error() {
				t.Errorf("test: %s failed. got: %v, wanted: %v, response: %v ", test.name, err, test.want, response)
			}
		})
	}

}
