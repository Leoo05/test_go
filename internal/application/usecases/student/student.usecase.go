package student

import (
	"fmt"

	"github.com/Leoo05/test_go/internal/domain/student/entities"
	"github.com/Leoo05/test_go/internal/domain/student/repository"
)

type StudentUseCase struct {
	StudentRepository repository.IStudentRepository
}

func NewStudentCase(studentRepository repository.IStudentRepository) *StudentUseCase {
	return &StudentUseCase{
		StudentRepository: studentRepository,
	}
}

func (usecase *StudentUseCase) PostStudent(student *entities.Student) error {
	err := usecase.StudentRepository.CreateStudent(student)
	if err != nil {
		return err
	}
	return nil
}

func (usecase *StudentUseCase) GetStudent(idStudent int) (*entities.Student, error) {
	if idStudent < 0 {
		return &entities.Student{}, fmt.Errorf("error Invalid ID")
	}
	student, err := usecase.StudentRepository.ReadStudent(idStudent)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (usecase *StudentUseCase) PutStudent(idStudent int, student *entities.Student) error {
	if idStudent < 0 {
		return fmt.Errorf("error Invalid ID")
	}
	err := usecase.StudentRepository.UpdateStudent(idStudent, student)
	if err != nil {
		return err
	}
	return nil
}

func (usecase *StudentUseCase) DeleteStudent(idStudent int) error {
	if idStudent < 0 {
		return fmt.Errorf("error Invalid ID")
	}
	err := usecase.StudentRepository.DeleteStudent(idStudent)
	if err != nil {
		return err
	}
	return nil
}
