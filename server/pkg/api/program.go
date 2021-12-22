package api

import "errors"

//TODO:Check pointer Program in ProgramService/New/handlers/storage
type ProgramService interface {
	New(program Program) (string, error)
	Get(getBy string) ([]Program, error)
	GetList() ([]ProgramList, error)
}

type ProgramRepository interface {
	SaveProgram(Program) (string, error)
	GetProgram(getBy string) ([]Program, error)
	GetProgramList() ([]ProgramList, error)
}

type programService struct {
	storage ProgramRepository
}

func NewProgramService(programRepository ProgramRepository) ProgramService {
	return &programService{
		storage: programRepository,
	}
}

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

func (p *programService) Get(getBy string) ([]Program, error) {
	response, err := p.storage.GetProgram(getBy)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (p *programService) GetList() ([]ProgramList, error) {
	response, err := p.storage.GetProgramList()
	if err != nil {
		return nil, err
	}
	return response, nil
}
