package professor

import (
	"fmt"

	"github.com/Leoo05/test_go/internal/domain/professor/entities"
	"github.com/Leoo05/test_go/internal/domain/professor/repository"
)

type ProfessorUseCase struct {
	ProfessorRepository repository.IProfessorRepository
}

func NewProfessorCase(professorRepository repository.IProfessorRepository) *ProfessorUseCase {
	return &ProfessorUseCase{
		ProfessorRepository: professorRepository,
	}
}

func (usecase *ProfessorUseCase) PostProfessor(professor *entities.Professor) error {
	err := usecase.ProfessorRepository.CreateProfessor(professor)
	if err != nil {
		return err
	}
	return nil
}

func (usecase *ProfessorUseCase) GetProfessorByID(idProfessor int) (*entities.Professor, error) {
	if idProfessor <= 0 {
		return &entities.Professor{}, fmt.Errorf("error Invalid ID")
	}
	professor, err := usecase.ProfessorRepository.ReadProfessor(idProfessor)
	if err != nil {
		return nil, err
	}
	return professor, err
}

func (usecase *ProfessorUseCase) PutProfessor(idProfessor int, professor *entities.Professor) error {
	if idProfessor <= 0 {
		return fmt.Errorf("error Invalid ID")
	}
	err := usecase.ProfessorRepository.UpdateProfessor(idProfessor, professor)
	if err != nil {
		return err
	}
	return nil
}

func (usecase *ProfessorUseCase) DeleteProfessor(idProfessor int) error {
	if idProfessor <= 0 {
		return fmt.Errorf("error Invalid ID")
	}
	err := usecase.ProfessorRepository.DeleteProfessor(idProfessor)
	if err != nil {
		return err
	}
	return nil
}
