package service

import "github.com/Leoo05/test_go/internal/domain/course/entities"

type ICourseService interface {
	GetCourseByID(string) (*entities.Course, error)
	PostCourse(*entities.Course) error
	PutCourse(string, *entities.Course) error
	DeleteCourse(string) error
}
