package api

import "errors"

//TODO:Check pointer Program in ProgramService/New/handlers/storage
type ProgramService interface {
	New(program Program) error
	// Get()
	// GetList()
}

type ProgramRepository interface {
	SaveProgram(Program) error
	GetProgram(program Program) ([]Program, error)
	// GetListPrograms(program Program) error
}

type programService struct {
	storage ProgramRepository
}

func NewProgramService(programRepository ProgramRepository) ProgramService {
	return &programService{
		storage: programRepository,
	}
}

//TODO: add normalization an validations
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

func (p *programService) Get(program Program) ([]Program, error) {
	p.storage.GetProgram(program)
	return nil, nil
}
