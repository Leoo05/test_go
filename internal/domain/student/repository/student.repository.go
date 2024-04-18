package repository

import "github.com/Leoo05/test_go/internal/domain/student/entities"

type IStudentRepository interface {
	CreateStudent(*entities.Student) error
	ReadStudent(int) (*entities.Student, error)
	UpdateStudent(int, *entities.Student) error
	DeleteStudent(int) error
}
