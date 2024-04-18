package service

import "github.com/Leoo05/test_go/internal/domain/student/entities"

type IStudentService interface {
	GetStudent(int) (*entities.Student, error)
	PostStudent(*entities.Student) error
	PutStudent(int, *entities.Student) error
	DeleteStudent(int) error
}
