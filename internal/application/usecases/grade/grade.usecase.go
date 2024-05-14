package grade

import (
	"fmt"

	"github.com/Leoo05/test_go/internal/domain/grade/entities"
	"github.com/Leoo05/test_go/internal/domain/grade/repository"
)

type GradeUseCase struct {
	GradeRepository repository.IGradesRepository
}

func NewGradeCase(gradeRepository repository.IGradesRepository) *GradeUseCase {
	return &GradeUseCase{
		GradeRepository: gradeRepository,
	}
}

func (usecase *GradeUseCase) PostGrade(grade *entities.Grade) error {
	if grade.Grade < 0 || grade.Grade > 5 {
		return fmt.Errorf("error invalid Grade")
	}
	err := usecase.GradeRepository.CreateGrade(grade)
	if err != nil {
		return err
	}
	return nil
}

func (usecase *GradeUseCase) GetGradeByID(idGrade int) (*entities.Grade, error) {
	if idGrade < 0 {
		return &entities.Grade{}, fmt.Errorf("error invalid idGrade")
	}
	grade, err := usecase.GradeRepository.ReadGrade(idGrade)
	if err != nil {
		return &entities.Grade{}, err
	}
	return grade, nil
}

func (usecase *GradeUseCase) PutGrade(idGrade int, grade *entities.Grade) error {
	if idGrade <= 0 {
		return fmt.Errorf("error invalid idGrade")
	}
	if grade.Grade < 0 || grade.Grade > 5 {
		return fmt.Errorf("error invalid Grade")
	}
	err := usecase.GradeRepository.UpdateGrade(idGrade, grade)
	if err != nil {
		return err
	}
	return nil
}

func (usecase *GradeUseCase) DeleteGrade(idGrade int) error {
	if idGrade <= 0 {
		return fmt.Errorf("error Invalid ID")
	}
	err := usecase.GradeRepository.DeleteGrade(idGrade)
	if err != nil {
		return err
	}
	return nil
}
