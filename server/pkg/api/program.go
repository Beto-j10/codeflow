package api

import "errors"

// ProgramService contains the mothods of the program service
type ProgramService interface {
	New(program Program) (string, error)
	Get(getBy string) ([]Program, error)
	GetList() ([]ProgramList, error)
}

// ProgramRepository lets program service do db operations
type ProgramRepository interface {
	SaveProgram(Program) (string, error)
	GetProgram(getBy string) ([]Program, error)
	GetProgramList() ([]ProgramList, error)
}

type programService struct {
	storage ProgramRepository
}

// NewProgramService creates a new program service
func NewProgramService(programRepository ProgramRepository) ProgramService {
	return &programService{
		storage: programRepository,
	}
}

// New tells db to save a new program and returns the uid
func (p *programService) New(program Program) (string, error) {

	if program.Name == "" {
		return "", errors.New("name required")
	}

	if program.Program == "" {
		return "", errors.New("program required")
	}

	response, err := p.storage.SaveProgram(program)
	if err != nil {
		return "", err
	}
	return response, nil
}

// Get tells db to get a program by uid and returns an array with one program
func (p *programService) Get(getBy string) ([]Program, error) {
	response, err := p.storage.GetProgram(getBy)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// GetList tells db to get all programs and returns an array with all programs
func (p *programService) GetList() ([]ProgramList, error) {
	response, err := p.storage.GetProgramList()
	if err != nil {
		return nil, err
	}
	return response, nil
}
