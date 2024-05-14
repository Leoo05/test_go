package course

import (
	"github.com/Leoo05/test_go/internal/domain/course/entities"
	"github.com/Leoo05/test_go/internal/domain/course/repository"
)

type CourseUseCase struct {
	CourseRepository repository.ICourseRepository
}

func NewCourseCase(courseRepository repository.ICourseRepository) *CourseUseCase {
	return &CourseUseCase{
		CourseRepository: courseRepository,
	}
}

func (usecase *CourseUseCase) PostCourse(course *entities.Course) error {
	err := usecase.CourseRepository.CreateCourse(course)
	if err != nil {
		return err
	}
	return nil
}

func (usecase *CourseUseCase) GetCourseByID(idCourse string) (*entities.Course, error) {
	course, err := usecase.CourseRepository.ReadCourse(idCourse)
	if err != nil {
		return &entities.Course{}, err
	}
	return course, err
}

func (usecase *CourseUseCase) PutCourse(idCourse string, course *entities.Course) error {
	err := usecase.CourseRepository.UpdateCourse(idCourse, course)
	if err != nil {
		return err
	}
	return nil
}

func (usecase *CourseUseCase) DeleteCourse(idCourse string) error {
	err := usecase.CourseRepository.DeleteCourse(idCourse)
	if err != nil {
		return err
	}
	return nil
}
