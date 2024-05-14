package repository

import "github.com/Leoo05/test_go/internal/domain/course/entities"

type ICourseRepository interface {
	CreateCourse(*entities.Course) error
	ReadCourse(string) (*entities.Course, error)
	UpdateCourse(string, *entities.Course) error
	DeleteCourse(string) error
}
