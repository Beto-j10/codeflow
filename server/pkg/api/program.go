package api

import "errors"

//TODO:Check pointer Program in ProgramService/New/handlers/storage
type ProgramService interface {
	New(program Program) error
	Get(getBy Program) ([]Program, error)
	GetList() ([]ProgramList, error)
}

type ProgramRepository interface {
	SaveProgram(Program) error
	GetProgram(getBy Program) ([]Program, error)
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

func (p *programService) New(program Program) error {
	if program.Name == "" {
		return errors.New("name required")
	}
	err := p.storage.SaveProgram(program)
	if err != nil {
		return err
	}
	return nil
}

func (p *programService) Get(getBy Program) ([]Program, error) {
	if getBy.Uid == "" {
		return nil, errors.New("uid required")
	}
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
