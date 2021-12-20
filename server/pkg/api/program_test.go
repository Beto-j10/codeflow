package api

import (
	"errors"
	"testing"
)

//TODO: study test mock
type mockProgramRepository struct{}

func (m mockProgramRepository) SaveProgram(request Program) error {
	if request.Name == "test program name already created" {
		return errors.New("name already exists")
	}
	return nil
}

func (m mockProgramRepository) GetProgram(request Program) ([]Program, error) {
	if request.Uid == "no uid exists in db" {
		return nil, errors.New("not found")
	}
	p := GetPrograms{
		Program: []Program{
			Program{
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
			want: errors.New("name required"),
		},
		{
			name: "should return an error because the name already exists",
			request: Program{
				Name:    "test program name already created",
				Program: "code to compile",
			},
			want: errors.New("name already exists"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := mockProgramService.New(test.request)
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
		name    string
		request Program
		want    error
	}{
		{
			name: "run query successfully and return error nil",
			request: Program{
				Uid: "uid test",
			},
			want: nil,
		},
		{
			name: "should return an error because Uid is missing",
			request: Program{
				Uid: "",
			},
			want: errors.New("uid required"),
		},
		{
			name: "should return error because no resource found",
			request: Program{
				Uid: "no uid exists in db",
			},
			want: errors.New("not found"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response, err := mockProgramService.Get(test.request)
			if !errors.Is(err, test.want) && err.Error() != test.want.Error() {
				t.Errorf("test: %s failed. got: %v, wanted: %v, response: %v ", test.name, err, test.want, response)
			}
		})
	}

}
